package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var setLastContactRequest = json.RawMessage(`{ "agent_id": 1, "lastcontact": "2025-01-04 23:20:41" }`)
var setLastContactResponse = json.RawMessage(`true`)

func Test_SetLastContact(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call SetLastContact",
				out:  mustUnmarshal[bool](t, setLastContactResponse),
				method: CallMethod[SetLastContactRequest, bool](ApiCall[SetLastContactRequest, bool]{
					req:    mustUnmarshal[SetLastContactRequest](t, setLastContactRequest),
					res:    mustUnmarshal[bool](t, setLastContactResponse),
					mux:    mux,
					method: api.SetLastContact,
				}),
			},
			{
				name: "call SetLastContact with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[SetLastContactRequest, bool](ApiCall[SetLastContactRequest, bool]{
					req:    mustUnmarshal[SetLastContactRequest](t, setLastContactRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.SetLastContact,
				}),
			},
		}
	})
}
