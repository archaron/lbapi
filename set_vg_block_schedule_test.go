package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var setVgBlockScheduleRequest = json.RawMessage(`{"blk_req":0,"change_time":"2026-06-21 00:00:00","vg_id":213}`)
var setVgBlockScheduleResponse = json.RawMessage(`703`)

func Test_SetVgBlockSchedule(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call setVgBlockSchedule success",
				out:  mustUnmarshal[int](t, setVgBlockScheduleResponse),
				method: CallMethod[SetVgBlockScheduleRequest, int](ApiCall[SetVgBlockScheduleRequest, int]{
					req:    mustUnmarshal[SetVgBlockScheduleRequest](t, setVgBlockScheduleRequest),
					res:    mustUnmarshal[int](t, setVgBlockScheduleResponse),
					mux:    mux,
					method: api.SetVgBlockSchedule,
				}),
			},
			{
				name: "call setVgBlockSchedule with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[SetVgBlockScheduleRequest, int](ApiCall[SetVgBlockScheduleRequest, int]{
					req:    mustUnmarshal[SetVgBlockScheduleRequest](t, setVgBlockScheduleRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.SetVgBlockSchedule,
				}),
			},
		}
	})
}
