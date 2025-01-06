package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetRNasRequest struct {
	AgentID       int `json:"agent_id"`
	TimeMarkStart int `json:"time_mark_start"`

	types.Pagination
}

func (GetRNasRequest) Method() string { return "getRnas" }

func (api *Client) GetRNas(ctx context.Context, request GetRNasRequest) ([]types.RNas, error) {
	var response []types.RNas
	if err := api.client.CallResult(ctx, "getRnas", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
