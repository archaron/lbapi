package lbapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"

	"github.com/archaron/lbapi/events"
	"github.com/archaron/lbapi/types"
)

type (
	Client struct {
		cfg ClientConfig
		log *slog.Logger

		client     *jrpc2.Client
		connection net.Conn

		reconnectTimer         *time.Timer
		fails                  int
		currentReconnectPeriod time.Duration

		connector func(context.Context) error

		subscriptions []events.LBEvent
		eventsChannel events.Channel

		hookOnNotify   func(*jrpc2.Request)
		hookOnCallback jrpc2.Handler
		hookOnCancel   func(cli *jrpc2.Client, rsp *jrpc2.Response)
		hookOnStop     func(cli *jrpc2.Client, err error)
	}

	ClientConfig struct {
		Address  string `env:"LB_API_ADDRESS" yaml:"address" usage:"LANBilling API address"`     // LanBilling server address
		Username string `env:"LB_API_USERNAME" yaml:"username" usage:"LANBilling API username "` // Agent username
		Password string `env:"LB_API_PASSWORD" yaml:"password" usage:"LANBilling API password"`  // Agent password

		MaxFails int           `env:"LB_API_MAX_FAILS" yaml:"max_fails" default:"5"` // Maximum reconnect fails before reconnect
		Timeout  time.Duration `env:"LB_API_TIMEOUT" yaml:"timeout" default:"5s"`    // (re)Connect timeout

		// Periodic ping interval to check if server is reachable.
		ReconnectPeriod time.Duration `env:"LB_API_RECONNECT_PERIOD" yaml:"reconnect_period" default:"10s"`

		// Reconnect backoff configuration.
		// After a failed reconnect, period is multiplied by BackoffFactor up to BackoffMax.
		// On successful reconnect, period resets to ReconnectPeriod.
		ReconnectBackoffMin    time.Duration `env:"LB_API_RECONNECT_BACKOFF_MIN" yaml:"reconnect_backoff_min" default:"10s"`
		ReconnectBackoffMax    time.Duration `env:"LB_API_RECONNECT_BACKOFF_MAX" yaml:"reconnect_backoff_max" default:"5m"`
		ReconnectBackoffFactor float64       `env:"LB_API_RECONNECT_BACKOFF_FACTOR" yaml:"reconnect_backoff_factor" default:"2"`
	}

	ClientOption func(*Client)
)

func WithOnNotify(handler func(*jrpc2.Request)) ClientOption {
	return func(api *Client) {
		api.hookOnNotify = handler
	}
}

func WithOnCallback(handler jrpc2.Handler) ClientOption {
	return func(api *Client) {
		api.hookOnCallback = handler
	}
}

func WithOnCancel(handler func(cli *jrpc2.Client, rsp *jrpc2.Response)) ClientOption {
	return func(api *Client) {
		api.hookOnCancel = handler
	}
}

func WithOnStop(handler func(cli *jrpc2.Client, err error)) ClientOption {
	return func(api *Client) {
		api.hookOnStop = handler
	}
}

func defaultConnector(api *Client) {
	api.connector = func(top context.Context) error {
		ctx, cancel := context.WithTimeout(top, api.cfg.Timeout)
		defer cancel()

		network, address := jrpc2.Network(api.cfg.Address)
		conn, err := new(net.Dialer).DialContext(ctx, network, address)
		if err != nil {
			return err
		}

		if api.client == nil {
			api.client = jrpc2.NewClient(channel.Line(conn, conn), &jrpc2.ClientOptions{
				Logger: func(text string) {
					api.log.Debug("API: " + text)
				},
				OnNotify:   api.hookOnNotify,
				OnCallback: api.onCallback,
				OnCancel:   api.hookOnCancel,
				OnStop:     api.hookOnStop,
			})
		}

		if api.cfg.ReconnectPeriod > 0 && api.reconnectTimer == nil {
			api.log.Debug("will setup LBreconnect", slog.Duration("reconnect_period", api.cfg.ReconnectPeriod))
			go api.connectionWatchdog(top)
		}

		return nil
	}
}

