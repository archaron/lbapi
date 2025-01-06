package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getAccountsCountRequest = json.RawMessage(`{"category":0,"count":true,"is_archive":false,"is_template":0}`)
var getAccountsCountResponse = json.RawMessage(`296`)

var getAccountsRequest = json.RawMessage(`{"category":0,"count":false,"is_archive":false,"is_template":0}`)
var getAccountsResponse = json.RawMessage(`[
	{ "addresses": [  ], "agreements": [  ], "app_id": null, "category": 0, "descr": "", "email": "test1@test.ru", "is_archive": false, "is_default": false, "is_template": 0, "login": "test1", "mobile": "+79795433262", "name": "Тест Тестович Тестовский", "phone": "+798754354354", "sole_proprietor": null, "type": 1, "uid": 2 }, 
	{ "addresses": [  ], "agreements": [  ], "app_id": null, "category": 0, "descr": "", "email": "test2@test.ru", "is_archive": false, "is_default": false, "is_template": 0, "login": "test2", "mobile": "+79897987987", "name": "Тестов Тест Тестович", "phone": "+798756456465", "sole_proprietor": null, "type": 1, "uid": 3 }
]`)

func Test_GetAccounts(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getAccounts count",
				out:  mustUnmarshal[int](t, getAccountsCountResponse),
				method: CallMethod[GetAccountsFilter, int](ApiCall[GetAccountsFilter, int]{
					req:    mustUnmarshal[GetAccountsFilter](t, getAccountsCountRequest),
					res:    mustUnmarshal[int](t, getAccountsCountResponse),
					mux:    mux,
					method: api.GetAccountsCount,
				}),
			},
			{
				name: "call getAccounts count with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetAccountsFilter, int](ApiCall[GetAccountsFilter, int]{
					req:    mustUnmarshal[GetAccountsFilter](t, getAccountsCountRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetAccountsCount,
				}),
			},
			{
				name: "call getAccounts",
				out:  mustUnmarshal[[]types.LBAccount](t, getAccountsResponse),
				method: CallMethod[GetAccountsFilter, []types.LBAccount](ApiCall[GetAccountsFilter, []types.LBAccount]{
					req:    mustUnmarshal[GetAccountsFilter](t, getAccountsRequest),
					res:    mustUnmarshal[[]types.LBAccount](t, getAccountsResponse),
					mux:    mux,
					method: api.GetAccounts,
				}),
			},
			{
				name: "call getAccounts with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetAccountsFilter, []types.LBAccount](ApiCall[GetAccountsFilter, []types.LBAccount]{
					req:    mustUnmarshal[GetAccountsFilter](t, getAccountsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetAccounts,
				}),
			},
		}
	})
}
