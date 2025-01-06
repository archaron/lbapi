package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getDeviceGroupsRequest = json.RawMessage(`{"time_mark_start":0,"pg_num":0,"pg_size":0}`)

var getDeviceGroupsResponse = json.RawMessage(`[ { "agent_id": null, "agent_name": null, "desc": "Тестовая группа", "group_id": 1, "name": "Тестовая", "time_mark": { "is_deleted": false, "mark": 1708789851 } } ] `)

func Test_GetDeviceGroups(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetDeviceGroups",
				out:  mustUnmarshal[[]types.LBDeviceGroup](t, getDeviceGroupsResponse),
				method: CallMethod[GetDeviceGroupsRequest, []types.LBDeviceGroup](ApiCall[GetDeviceGroupsRequest, []types.LBDeviceGroup]{
					req:    mustUnmarshal[GetDeviceGroupsRequest](t, getDeviceGroupsRequest),
					res:    mustUnmarshal[[]types.LBDeviceGroup](t, getDeviceGroupsResponse),
					mux:    mux,
					method: api.GetDeviceGroups,
				}),
			},
			{
				name: "call GetDeviceGroups with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetDeviceGroupsRequest, []types.LBDeviceGroup](ApiCall[GetDeviceGroupsRequest, []types.LBDeviceGroup]{
					req:    mustUnmarshal[GetDeviceGroupsRequest](t, getDeviceGroupsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetDeviceGroups,
				}),
			},
		}
	})
}
