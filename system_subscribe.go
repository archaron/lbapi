package lbapi

import (
	"context"

	"github.com/archaron/lbapi/events"
	"github.com/archaron/lbapi/types"
)

type (
	subscribeRequest struct {
		Event events.LBEvent `json:"message"`
	}

	SystemSubscribeRequest events.LBEvent

	SystemSubscribeMultipleRequest []events.LBEvent
)

func (SystemSubscribeRequest) Method() string         { return "system_subscribe" }
func (SystemSubscribeMultipleRequest) Method() string { return "system_subscribe" }

func (api *Client) SystemSubscribeMultiple(ctx context.Context, request SystemSubscribeMultipleRequest) (bool, error) {
	var result bool
	for i := 0; i < len(request); i++ {

		if err := api.client.CallResult(ctx, "system_subscribe", subscribeRequest{Event: request[i]}, &result); err != nil {
			return false, types.ParseJSONRPCError(err)
		}
		if !result {
			return false, nil
		}
	}

	return result, nil
}

func (api *Client) SystemSubscribe(ctx context.Context, request SystemSubscribeRequest) (bool, error) {
	var result bool
	if err := api.client.CallResult(ctx, "system_subscribe", subscribeRequest{Event: events.LBEvent(request)}, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
