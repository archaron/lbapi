package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// getVGroups
type GetVGroupsRequest struct {
	IsTemplate      bool   `json:"is_template"`
	Address         string `json:"address"`
	Name            string `json:"name"`
	AgrmNum         string `json:"agrm_num"`
	SmartCardSerial string `json:"smart_card_serial"`
	Descr           string `json:"descr"`
	EquipChipid     string `json:"equip_chipid"`
	EquipMac        string `json:"equip_mac"`
	EquipSerial     string `json:"equip_serial"`
	IP              string `json:"ip"`
	PayCode         string `json:"pay_code"`
	Phone           string `json:"phone"`

	VgMac      string            `json:"vg_mac"`
	Archive    int               `json:"archive"`
	Blocked    int               `json:"blocked"`
	AgentTypes []types.AgentType `json:"agent_types"`

	GetFullData bool `json:"get_full_data"`
	Count       bool `json:"count"`
	Nodata      bool `json:"nodata"`
	types.Pagination
}

func (GetVGroupsRequest) Method() string {
	return "getVgroups"
}

func (api *Client) GetVGroups(ctx context.Context, request GetVGroupsRequest) ([]types.VGroupRecord, error) {
	var result []types.VGroupRecord
	if err := api.client.CallResult(ctx, "getVgroups", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
