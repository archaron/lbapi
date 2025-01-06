package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetAccountsFilter struct {
		Category   int  `json:"category"`
		IsArchive  bool `json:"is_archive"`
		IsTemplate int  `json:"is_template"`
	}
	GetAccountsRequest struct {
		GetAccountsFilter
		Count bool `json:"count"`
	}
)

func (GetAccountsFilter) Method() string {
	return "getAccounts"
}

func (api *Client) GetAccounts(ctx context.Context, filter GetAccountsFilter) ([]types.LBAccount, error) {
	var result []types.LBAccount
	request := GetAccountsRequest{GetAccountsFilter: filter, Count: false}
	if err := api.client.CallResult(ctx, "getAccounts", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}

func (api *Client) GetAccountsCount(ctx context.Context, filter GetAccountsFilter) (int, error) {
	var result int
	request := GetAccountsRequest{GetAccountsFilter: filter, Count: true}
	if err := api.client.CallResult(ctx, "getAccounts", &request, &result); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}
	return result, nil
}
