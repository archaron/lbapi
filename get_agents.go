package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetAgentsRequest struct {
	AgentID int `json:"agent_id"`
	types.Pagination
}

func (GetAgentsRequest) Method() string {
	return "getAgents"
}

func (api *Client) GetAgents(ctx context.Context, request GetAgentsRequest) ([]types.LBAgent, error) {
	var result []types.LBAgent
	if err := api.client.CallResult(ctx, "getAgents", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
