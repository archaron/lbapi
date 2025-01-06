package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getAgentOptionsRequest = json.RawMessage(`{ "agent_id": 1, "time_mark_start": 0 }`)
var getAgentOptionsResponse = json.RawMessage(`[ { "agent_id": 1, "descr": "auth_history", "name": "auth_history", "time_mark": { "is_deleted": false, "mark": 1632760537 }, "value": "" }, { "agent_id": 1, "descr": "check-ip-when-static", "name": "check-ip-when-static", "time_mark": { "is_deleted": false, "mark": 1632760537 }, "value": "0" } ] `)

func Test_GetAgentOptions(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetAgentOptions",
				out:  mustUnmarshal[[]types.AgentOption](t, getAgentOptionsResponse),
				method: CallMethod[GetAgentOptionsRequest, []types.AgentOption](ApiCall[GetAgentOptionsRequest, []types.AgentOption]{
					req:    mustUnmarshal[GetAgentOptionsRequest](t, getAgentOptionsRequest),
					res:    mustUnmarshal[[]types.AgentOption](t, getAgentOptionsResponse),
					mux:    mux,
					method: api.GetAgentOptions,
				}),
			},
			{
				name: "call GetAgentOptions with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetAgentOptionsRequest, []types.AgentOption](ApiCall[GetAgentOptionsRequest, []types.AgentOption]{
					req:    mustUnmarshal[GetAgentOptionsRequest](t, getAgentOptionsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetAgentOptions,
				}),
			},
		}
	})
}
