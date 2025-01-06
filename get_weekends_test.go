package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getWeekendsRequest = json.RawMessage(`{"date_from":"2025-01-03","date_to":"2025-01-05","pg_num":0,"pg_size":0}`)
var getWeekendsResponse = json.RawMessage(`[ { "date": "2025-01-03", "is_holiday": true, "time_mark": null }, { "date": "2025-01-04", "is_holiday": true, "time_mark": null } ]`)

func TestClient_GetWeekends(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetWeekends",
				out:  mustUnmarshal[[]types.Weekend](t, getWeekendsResponse),
				method: CallMethod[GetWeekendsRequest, []types.Weekend](ApiCall[GetWeekendsRequest, []types.Weekend]{
					req:    mustUnmarshal[GetWeekendsRequest](t, getWeekendsRequest),
					res:    mustUnmarshal[[]types.Weekend](t, getWeekendsResponse),
					mux:    mux,
					method: api.GetWeekends,
				}),
			},
			{
				name: "call GetWeekends with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetWeekendsRequest, []types.Weekend](ApiCall[GetWeekendsRequest, []types.Weekend]{
					req:    mustUnmarshal[GetWeekendsRequest](t, getWeekendsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetWeekends,
				}),
			},
		}
	})
}