// NewClient creates a new LanBilling API client
func NewClient(cfg ClientConfig, log *slog.Logger, opts ...ClientOption) *Client {
	log.Debug("LanBilling API created")

	if cfg.Timeout <= 0 {
		cfg.Timeout = 5 * time.Second
	}
	if cfg.ReconnectBackoffMin <= 0 {
		cfg.ReconnectBackoffMin = 10 * time.Second
	}
	if cfg.ReconnectBackoffMax <= 0 {
		cfg.ReconnectBackoffMax = 5 * time.Minute
	}
	if cfg.ReconnectBackoffFactor <= 0 {
		cfg.ReconnectBackoffFactor = 2
	}

	api := &Client{
		cfg:                    cfg,
		log:                    log,
		currentReconnectPeriod: cfg.ReconnectBackoffMin,
	}

	for _, opt := range opts {
		opt(api)
	}

	if api.connector == nil {
		defaultConnector(api)
	}

	return api
}

// ConnectAndLogin - connects to LanBilling server and tries to log in
func (api *Client) ConnectAndLogin(ctx context.Context) error {
	if err := api.connector(ctx); err != nil {
		return err
	}

	if ok, err := api.Login(ctx, LoginRequest{
		Login:    api.cfg.Username,
		Password: api.cfg.Password,
	}); err != nil {
		return types.ParseJSONRPCError(err)
	} else if !ok {
		return types.ErrLogin
	}
	return nil
}

func (api *Client) connectionWatchdog(top context.Context) {
	api.log.Debug("reconnect timer callback start")
	api.reconnectTimer = time.NewTimer(api.cfg.ReconnectPeriod)

	for {
		select {
		case <-top.Done():
			api.log.Debug("reconnect callback stop")
			api.reconnectTimer.Stop()
			api.reconnectTimer = nil
			return

		case <-api.reconnectTimer.C:
			api.checkAndReconnect(top)

			api.reconnectTimer.Reset(api.currentReconnectPeriod)
		}
	}
}

func (api *Client) checkAndReconnect(top context.Context) {
	ctx, cancel := context.WithTimeout(top, api.cfg.Timeout)
	defer cancel()

	if api.client != nil {
		if _, err := api.GetJAPIVersion(ctx); err != nil {
			api.log.Warn("LANBilling api ping failed", slog.Int("fails", api.fails), slog.Any("error", err))
			api.fails++
		} else {
			api.fails = 0
			api.currentReconnectPeriod = api.cfg.ReconnectBackoffMin
		}
	}

	if api.fails <= api.cfg.MaxFails {
		return
	}

	api.log.Warn("LANBilling API max ping tries reached, reconnecting")

	if err := api.Close(); err != nil {
		api.log.Warn("cannot close client", slog.Any("error", err))
	}

	if err := api.ConnectAndLogin(ctx); err != nil {
		api.log.Warn("cannot reconnect client", slog.Any("error", err))
		// Экспоненциальный бэкофф
		api.currentReconnectPeriod = time.Duration(float64(api.currentReconnectPeriod) * api.cfg.ReconnectBackoffFactor)
		if api.currentReconnectPeriod > api.cfg.ReconnectBackoffMax {
			api.currentReconnectPeriod = api.cfg.ReconnectBackoffMax
		}
		api.log.Info("next reconnect attempt", slog.Duration("after", api.currentReconnectPeriod))
		api.fails = 0
		return
	}

	api.log.Info("reconnected to LANBilling")

	if len(api.subscriptions) == 0 {
		api.fails = 0
		api.currentReconnectPeriod = api.cfg.ReconnectBackoffMin
		return
	}

	if ok, err := api.SystemSubscribeMultiple(ctx, api.subscriptions); err != nil || !ok {
		api.log.Warn("cannot resubscribe to LB events", slog.Any("error", err))
		return
	}

	api.fails = 0
	api.currentReconnectPeriod = api.cfg.ReconnectBackoffMin
}

