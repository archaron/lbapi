package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getRnasOptionsRequest = json.RawMessage(`{ "agent_id": 1, "pg_num": 1, "pg_size": 10000, "time_mark_start": 1600203135}`)
var getRnasOptionsResponse = json.RawMessage(` [ { "name": "interim_update_time", "nas_id": 1, "option_id": 1, "time_mark": { "is_deleted": false, "mark": 1600203135 }, "value": "300" }, { "name": "session_lifetime", "nas_id": 1, "option_id": 2, "time_mark": { "is_deleted": false, "mark": 1600203135 }, "value": "86400" }, { "name": "max_radius_timeout", "nas_id": 1, "option_id": 3, "time_mark": { "is_deleted": false, "mark": 1600203135 }, "value": "86400" } ]`)

func Test_GetRNasOptions(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetRNASOptions",
				out:  mustUnmarshal[[]types.LBRNasOption](t, getRnasOptionsResponse),
				method: CallMethod[GetRNasOptionsRequest, []types.LBRNasOption](ApiCall[GetRNasOptionsRequest, []types.LBRNasOption]{
					req:    mustUnmarshal[GetRNasOptionsRequest](t, getRnasOptionsRequest),
					res:    mustUnmarshal[[]types.LBRNasOption](t, getRnasOptionsResponse),
					mux:    mux,
					method: api.GetRNasOptions,
				}),
			},
			{
				name: "call GetRNASOptions with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetRNasOptionsRequest, []types.LBRNasOption](ApiCall[GetRNasOptionsRequest, []types.LBRNasOption]{
					req:    mustUnmarshal[GetRNasOptionsRequest](t, getRnasOptionsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetRNasOptions,
				}),
			},
		}
	})
}
