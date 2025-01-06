package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetVgMacAddrsRequest struct {
	AgentID       int   `json:"agent_id"`
	TimeMarkStart int64 `json:"time_mark_start"`
	types.Pagination
}

func (GetVgMacAddrsRequest) Method() string {
	return "getVgMacAddrs"
}

func (api *Client) GetVgMacAddrs(ctx context.Context, request GetVgMacAddrsRequest) ([]types.LBMacRecord, error) {
	var result []types.LBMacRecord
	if err := api.client.CallResult(ctx, "getVgMacAddrs", request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return result, nil
}