// Close all connections to server and clear subscriptions
func (api *Client) Close() error {

	if api.client != nil {
		if err := api.client.Close(); err != nil {
			return err
		}
	}

	if api.connection != nil {
		if err := api.connection.Close(); err != nil {
			return err
		}
	}

	api.connection = nil
	api.client = nil
	if api.eventsChannel != nil {
		close(api.eventsChannel)
		api.eventsChannel = nil
	}
	api.subscriptions = nil

	return nil
}

func IPv4MaskString(m []byte) string {
	if len(m) != 4 {
		panic("ipv4Mask: len must be 4 bytes")
	}

	return fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
}

// Subscribe to LanBilling events, use GetEventsChannel to obtain notifications channel
func (api *Client) Subscribe(ctx context.Context, eventsList []events.LBEvent) (result bool, err error) {
	if api.eventsChannel == nil {
		api.eventsChannel = make(events.Channel, 100)
	}

	api.subscriptions = append(api.subscriptions, eventsList...)
	return api.SystemSubscribeMultiple(ctx, eventsList)
}

// GetEventsChannel returns EventsChannel which will receive notifications to subscribed events
func (api *Client) GetEventsChannel() events.Channel {
	return api.eventsChannel
}

// Call the LanBilling api
func (api *Client) Call(ctx context.Context, req types.LanbillingRequest, res interface{}) error {
	return api.client.CallResult(ctx, req.Method(), req, &res)
}

