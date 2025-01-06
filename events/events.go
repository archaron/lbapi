package events

import (
	"fmt"

	"github.com/archaron/lbapi/types"
)

const (
	CreateAgreementEvent                        = "create_agreement"
	CloseAgreementEvent                         = "close_agreement"
	ChangeTariffVGEvent                         = "change_tariff_vg"
	DelAgentEvent                               = "del_agent"
	ChangeAgentEvent                            = "change_agent"
	ChangeAgentOptionEvent                      = "change_agent_option"
	SetDeviceEvent                              = "set_device"
	DelDeviceEvent                              = "del_device"
	ChangeAuthHistoryEvent                      = "change_auth_history"
	ChangeCategoryEvent                         = "change_category"
	ChangeCategoryModifierEvent                 = "change_category_modifier"
	ChangeDeviceGroupEvent                      = "change_device_group"
	ChangeDeviceGroupMemberEvent                = "change_device_group_member"
	ChangeDeviceGroupsOptionEvent               = "change_device_groups_option"
	ChangeDeviceGroupVlanEvent                  = "change_device_group_vlan"
	ChangeDictionaryEvent                       = "change_dictionary"
	DeleteDictionaryEvent                       = "delete_dictionary"
	ChangeGrStaffEvent                          = "change_gr_staff"
	ChangeVgNetworkEvent                        = "change_vg_network"
	ChangeOptionEvent                           = "change_option"
	ChangeOptionsEvent                          = "change_options"
	ChangePayCardEvent                          = "change_pay_card"
	ChangePortEvent                             = "change_port"
	ChangeVgPortEvent                           = "change_vg_port"
	ChangeRadiusAttrEvent                       = "change_radius_attr"
	ChangeSegmentEvent                          = "change_segment"
	ChangeStaffEvent                            = "change_staff"
	ChangeTimeDiscountEvent                     = "change_time_discount"
	ChangeSizeDiscountEvent                     = "change_size_discount"
	ChangeVlanEvent                             = "change_vlan"
	ChangeWeekendEvent                          = "change_weekend"
	CreateVgroupEvent                           = "create_vgroup"
	ChangeVgroupEvent                           = "change_vgroup"
	EndRadiusStatEvent                          = "end_radius_stat"
	BlockVgEvent                                = "block_vg"
	BlockVgTaskEvent                            = "block_vg_task"
	DelRadiusVgroupEvent                        = "del_radius_vgroup"
	ChangeTariffEvent                           = "change_tariff"
	ChangeSizeShapeEvent                        = "change_size_shape"
	ChangeTimeShapeEvent                        = "change_time_shape"
	ChangeShapePolicyEvent                      = "change_shape_policy"
	RunUsboxScriptEvent                         = "run_usbox_script"
	ToInduceArcdTurboshapesSynchronizationEvent = "to_induce_arcd_turboshapes_synchronization"
	ChangeRnasEvent                             = "change_rnas"
	DeleteRnasEvent                             = "delete_rnas"
	ChangeRnasOptionsEvent                      = "change_rnas_options"
	DeleteRnasOptionsEvent                      = "delete_rnas_options"
	PaymentEvent                                = "payment"
	PromisePaymentEvent                         = "promise_payment"
	GetRadiussessionsFromAgentEvent             = "get_radiussessions_from_agent"
	GetRadiusAuthsFromAgentEvent                = "get_radius_auths_from_agent"
	StopRadSessionEvent                         = "stop_rad_session"
	DeleteRadSessionEvent                       = "delete_rad_session"
	StartRecalcEvent                            = "start_recalc"
	ChangeGponPortsEvent                        = "change_gpon_ports"
	ChangeVgroupsAccessOverrideEvent            = "change_vgroups_access_override"
)

