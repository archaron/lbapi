package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/events"
	"github.com/archaron/lbapi/types"
)

var systemSubscribeRequest = json.RawMessage(`{ "id": 2, "method": "system_subscribe", "params": { "message": "create_agreement" } }`)
var systemSubscribeResponse = json.RawMessage(`{ "id": 2, "jsonrpc": "2.0", "result": true }`)

func Test_SystemSubscribe(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call systemSubscribe",
				out:  true,
				method: CallMethod[SystemSubscribeRequest, bool](ApiCall[SystemSubscribeRequest, bool]{
					req:    events.BlockVgEvent,
					res:    true,
					mux:    mux,
					method: api.SystemSubscribe,
				}),
			},
			{
				name: "call systemSubscribe with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[SystemSubscribeRequest, bool](ApiCall[SystemSubscribeRequest, bool]{
					req:    events.BlockVgEvent,
					err:    types.ErrClient,
					mux:    mux,
					method: api.SystemSubscribe,
				}),
			},
			{
				name: "call systemSubscribe multiple",
				out:  true,
				method: CallMethod[SystemSubscribeMultipleRequest, bool](ApiCall[SystemSubscribeMultipleRequest, bool]{
					req:    SystemSubscribeMultipleRequest{events.BlockVgEvent, events.ChangeAgentEvent},
					res:    true,
					mux:    mux,
					method: api.SystemSubscribeMultiple,
				}),
			},
			{
				name: "call systemSubscribe multiple with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[SystemSubscribeMultipleRequest, bool](ApiCall[SystemSubscribeMultipleRequest, bool]{
					req:    SystemSubscribeMultipleRequest{events.BlockVgEvent, events.ChangeAgentEvent},
					err:    types.ErrClient,
					mux:    mux,
					method: api.SystemSubscribeMultiple,
				}),
			},
			{
				name: "call systemSubscribe multiple with unsuccessful subscribe",
				out:  false,
				method: CallMethod[SystemSubscribeMultipleRequest, bool](ApiCall[SystemSubscribeMultipleRequest, bool]{
					req:    SystemSubscribeMultipleRequest{events.BlockVgEvent, events.ChangeAgentEvent},
					mux:    mux,
					method: api.SystemSubscribeMultiple,
				}),
			},
		}
	})
}
