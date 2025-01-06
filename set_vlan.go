package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type SetVlanRequest types.VlanRecord

func (SetVlanRequest) Method() string {
	return "setVlan"
}

// SetVlan sets vlan information, returns vlan ID
func (api *Client) SetVlan(ctx context.Context, params types.VlanRecord) (int, error) {
	var result int
	if err := api.client.CallResult(ctx, "setVlan", params, &result); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}

	return result, nil
}