var AllLBEvents = []LBEvent{
	CreateAgreementEvent,
	CloseAgreementEvent,
	ChangeTariffVGEvent,
	DelAgentEvent,
	ChangeAgentEvent,
	ChangeAgentOptionEvent,
	SetDeviceEvent,
	DelDeviceEvent,
	ChangeAuthHistoryEvent,
	ChangeCategoryEvent,
	ChangeCategoryModifierEvent,
	ChangeDeviceGroupEvent,
	ChangeDeviceGroupMemberEvent,
	ChangeDeviceGroupsOptionEvent,
	ChangeDeviceGroupVlanEvent,
	ChangeDictionaryEvent,
	DeleteDictionaryEvent,
	ChangeGrStaffEvent,
	ChangeVgNetworkEvent,
	ChangeOptionEvent,
	ChangeOptionsEvent,
	ChangePayCardEvent,
	ChangePortEvent,
	ChangeVgPortEvent,
	ChangeRadiusAttrEvent,
	ChangeSegmentEvent,
	ChangeStaffEvent,
	ChangeTimeDiscountEvent,
	ChangeSizeDiscountEvent,
	ChangeVlanEvent,
	ChangeWeekendEvent,
	CreateVgroupEvent,
	ChangeVgroupEvent,
	EndRadiusStatEvent,
	BlockVgEvent,
	BlockVgTaskEvent,
	DelRadiusVgroupEvent,
	ChangeTariffEvent,
	ChangeSizeShapeEvent,
	ChangeTimeShapeEvent,
	ChangeShapePolicyEvent,
	RunUsboxScriptEvent,
	ToInduceArcdTurboshapesSynchronizationEvent,
	ChangeRnasEvent,
	DeleteRnasEvent,
	ChangeRnasOptionsEvent,
	DeleteRnasOptionsEvent,
	PaymentEvent,
	PromisePaymentEvent,
	GetRadiussessionsFromAgentEvent,
	GetRadiusAuthsFromAgentEvent,
	StopRadSessionEvent,
	DeleteRadSessionEvent,
	StartRecalcEvent,
	ChangeGponPortsEvent,
	ChangeVgroupsAccessOverrideEvent,
}

type EventInterface interface {
}

type (
	LBEvent string

	LBEventSubscription struct {
		event LBEvent
	}

	Channel chan EventInterface

	LBEventCreateAgreement struct {
	}

	LBEventCloseAgreement struct {
	}

	LBEventChangeTariffVG struct {
		Changetime string `json:"changetime"`
		Requestby  int    `json:"requestby"`
		Tarid      int    `json:"tarid"`
		TaridOld   int    `json:"tarid_old"`
		VgID       int    `json:"vg_id"`
	}

	LBEventDelAgent struct {
		ID int `json:"id"`
	}

	LBEventChangeAgent struct {
		ID int `json:"id"`
	}

	LBEventChangeAgentOption struct {
		AgentID int    `json:"agent_id"`
		Name    string `json:"name"`
	}

	LBEventSetDevice struct {
	}

	LBEventDelDevice struct {
	}

	LBEventChangeAuthHistory struct {
	}

	LBEventChangeCategory struct {
	}

	LBEventChangeCategoryModifier struct {
		VgID int `json:"vg_id"`
	}

	LBEventChangeDeviceGroup struct {
	}

	LBEventChangeDeviceGroupMember struct {
	}

	LBEventChangeDeviceGroupsOption struct {
	}

	LBEventChangeDeviceGroupVlan struct {
	}

	LBEventChangeDictionary struct {
	}

	LBEventDeleteDictionary struct {
	}

	LBEventChangeGrStaff struct {
	}

	LBEventChangeVgNetwork struct {
		AgentID         int                      `json:"agent_id"`
		VgID            int                      `json:"vg_id"`
		DeletedMAC      string                   `json:"deleted_mac"`
		DeletedSegments []types.LBNetworkSegment `json:"deleted_segments"`
	}

	LBEventChangeOption struct {
	}

	LBEventChangeOptions struct {
	}

	LBEventChangePayCard struct {
	}

	LBEventChangePort struct {
	}

	LBEventChangeVgPort struct {
	}

	LBEventChangeRadiusAttr struct {
	}

	LBEventChangeSegment struct {
		SegmentID int  `json:"segment_id"`
		Deleted   bool `json:"deleted"`
	}

	LBEventChangeStaff struct {
	}

	LBEventChangeTimeDiscount struct {
	}

	LBEventChangeSizeDiscount struct {
	}

	LBEventChangeVlan struct {
		RecordID int `json:"record_id"`
	}

	LBEventChangeWeekend struct {
	}

	LBEventCreateVgroup struct {
	}

	LBEventChangeVgroup struct {
		AgreementID int `json:"agrm_id"`
		VgID        int `json:"vg_id"`
	}

	LBEventEndRadiusStat struct {
	}

	LBEventBlockVg struct {
		BlkReq      types.LBBlockRequest `json:"blk_req"`
		BlockRaspID int                  `json:"block_rasp_id"`
		ChangeTime  *types.LBTime        `json:"change_time"`
		RequestBy   int                  `json:"request_by"`
		VgID        int                  `json:"vgid"`
	}

	LBEventBlockVgTask struct {
		BlkReq   types.LBBlockRequest `json:"blk_req"`
		RecordID int                  `json:"recordid"`
		VgID     int                  `json:"vgid"`
	}

	LBEventDelRadiusVgroup struct {
	}

	LBEventChangeTariff struct {
	}

	LBEventChangeSizeShape struct {
	}

	LBEventChangeTimeShape struct {
	}

	LBEventChangeShapePolicy struct {
	}

	LBEventRunUsboxScript struct {
	}

	LBEventToInduceArcdTurboshapesSynchronization struct {
	}

	LBEventChangeRnas struct {
		NasID int `json:"nas_id"`
	}

	LBEventDeleteRnas struct {
	}

	LBEventChangeRnasOptions struct {
		OptionID int `json:"option_id"`
	}

	LBEventDeleteRnasOptions struct {
	}

	LBEventPayment struct {
		AgrmID      int     `json:"agrmid"`
		Amount      float64 `json:"amount"`
		Balance     float64 `json:"balance"`
		CashCode    int     `json:"cashcode"`
		FromAgrmID  int     `json:"fromagrmid"`
		IsCancelled bool    `json:"is_cancelled"`
		Modperson   int     `json:"modperson"`
		Payclass    int     `json:"payclass"`
		PaymentID   int     `json:"paymentid"`
	}

	LBEventPromisePayment struct {
	}

	LBEventGetRadiussessionsFromAgent struct {
	}

	LBEventGetRadiusAuthsFromAgent struct {
		AgentID    int    `json:"agent_id"`
		Agrmids    []int  `json:"agrmids"`
		Ani        string `json:"ani"`
		DeviceName string `json:"device_name"`
		IP         string `json:"ip"`
		Nodata     bool   `json:"nodata"`
		SessionID  string `json:"sessionid"`
		VgLogin    string `json:"vglogin"`
		WaitingID  string `json:"waiting_id"`
	}

	LBEventStopRadSession struct {
	}

	LBEventDeleteRadSession struct {
	}

	LBEventStartRecalc struct {
	}

	LBEventChangeGponPorts struct {
	}

	LBEventChangeVgroupsAccessOverride struct {
	}
)

