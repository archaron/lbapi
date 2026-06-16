package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetAccountsAddonsParams struct {
		Uid int `json:"uid"`
	}

	GetAccountAddonsRequest struct {
		GetAccountsAddonsParams
	}
)

func (GetAccountAddonsRequest) Method() string {
	return "getAccountAddons"
}

// GetAccountAddons Возвращает список дополнительных полей аккаунта
func (api *Client) GetAccountAddons(ctx context.Context, uid int) ([]types.AccountAddon, error) {
	var result []types.AccountAddon
	request := GetAccountAddonsRequest{
		GetAccountsAddonsParams{
			Uid: uid,
		},
	}
	if err := api.client.CallResult(ctx, "getAccountAddons", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
