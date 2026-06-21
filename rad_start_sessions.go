package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// RadStartSessionsRequest — запрос на старт Radius-сессий.
type RadStartSessionsRequest []types.RadSession

func (RadStartSessionsRequest) Method() string {
	return "radStartSessions"
}

// RadStartSessions — сообщить серверу о старте Radius-сессий.
func (api *Client) RadStartSessions(ctx context.Context, request RadStartSessionsRequest) error {
	var result interface{}
	if err := api.client.CallResult(ctx, "radStartSessions", &request, &result); err != nil {
		return types.ParseJSONRPCError(err)
	}
	return nil
}
