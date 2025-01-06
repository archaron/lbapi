package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// getVlans
type (
	GetVlansRequest struct {
		VlanId    int    `json:"vlan_id"`
		OuterVlan int    `json:"outer_vlan,omitempty"`
		Name      string `json:"name,omitempty"`
	}
)

func (GetVlansRequest) Method() string {
	return "getVlans"
}

func (api *Client) GetVlans(ctx context.Context, request GetVlansRequest) (vlans *[]types.VlanRecord, err error) {
	if err := api.client.CallResult(ctx, "getVlans", &request, &vlans); err != nil {
		return nil, err
	}
	return
}
