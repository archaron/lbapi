package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// RadStopSessionsPacketRequest — запрос на остановку Radius-сессий (пакетная обработка).
type RadStopSessionsPacketRequest []types.RadStopSession

func (RadStopSessionsPacketRequest) Method() string {
	return "radStopSessionsPacket"
}

// RadStopSessionsPacket — сообщить серверу об остановке Radius-сессий.
func (api *Client) RadStopSessionsPacket(ctx context.Context, request RadStopSessionsPacketRequest) error {
	var result interface{}
	if err := api.client.CallResult(ctx, "radStopSessionsPacket", &request, &result); err != nil {
		return types.ParseJSONRPCError(err)
	}
	return nil
}
