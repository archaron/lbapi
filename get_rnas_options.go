package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetRNasOptionsRequest struct {
	AgentID       int `json:"agent_id"`
	TimeMarkStart int `json:"time_mark_start"`
	types.Pagination
}

func (GetRNasOptionsRequest) Method() string {
	return "getRnasOptions"
}

func (api *Client) GetRNasOptions(ctx context.Context, request GetRNasOptionsRequest) ([]types.LBRNasOption, error) {
	var result []types.LBRNasOption
	if err := api.client.CallResult(ctx, "getRnasOptions", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
