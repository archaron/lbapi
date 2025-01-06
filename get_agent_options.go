package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetAgentOptionsRequest struct {
	TimeMarkStart int `json:"time_mark_start"`
	types.Pagination
}

func (GetAgentOptionsRequest) Method() string {
	return "getAgentOptions"
}

func (api *Client) GetAgentOptions(ctx context.Context, request GetAgentOptionsRequest) ([]types.AgentOption, error) {
	var result []types.AgentOption
	if err := api.client.CallResult(ctx, "getAgentOptions", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
