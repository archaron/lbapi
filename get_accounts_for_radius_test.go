package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getAccountsForRadiusRequest = json.RawMessage(`{ "agent_id": 1, "pg_num": 1, "pg_size": 10000, "time_mark": 1646917558 }`)
var getAccountsForRadiusResponse = json.RawMessage(`[ { "archive": false, "name": "Тестов Тест Тестович", "template": 0, "time_mark": 1646917558, "uid": 216 } ] `)

func Test_GetAccountsForRadius(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetAccountsForRadius",
				out:  mustUnmarshal[[]types.AccountForRadius](t, getAccountsForRadiusResponse),
				method: CallMethod[GetAccountsForRadiusRequest, []types.AccountForRadius](ApiCall[GetAccountsForRadiusRequest, []types.AccountForRadius]{
					req:    mustUnmarshal[GetAccountsForRadiusRequest](t, getAccountsForRadiusRequest),
					res:    mustUnmarshal[[]types.AccountForRadius](t, getAccountsForRadiusResponse),
					mux:    mux,
					method: api.GetAccountsForRadius,
				}),
			},
			{
				name: "call GetAccountsForRadius with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetAccountsForRadiusRequest, []types.AccountForRadius](ApiCall[GetAccountsForRadiusRequest, []types.AccountForRadius]{
					req:    mustUnmarshal[GetAccountsForRadiusRequest](t, getAccountsForRadiusRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetAccountsForRadius,
				}),
			},
		}
	})
}
