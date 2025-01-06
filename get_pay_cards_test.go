package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getPayCardsRequest = json.RawMessage(`{"is_activated":false,"set_id":0,"ser_no":"","card_key":"","create_date":null,"activate_date":null,"pg_num":0,"pg_size":0}`)

var getPayCardsResponse = json.RawMessage(`[{"act_til": "2025-01-04 00:00:00", "activate_date": null, "agrm_id": null, "agrm_number": null, "card_key": "10364713", "create_date": "2025-01-04 01:20:55", "cur_id": 1, "expire_period": 0, "mod_person": 0, "mod_person_name": "Администратор", "ser_no": 2, "set_id": 1, "set_name": "Test set", "sum": 0.000000, "symbol": "руб", "time_mark": null, "uid": null, "used": 0, "user_name": null}]`)

func Test_GetPayCards(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetPayCards",
				out:  mustUnmarshal[[]types.PayCard](t, getPayCardsResponse),
				method: CallMethod[GetPayCardsRequest, []types.PayCard](ApiCall[GetPayCardsRequest, []types.PayCard]{
					req:    mustUnmarshal[GetPayCardsRequest](t, getPayCardsRequest),
					res:    mustUnmarshal[[]types.PayCard](t, getPayCardsResponse),
					mux:    mux,
					method: api.GetPayCards,
				}),
			},
			{
				name: "call GetPayCards with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetPayCardsRequest, []types.PayCard](ApiCall[GetPayCardsRequest, []types.PayCard]{
					req:    mustUnmarshal[GetPayCardsRequest](t, getPayCardsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetPayCards,
				}),
			},
		}
	})
}
