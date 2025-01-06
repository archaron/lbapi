package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getGponPortsRequest = json.RawMessage(`{"port_id":281,"pg_num":0,"pg_size":0}`)
var getGponPortsResponse = json.RawMessage(`[ { "id": 2, "port_id": 281, "time_mark": { "is_deleted": false, "mark": 1735998887 }, "vg_id": 303, "virtual_id": "00:d3:9e:a8:8c:18" } ]`)

func Test_GetGPONPorts(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetGPONPorts",
				out:  mustUnmarshal[[]types.GPONPort](t, getGponPortsResponse),
				method: CallMethod[GetGPONPortsRequest, []types.GPONPort](ApiCall[GetGPONPortsRequest, []types.GPONPort]{
					req:    mustUnmarshal[GetGPONPortsRequest](t, getGponPortsRequest),
					res:    mustUnmarshal[[]types.GPONPort](t, getGponPortsResponse),
					mux:    mux,
					method: api.GetGPONPorts,
				}),
			},
			{
				name: "call GetGPONPorts with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetGPONPortsRequest, []types.GPONPort](ApiCall[GetGPONPortsRequest, []types.GPONPort]{
					req:    mustUnmarshal[GetGPONPortsRequest](t, getGponPortsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetGPONPorts,
				}),
			},
		}
	})
}
