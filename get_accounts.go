package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetAccountsFilter struct {
		Category            int    `json:"category"`
		IsArchive           int    `json:"is_archive"`
		IsTemplate          int    `json:"is_template"`
		GetFull             bool   `json:"get_full"`
		IncludePreactivated bool   `json:"include_preactivated"`
		AgrmNum             string `json:"agrm_num,omitempty"`
		ManagerID           int    `json:"manager_id,omitempty"`

		Inn string `json:"inn,omitempty"`
		Name string `json:"name,omitempty"`
		// Addons              GetAccountsFilterAddon `json:"addons,omitempty"`
		IsDefault int `json:"is_default,omitempty"`
		types.Pagination
	}
	// { "id": 1, "method": "getAccounts", "params": { "category": -1, "count": true, "get_full": true, "include_preactivated": false, "inn": "7703191457", "is_archive": 0, "is_template": 0, "japi": 13, "manager_id": -2, "pg_num": 1, "pg_size": 100 } }
	GetAccountsFilterAddon struct {
		StrValue string `json:"str_value"`
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
	result := make([]types.LBAccount, 0)
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
