package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type SetPortRequest struct {
	Comment     string `json:"comment,omitempty"`
	DeviceId    int    `json:"device_id,omitempty"`
	Media       string `json:"media,omitempty"`
	Name        int    `json:"name,omitempty"`
	PolicyId    int    `json:"policy_id,omitempty"`
	PortId      int    `json:"port_id,omitempty"`
	PrototypeId int    `json:"prototype_id,omitempty"`
	Speed       string `json:"speed,omitempty"`
	Tpl         int    `json:"tpl,omitempty"`
	VlanId      int    `json:"vlan_id,omitempty"`
	VgId        int    `json:"vg_id,omitempty"`
}

func (SetPortRequest) Method() string {
	return "setPort"
}

func (api *Client) SetPort(ctx context.Context, params SetPortRequest) (int, error) {
	var PortId int
	if err := api.client.CallResult(ctx, "setPort", params, &PortId); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}

	return PortId, nil
}
