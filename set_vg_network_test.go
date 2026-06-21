package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var setVgNetworkRequest = json.RawMessage(`{"as_num":0,"netmask":"255.255.255.255","network":"10.0.0.100","record_id":0,"segment_id":20,"type":0,"vg_id":213}`)
var setVgNetworkResponse = json.RawMessage(`512`)

func Test_SetVgNetwork(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call setVgNetwork success",
				out:  mustUnmarshal[int](t, setVgNetworkResponse),
				method: CallMethod[SetVgNetworkRequest, int](ApiCall[SetVgNetworkRequest, int]{
					req:    mustUnmarshal[SetVgNetworkRequest](t, setVgNetworkRequest),
					res:    mustUnmarshal[int](t, setVgNetworkResponse),
					mux:    mux,
					method: api.SetVgNetwork,
				}),
			},
			{
				name: "call setVgNetwork with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[SetVgNetworkRequest, int](ApiCall[SetVgNetworkRequest, int]{
					req:    mustUnmarshal[SetVgNetworkRequest](t, setVgNetworkRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.SetVgNetwork,
				}),
			},
		}
	})
}
