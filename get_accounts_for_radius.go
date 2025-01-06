package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetAccountsForRadiusRequest struct {
	AgentID  int `json:"agent_id"`
	TimeMark int `json:"time_mark"`
	types.Pagination
}

func (GetAccountsForRadiusRequest) Method() string {
	return "getAccountsForRadius"
}

func (api *Client) GetAccountsForRadius(ctx context.Context, request GetAccountsForRadiusRequest) ([]types.AccountForRadius, error) {
	var result []types.AccountForRadius
	if err := api.client.CallResult(ctx, "getAccountsForRadius", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
