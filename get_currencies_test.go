package lbapi

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getCurrenciesResponse = json.RawMessage(`[{"code_okv":null,"cur_id":0,"is_def":false,"name":"Расчетная единица","symbol":"р.е."},{"code_okv":643,"cur_id":1,"is_def":true,"name":"RUR","symbol":"руб"}]`)

func Test_GetCurrencies(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getCurrencies success",
				out:  mustUnmarshal[[]types.LBCurrency](t, getCurrenciesResponse),
				method: CallMethod[GetCurrenciesRequest, []types.LBCurrency](ApiCall[GetCurrenciesRequest, []types.LBCurrency]{
					req:    GetCurrenciesRequest{},
					res:    mustUnmarshal[[]types.LBCurrency](t, getCurrenciesResponse),
					mux:    mux,
					method: func(ctx context.Context, _ GetCurrenciesRequest) ([]types.LBCurrency, error) {
						return api.GetCurrencies(ctx)
					},
				}),
			},
			{
				name: "call getCurrencies with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetCurrenciesRequest, []types.LBCurrency](ApiCall[GetCurrenciesRequest, []types.LBCurrency]{
					req:    GetCurrenciesRequest{},
					err:    types.ErrClient,
					mux:    mux,
					method: func(ctx context.Context, _ GetCurrenciesRequest) ([]types.LBCurrency, error) {
						return api.GetCurrencies(ctx)
					},
				}),
			},
		}
	})
}
