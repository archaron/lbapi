package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetCurrenciesRequest — запрос списка валют.
type GetCurrenciesRequest struct{}

func (GetCurrenciesRequest) Method() string { return "getCurrencies" }

// GetCurrencies — получить список валют.
func (api *Client) GetCurrencies(ctx context.Context) ([]types.LBCurrency, error) {
	var result []types.LBCurrency
	if err := api.client.CallResult(ctx, "getCurrencies", &GetCurrenciesRequest{}, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
