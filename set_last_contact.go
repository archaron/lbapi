package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	SetLastContactRequest struct {
		AgentId     int    `json:"agent_id"`
		LastContact string `json:"lastcontact"`
	}
)

func (SetLastContactRequest) Method() string {
	return "setLastContact"
}

func (api *Client) SetLastContact(ctx context.Context, request SetLastContactRequest) (bool, error) {
	var result bool

	if err := api.client.CallResult(ctx, "setLastContact", request, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}

	return result, nil
}
