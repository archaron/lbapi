package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var setVgroupStateRequest = json.RawMessage(`{"id":1,"vg_id":424,"changed":0,"time_mark":1781851977}`)
var setVgroupStateResponse = json.RawMessage(`true`)

func Test_SetVgroupState(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call setVgroupState success",
				out:  mustUnmarshal[bool](t, setVgroupStateResponse),
				method: CallMethod[SetVgroupStateRequest, bool](ApiCall[SetVgroupStateRequest, bool]{
					req:    mustUnmarshal[SetVgroupStateRequest](t, setVgroupStateRequest),
					res:    mustUnmarshal[bool](t, setVgroupStateResponse),
					mux:    mux,
					method: api.SetVgroupState,
				}),
			},
			{
				name: "call setVgroupState with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[SetVgroupStateRequest, bool](ApiCall[SetVgroupStateRequest, bool]{
					req:    mustUnmarshal[SetVgroupStateRequest](t, setVgroupStateRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.SetVgroupState,
				}),
			},
		}
	})
}
