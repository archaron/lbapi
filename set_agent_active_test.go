package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var setAgentActiveRequest = json.RawMessage(`{"agent_id":1}`)
var setAgentActiveResponse = json.RawMessage(`true`)

func Test_SetAgentActive(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call setAgentActive success",
				out:  mustUnmarshal[bool](t, setAgentActiveResponse),
				method: CallMethod[SetAgentActiveRequest, bool](ApiCall[SetAgentActiveRequest, bool]{
					req:    mustUnmarshal[SetAgentActiveRequest](t, setAgentActiveRequest),
					res:    mustUnmarshal[bool](t, setAgentActiveResponse),
					mux:    mux,
					method: api.SetAgentActive,
				}),
			},
			{
				name: "call setAgentActive with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[SetAgentActiveRequest, bool](ApiCall[SetAgentActiveRequest, bool]{
					req:    mustUnmarshal[SetAgentActiveRequest](t, setAgentActiveRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.SetAgentActive,
				}),
			},
		}
	})
}
