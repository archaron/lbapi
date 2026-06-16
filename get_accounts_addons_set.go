package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetAccountsAddonsSetSort struct {
		Ascdesc int    `json:"ascdesc"`
		Name    string `json:"name"`
	}

	GetAccountsAddonsSetParams struct {
		GetFull  bool                     `json:"get_full"` // Получить вместе со значениями выбора
		ShowOnHp bool                     `json:"show_on_hp"`
		Sort     GetAccountsAddonsSetSort `json:"sort"`
	}

	GetAccountsAddonsSetRequest struct {
		GetAccountsAddonsSetParams
	}
)

func (GetAccountsAddonsSetRequest) Method() string {
	return "getAccountsAddonsSet"
}

// GetAccountsAddonsSet Возвращает список дополнительных полей пользователя
func (api *Client) GetAccountsAddonsSet(ctx context.Context, params GetAccountsAddonsSetParams) ([]types.AccountsAddonsSet, error) {
	var result []types.AccountsAddonsSet
	request := GetAccountsAddonsSetRequest{GetAccountsAddonsSetParams: params}
	if err := api.client.CallResult(ctx, "getAccountsAddonsSet", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