func (api *Client) onCallback(ctx context.Context, request *jrpc2.Request) (interface{}, error) {
	var (
		hookResult interface{}
		hookError  error
	)

	if api.hookOnCallback != nil {
		hookResult, hookError = api.hookOnCallback(ctx, request)
	}

	if request.Method() != "notify" {
		return hookResult, hookError
	}

	// ignore if there is no subscriptions
	if api.eventsChannel == nil {
		return hookResult, hookError
	}

	note := &types.LBAPINotification{}

	if err := request.UnmarshalParams(note); err != nil {
		return nil, err
	}

	var x events.EventInterface

	switch note.Message {
	case events.CreateAgreementEvent:
		x = &events.LBEventCreateAgreement{}
	case events.CloseAgreementEvent:
		x = &events.LBEventCloseAgreement{}
	case events.ChangeTariffVGEvent:
		x = &events.LBEventChangeTariffVG{}
	case events.DelAgentEvent:
		x = &events.LBEventDelAgent{}
	case events.ChangeAgentEvent:
		x = &events.LBEventChangeAgent{}
	case events.ChangeAgentOptionEvent:
		x = &events.LBEventChangeAgentOption{}
	case events.SetDeviceEvent:
		x = &events.LBEventSetDevice{}
	case events.DelDeviceEvent:
		x = &events.LBEventDelDevice{}
	case events.ChangeAuthHistoryEvent:
		x = &events.LBEventChangeAuthHistory{}
	case events.ChangeCategoryEvent:
		x = &events.LBEventChangeCategory{}
	case events.ChangeCategoryModifierEvent:
		x = &events.LBEventChangeCategoryModifier{}
	case events.ChangeDeviceGroupEvent:
		x = &events.LBEventChangeDeviceGroup{}
	case events.ChangeDeviceGroupMemberEvent:
		x = &events.LBEventChangeDeviceGroupMember{}
	case events.ChangeDeviceGroupsOptionEvent:
		x = &events.LBEventChangeDeviceGroupsOption{}
	case events.ChangeDeviceGroupVlanEvent:
		x = &events.LBEventChangeDeviceGroupVlan{}
	case events.ChangeDictionaryEvent:
		x = &events.LBEventChangeDictionary{}
	case events.DeleteDictionaryEvent:
		x = &events.LBEventDeleteDictionary{}
	case events.ChangeGrStaffEvent:
		x = &events.LBEventChangeGrStaff{}
	case events.ChangeVgNetworkEvent:
		x = &events.LBEventChangeVgNetwork{}
	case events.ChangeOptionEvent:
		x = &events.LBEventChangeOption{}
	case events.ChangeOptionsEvent:
		x = &events.LBEventChangeOptions{}
	case events.ChangePayCardEvent:
		x = &events.LBEventChangePayCard{}
	case events.ChangePortEvent:
		x = &events.LBEventChangePort{}
	case events.ChangeVgPortEvent:
		x = &events.LBEventChangeVgPort{}
	case events.ChangeRadiusAttrEvent:
		x = &events.LBEventChangeRadiusAttr{}
	case events.ChangeSegmentEvent:
		x = &events.LBEventChangeSegment{}
	case events.ChangeStaffEvent:
		x = &events.LBEventChangeStaff{}
	case events.ChangeTimeDiscountEvent:
		x = &events.LBEventChangeTimeDiscount{}
	case events.ChangeSizeDiscountEvent:
		x = &events.LBEventChangeSizeDiscount{}
	case events.ChangeVlanEvent:
		x = &events.LBEventChangeVlan{}
	case events.ChangeWeekendEvent:
		x = &events.LBEventChangeWeekend{}
	case events.CreateVgroupEvent:
		x = &events.LBEventCreateVgroup{}
	case events.ChangeVgroupEvent:
		x = &events.LBEventChangeVgroup{}
	case events.EndRadiusStatEvent:
		x = &events.LBEventEndRadiusStat{}
	case events.BlockVgEvent:
		x = &events.LBEventBlockVg{}
	case events.BlockVgTaskEvent:
		x = &events.LBEventBlockVgTask{}
	case events.DelRadiusVgroupEvent:
		x = &events.LBEventDelRadiusVgroup{}
	case events.ChangeTariffEvent:
		x = &events.LBEventChangeTariff{}
	case events.ChangeSizeShapeEvent:
		x = &events.LBEventChangeSizeShape{}
	case events.ChangeTimeShapeEvent:
		x = &events.LBEventChangeTimeShape{}
	case events.ChangeShapePolicyEvent:
		x = &events.LBEventChangeShapePolicy{}
	case events.RunUsboxScriptEvent:
		x = &events.LBEventRunUsboxScript{}
	case events.ToInduceArcdTurboshapesSynchronizationEvent:
		x = &events.LBEventToInduceArcdTurboshapesSynchronization{}
	case events.ChangeRnasEvent:
		x = &events.LBEventChangeRnas{}
	case events.DeleteRnasEvent:
		x = &events.LBEventDeleteRnas{}
	case events.ChangeRnasOptionsEvent:
		x = &events.LBEventChangeRnasOptions{}
	case events.DeleteRnasOptionsEvent:
		x = &events.LBEventDeleteRnasOptions{}
	case events.PaymentEvent:
		x = &events.LBEventPayment{}
	case events.PromisePaymentEvent:
		x = &events.LBEventPromisePayment{}
	case events.GetRadiussessionsFromAgentEvent:
		x = &events.LBEventGetRadiussessionsFromAgent{}
	case events.GetRadiusAuthsFromAgentEvent:
		x = &events.LBEventGetRadiusAuthsFromAgent{}
	case events.StopRadSessionEvent:
		x = &events.LBEventStopRadSession{}
	case events.DeleteRadSessionEvent:
		x = &events.LBEventDeleteRadSession{}
	case events.StartRecalcEvent:
		x = &events.LBEventStartRecalc{}
	case events.ChangeGponPortsEvent:
		x = &events.LBEventChangeGponPorts{}
	case events.ChangeVgroupsAccessOverrideEvent:
		x = &events.LBEventChangeVgroupsAccessOverride{}
	default:
		return nil, types.NewComplexErrorf("unknown notify type: %s, params: %s", note.Message, note.Params)
	}

	if err := json.Unmarshal(note.Params, x); err != nil {
		return nil, err
	}

	api.eventsChannel <- x

	return hookResult, hookError
}
