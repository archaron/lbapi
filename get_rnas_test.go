package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var rnasRequest = json.RawMessage(`{"agent_id":1,"time_mark_start":1647182460,"pg_num":1,"pg_size":10000}`)

var rnasResponse = json.RawMessage(`[{
    "agent_id": 1,
    "agent_name": "Интернет",
    "auth_method": "auto",
    "dae_port": 1700,
    "dae_server": "100.64.248.129",
    "descr": "BOB",
    "device_id": null,
    "device_name": null,
    "is_new": null,
    "nas_id": 1,
    "radblacklog": [],
    "rnas": "100.64.248.129",
    "secret": "bobsecret",
    "time_mark": {
        "is_deleted": false,
        "mark": 1647182460
    },
    "timezone": 0.0,
    "type": "generic"
}]`)

func Test_GetRNas(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getRnas",
				out:  mustUnmarshal[[]types.RNas](t, rnasResponse),
				method: CallMethod[GetRNasRequest, []types.RNas](ApiCall[GetRNasRequest, []types.RNas]{
					req:    mustUnmarshal[GetRNasRequest](t, rnasRequest),
					res:    mustUnmarshal[[]types.RNas](t, rnasResponse),
					mux:    mux,
					method: api.GetRNas,
				}),
			},
			{
				name: "call getRnas with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetRNasRequest, []types.RNas](ApiCall[GetRNasRequest, []types.RNas]{
					req:    mustUnmarshal[GetRNasRequest](t, rnasRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetRNas,
				}),
			},
		}
	})
}
