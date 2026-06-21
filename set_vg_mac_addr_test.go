package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var setVgMacAddrRequest = json.RawMessage(`{"mac":"AA:BB:CC:DD:EE:FF","network":"10.0.0.100","vg_id":213}`)
var setVgMacAddrResponse = json.RawMessage(`412`)

func Test_SetVgMacAddr(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call setVgMacAddr success",
				out:  mustUnmarshal[int](t, setVgMacAddrResponse),
				method: CallMethod[SetVgMacAddrRequest, int](ApiCall[SetVgMacAddrRequest, int]{
					req:    mustUnmarshal[SetVgMacAddrRequest](t, setVgMacAddrRequest),
					res:    mustUnmarshal[int](t, setVgMacAddrResponse),
					mux:    mux,
					method: api.SetVgMacAddr,
				}),
			},
			{
				name: "call setVgMacAddr with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[SetVgMacAddrRequest, int](ApiCall[SetVgMacAddrRequest, int]{
					req:    mustUnmarshal[SetVgMacAddrRequest](t, setVgMacAddrRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.SetVgMacAddr,
				}),
			},
		}
	})
}
