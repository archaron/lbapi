package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type SetVgMacAddrRequest struct {
	Mac      string `json:"mac"`
	VgId     int    `json:"vg_id"`
	RecordId *int   `json:"record_id,omitempty"`
}

func (SetVgMacAddrRequest) Method() string {
	return "setVgMacAddr"
}

func (api *Client) SetVgMacAddr(ctx context.Context, request SetVgMacAddrRequest) (int, error) {
	var result int
	if err := api.client.CallResult(ctx, "setVgMacAddr", request, &result); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}

	return result, nil
}
