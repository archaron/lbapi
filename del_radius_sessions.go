package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type DelRadiusSessionsRequest struct {
	//AgentID       int `json:"agent_id"`
	// TODO
}

func (DelRadiusSessionsRequest) Method() string { return "delRadiusSessions" }

// DelRadiusSessions - удалить сессии радиуса до того, как агент не синхронизируется
func (api *Client) DelRadiusSessions(ctx context.Context, request DelRadiusSessionsRequest) (interface{}, error) {
	var response interface{}
	if err := api.client.CallResult(ctx, "delRadiusSessions", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
