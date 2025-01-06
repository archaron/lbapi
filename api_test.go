package lbapi

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"reflect"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/handler"
	"github.com/creachadair/jrpc2/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/archaron/lbapi/events"
	"github.com/archaron/lbapi/types"
)

func testConnector(t *testing.T, m handler.Map, emulateEvents []events.LBEvent, emulateNotifications bool, eventsTrigger chan struct{}) ClientOption {
	t.Helper()
	return func(api *Client) {
		t.Helper()
		api.connector = func(ctx context.Context) error {
			srv := server.NewLocal(m, &server.LocalOptions{
				Server: &jrpc2.ServerOptions{AllowPush: true},
				Client: &jrpc2.ClientOptions{
					OnCallback: api.onCallback,
					OnCancel:   api.hookOnCancel,
					OnStop:     api.hookOnStop,
					OnNotify:   api.hookOnNotify,
				},
			})
			if emulateEvents != nil && eventsTrigger != nil {
				go func() {
					t.Helper()
					<-eventsTrigger
					time.Sleep(time.Millisecond)
					for _, event := range emulateEvents {
						t.Logf("send event: %s", string(event))
						_, err := srv.Server.Callback(context.Background(), "notify", types.LBAPINotification{
							Message: string(event),
							Params:  nil,
						})
						assert.NoError(t, err)
					}
				}()
			}

			if emulateNotifications {
				go func() {
					t.Helper()
					runtime.Gosched()
					assert.NoError(t, srv.Server.Notify(context.Background(), "test_notify", struct{}{}))
				}()
			}

			api.client = srv.Client

			return nil
		}
	}
}

func successfulLogin(ctx context.Context, request *jrpc2.Request) (any, error) {
	return true, nil
}

func unsuccessfulLogin(ctx context.Context, request *jrpc2.Request) (any, error) {
	return false, nil
}

func failedLogin(ctx context.Context, request *jrpc2.Request) (any, error) {
	return false, types.ErrClient
}

func TestClient(t *testing.T) {

	log := slog.New(slog.NewTextHandler(io.Discard, nil))

	t.Run("Fake success login", func(t *testing.T) {
		var (
			callbacked = make(chan struct{})
			notified   = make(chan struct{})
		)
		eventTrigger := make(chan struct{})
		client := NewClient(ClientConfig{}, log, testConnector(t, handler.Map{
			LoginRequest{}.Method():              successfulLogin,
			new(SystemSubscribeRequest).Method(): successfulLogin,
		}, []events.LBEvent{events.BlockVgEvent}, true, eventTrigger), WithOnCallback(func(ctx context.Context, request *jrpc2.Request) (any, error) {
			close(callbacked)
			return true, nil
		}), WithOnNotify(func(request *jrpc2.Request) {
			close(notified)
		}))

		now := time.Now()

		require.NoError(t, client.ConnectAndLogin(context.Background()))
		ok, err := client.Subscribe(context.Background(), []events.LBEvent{events.BlockVgEvent})
		require.NoError(t, err)
		require.True(t, ok)

		channels := map[string]chan struct{}{
			"notified":   notified,
			"callbacked": callbacked,
		}

		var wg sync.WaitGroup
		wg.Add(len(channels))
		eventTrigger <- struct{}{}
		for name, ch := range channels {
			go func() {
				defer wg.Done()
				select {
				case <-ch:
					t.Logf("%s: %s", name, time.Since(now))
				case <-time.After(5 * time.Millisecond):
					t.Errorf("Should be %s", name)
				}
			}()
		}

		wg.Wait()
	})

	t.Run("Fake unsuccessful login", func(t *testing.T) {
		eventTrigger := make(chan struct{})
		client := NewClient(ClientConfig{}, log, testConnector(t, handler.Map{
			"Login": unsuccessfulLogin,
		}, []events.LBEvent{events.BlockVgEvent}, true, eventTrigger))
		close(eventTrigger)
		require.ErrorContains(t, client.ConnectAndLogin(context.Background()), types.ErrLogin.Error())
	})

	t.Run("Fake fail login", func(t *testing.T) {
		eventTrigger := make(chan struct{})
		client := NewClient(ClientConfig{}, log, testConnector(t, handler.Map{
			"Login": failedLogin,
		}, []events.LBEvent{events.BlockVgEvent}, true, eventTrigger))

		err := client.ConnectAndLogin(context.Background())
		close(eventTrigger)
		require.Error(t, err)
		require.ErrorContains(t, errors.Unwrap(err), types.ErrClient.Error())
	})

	t.Run("Bad connector and cancel/stop hooks", func(t *testing.T) {
		var (
			cancelled = make(chan struct{})
			stopped   = make(chan struct{})
		)

		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
		}))

		client := NewClient(ClientConfig{
			Address: srv.Listener.Addr().String(),
		}, log, WithOnCancel(func(*jrpc2.Client, *jrpc2.Response) {
			close(cancelled)
		}), WithOnStop(func(*jrpc2.Client, error) {
			close(stopped)
		}))

		now := time.Now()
		require.ErrorContains(t, client.ConnectAndLogin(context.Background()), types.ErrClient.Error())

		channels := map[string]chan struct{}{
			"cancelled": cancelled,
			"stopped":   stopped,
		}

		var wg sync.WaitGroup
		wg.Add(len(channels))

		for name, ch := range channels {
			go func() {
				defer wg.Done()
				select {
				case <-ch:
					t.Logf("%s: %s", name, time.Since(now))
				case <-time.After(time.Millisecond):
					t.Errorf("Should be %s", name)
				}
			}()
		}

		wg.Wait()
	})

	t.Run("Close", func(t *testing.T) {
		client := NewClient(ClientConfig{}, log, testConnector(t, handler.Map{
			"Login": successfulLogin,
		}, nil, false, nil))

		require.NoError(t, client.ConnectAndLogin(context.Background()))
		require.NoError(t, client.Close())
	})

	t.Run("Test OnCallback", func(t *testing.T) {
		eventTrigger := make(chan struct{})
		client := NewClient(ClientConfig{}, log, testConnector(t, handler.Map{
			LoginRequest{}.Method():              successfulLogin,
			new(SystemSubscribeRequest).Method(): successfulLogin,
		}, events.AllLBEvents, false, eventTrigger))

		now := time.Now()
		require.NoError(t, client.ConnectAndLogin(context.Background()))
		ok, err := client.Subscribe(context.Background(), events.AllLBEvents)
		require.NoError(t, err)
		require.True(t, ok)

		ch := client.GetEventsChannel()
		var wg sync.WaitGroup
		wg.Add(len(events.AllLBEvents))
		close(eventTrigger)
		for i := range events.AllLBEvents {
			go func() {
				defer wg.Done()
				select {
				case e := <-ch:
					require.NotNil(t, e)
					event, ok := e.(events.EventInterface)
					require.Truef(t, ok, "event should be of type EventInterface, has:%s", reflect.TypeOf(e).String())
					t.Logf("%T: %s", event, time.Since(now))
					return
				case <-time.After(500 * time.Millisecond):
					t.Errorf("Event %d is not received", i)
					return
				}
			}()
		}

		wg.Wait()
		require.NoError(t, client.Close())
	})

	t.Run("Test OnCallback unknown event", func(t *testing.T) {

	})

	t.Run("Test OnCallback invalid params", func(t *testing.T) {

	})

}