func (L LBEventChangeSegment) String() string {
	if L.Deleted {
		return fmt.Sprintf("Удалить сегмент: SegmentId:%d", L.SegmentID)
	} else {
		return fmt.Sprintf("Изменить сегмент: SegmentId:%d", L.SegmentID)
	}
}

func (L LBEventBlockVgTask) String() string {
	return fmt.Sprintf("Изменить статус УЗ: VgID:%d [%s] №Записи: %d", L.VgID, L.BlkReq, L.RecordID)
}

func (L LBEventBlockVg) String() string {

	return fmt.Sprintf("Запланировать изменение статуса УЗ: VgID:%d [%s] №Записи: %d, Когда: %s, ЗапросОт: %d", L.VgID, L.BlkReq, L.BlockRaspID, L.ChangeTime, L.RequestBy)
}

func (L LBEventChangeVgNetwork) String() string {
	result := fmt.Sprintf("Изменена сеть УЗ: AgentID: %d, VgroupID: %d", L.AgentID, L.VgID)
	if L.DeletedMAC != "" {
		result += " Удален MAC: '" + L.DeletedMAC + "'"
	}

	if len(L.DeletedSegments) > 0 {
		result += " Удалены сети: [ "
		for i := 0; i < len(L.DeletedSegments); i++ {
			result += L.DeletedSegments[i].String() + " "
		}
		result += " ] "
	}
	return result
}

func (L LBEventChangeVgroup) String() string {
	if L.AgreementID > 0 {
		return fmt.Sprintf("ChangeVGroup: AgreementID:%d VgID: %d ", L.AgreementID, L.VgID)
	} else {
		return fmt.Sprintf("ChangeVGroup: VgID: %d ", L.VgID)
	}

}

func (L LBEventChangeAgentOption) String() string {
	return fmt.Sprintf("ChangeAgentOption: AgentID: %d, Option: %s", L.AgentID, L.Name)
}
