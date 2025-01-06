package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getSegmentsRequest = json.RawMessage(`{ "pg_num": 1, "pg_size": 10000, "time_mark_start": 1635714394 }`)
var getSegmentsResponse = json.RawMessage(`[ { "agent_id": 1, "delegated_prefix": null, "descr": "Тестовый сегмент", "device_group_id": null, "device_group_name": null, "gateway": null, "guest": 0, "ignore_local": 0, "ip": "100.100.250.0", "mask": "255.255.255.0", "nas_id": -1, "nat": 0, "outer_vlan": null, "priority": 0, "rnas": null, "segment_id": 55, "time_mark": { "is_deleted": false, "mark": 1635714394 }, "type": 0, "unavailable": 0, "vlan_id": null, "vlan_name": null } ] `)

func Test_GetSegments(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetSegments",
				out:  mustUnmarshal[[]types.Segment](t, getSegmentsResponse),
				method: CallMethod[GetSegmentsRequest, []types.Segment](ApiCall[GetSegmentsRequest, []types.Segment]{
					req:    mustUnmarshal[GetSegmentsRequest](t, getSegmentsRequest),
					res:    mustUnmarshal[[]types.Segment](t, getSegmentsResponse),
					mux:    mux,
					method: api.GetSegments,
				}),
			},
			{
				name: "call GetSegments with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetSegmentsRequest, []types.Segment](ApiCall[GetSegmentsRequest, []types.Segment]{
					req:    mustUnmarshal[GetSegmentsRequest](t, getSegmentsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetSegments,
				}),
			},
		}
	})
}
