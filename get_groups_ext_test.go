package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getGroupsExtRequest = json.RawMessage(`{"count":true,"japi":13,"pg_num":1,"pg_size":100}`)
var getGroupsExtResponse = json.RawMessage(`{"data":[{"agent_ids":[1,2],"descr":"Тестовая группа учётных записей","group_id":1,"name":"test_group","vg_count":2}],"total":1}`)

func Test_GetGroupsExt(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getGroupsExt success",
				out:  mustUnmarshal[*GetGroupsExtResponse](t, getGroupsExtResponse),
				method: CallMethod[GetGroupsExtRequest, *GetGroupsExtResponse](ApiCall[GetGroupsExtRequest, *GetGroupsExtResponse]{
					req:    mustUnmarshal[GetGroupsExtRequest](t, getGroupsExtRequest),
					res:    mustUnmarshal[*GetGroupsExtResponse](t, getGroupsExtResponse),
					mux:    mux,
					method: api.GetGroupsExt,
				}),
			},
			{
				name: "call getGroupsExt with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetGroupsExtRequest, *GetGroupsExtResponse](ApiCall[GetGroupsExtRequest, *GetGroupsExtResponse]{
					req:    mustUnmarshal[GetGroupsExtRequest](t, getGroupsExtRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetGroupsExt,
				}),
			},
		}
	})
}
