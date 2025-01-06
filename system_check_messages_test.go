package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var systemCheckMessagesRequest = json.RawMessage(`{}`)
var systemCheckMessagesResponse = json.RawMessage(`true`)

func Test_SystemCheckMessages(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call SystemCheckMessages",
				out:  mustUnmarshal[bool](t, systemCheckMessagesResponse),
				method: CallMethod[SystemCheckMessagesRequest, bool](ApiCall[SystemCheckMessagesRequest, bool]{
					req:    mustUnmarshal[SystemCheckMessagesRequest](t, systemCheckMessagesRequest),
					res:    mustUnmarshal[bool](t, systemCheckMessagesResponse),
					mux:    mux,
					method: api.SystemCheckMessages,
				}),
			},
			{
				name: "call SystemCheckMessages with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[SystemCheckMessagesRequest, bool](ApiCall[SystemCheckMessagesRequest, bool]{
					req:    mustUnmarshal[SystemCheckMessagesRequest](t, systemCheckMessagesRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.SystemCheckMessages,
				}),
			},
		}
	})
}
