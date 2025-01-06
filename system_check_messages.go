package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type SystemCheckMessagesRequest struct{}

func (SystemCheckMessagesRequest) Method() string {
	return "system_check_messages"
}

func (api *Client) SystemCheckMessages(ctx context.Context, request SystemCheckMessagesRequest) (bool, error) {
	var result bool
	if err := api.client.CallResult(ctx, "system_check_messages", &request, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
