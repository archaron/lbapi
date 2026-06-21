package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// SetAgentActiveRequest — запрос на отметку агента как активного.
type SetAgentActiveRequest struct {
	AgentID int `json:"agent_id"` // ID агента
}

func (SetAgentActiveRequest) Method() string {
	return "setAgentActive"
}

// SetAgentActive — сообщить серверу об активности агента.
func (api *Client) SetAgentActive(ctx context.Context, request SetAgentActiveRequest) (bool, error) {
	var result bool
	if err := api.client.CallResult(ctx, "setAgentActive", &request, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
