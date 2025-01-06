package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetRadiusAttrsRequest struct {
	AgentID       int `json:"agent_id"`
	TimeMarkStart int `json:"time_mark_start"`

	types.Pagination
}

func (GetRadiusAttrsRequest) Method() string { return "getRadiusAttrs" }

func (api *Client) GetRadiusAttrs(ctx context.Context, request GetRadiusAttrsRequest) ([]types.RadiusAttrs, error) {
	var response []types.RadiusAttrs
	if err := api.client.CallResult(ctx, "getRadiusAttrs", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
