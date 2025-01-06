package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getGroupsStaffRequest = json.RawMessage(`{"vg_id":213,"pg_num":0,"pg_size":0}`)
var getGroupsStaffResponse = json.RawMessage(`[ { "group_id": 1, "time_mark": null, "vg_id": 213 } ]`)

func Test_GetGroupsStaff(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetGroupsStaff",
				out:  mustUnmarshal[[]types.GroupStaff](t, getGroupsStaffResponse),
				method: CallMethod[GetGroupsStaffRequest, []types.GroupStaff](ApiCall[GetGroupsStaffRequest, []types.GroupStaff]{
					req:    mustUnmarshal[GetGroupsStaffRequest](t, getGroupsStaffRequest),
					res:    mustUnmarshal[[]types.GroupStaff](t, getGroupsStaffResponse),
					mux:    mux,
					method: api.GetGroupsStaff,
				}),
			},
			{
				name: "call GetGroupsStaff with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetGroupsStaffRequest, []types.GroupStaff](ApiCall[GetGroupsStaffRequest, []types.GroupStaff]{
					req:    mustUnmarshal[GetGroupsStaffRequest](t, getGroupsStaffRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetGroupsStaff,
				}),
			},
		}
	})
}
