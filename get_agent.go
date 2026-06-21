package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetAgentRequest — запрос информации об одном агенте.
type GetAgentRequest struct {
	AgentID int `json:"agent_id"`
}

func (GetAgentRequest) Method() string { return "getAgent" }

// GetAgent — получить полную информацию об агенте.
func (api *Client) GetAgent(ctx context.Context, agentID int) (*types.LBAgentFull, error) {
	var result types.LBAgentFull
	if err := api.client.CallResult(ctx, "getAgent", &GetAgentRequest{AgentID: agentID}, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
