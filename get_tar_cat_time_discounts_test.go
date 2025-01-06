package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getTarCatTimeDiscountsRequest = json.RawMessage(`{"time_mark_start":0,"pg_num":0,"pg_size":0}`)
var getTarCatTimeDiscountsResponse = json.RawMessage(`[ { "cat_idx": 1, "dis_id": 1, "discount": 5.000000, "tar_id": 45, "time_from": "00:00:00", "time_mark": { "is_deleted": false, "mark": 1736010049 }, "time_to": "00:10:00", "type": 0, "use_weekend": false, "week": { "fri": false, "mon": false, "sat": true, "sun": true, "thu": false, "tue": false, "value": 65, "wed": false } }, { "cat_idx": 1, "dis_id": 2, "discount": 50.000000, "tar_id": 45, "time_from": "00:30:00", "time_mark": { "is_deleted": false, "mark": 1736010067 }, "time_to": "00:40:00", "type": 1, "use_weekend": false, "week": { "fri": false, "mon": false, "sat": true, "sun": true, "thu": false, "tue": false, "value": 65, "wed": false } } ]`)

func Test_GetTarCatTimeDiscounts(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetTarCatTimeDiscounts",
				out:  mustUnmarshal[[]types.TarCatTimeDiscount](t, getTarCatTimeDiscountsResponse),
				method: CallMethod[GetTarCatTimeDiscountsRequest, []types.TarCatTimeDiscount](ApiCall[GetTarCatTimeDiscountsRequest, []types.TarCatTimeDiscount]{
					req:    mustUnmarshal[GetTarCatTimeDiscountsRequest](t, getTarCatTimeDiscountsRequest),
					res:    mustUnmarshal[[]types.TarCatTimeDiscount](t, getTarCatTimeDiscountsResponse),
					mux:    mux,
					method: api.GetTarCatTimeDiscounts,
				}),
			},
			{
				name: "call GetTarCatTimeDiscounts with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetTarCatTimeDiscountsRequest, []types.TarCatTimeDiscount](ApiCall[GetTarCatTimeDiscountsRequest, []types.TarCatTimeDiscount]{
					req:    mustUnmarshal[GetTarCatTimeDiscountsRequest](t, getTarCatTimeDiscountsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetTarCatTimeDiscounts,
				}),
			},
		}
	})
}
