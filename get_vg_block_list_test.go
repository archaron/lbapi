package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getVgBlockListRequest = json.RawMessage(`{"count":true,"japi":13,"pg_num":1,"pg_size":100,"vg_id":213}`)
var getVgBlockListResponse = json.RawMessage(`{"data":[{"agrm_id":198,"agrm_num":"500001","block_type":10,"change_time":"0000-00-00 00:00:00","comment":"","cur_id":1,"curr_symbol":"руб","is_freewill":false,"is_history":true,"login":"test","manager_login":null,"mgr_descr":null,"mgr_name":null,"record_id":236,"request_by":null,"time_to":"2022-03-12 00:00:00","uid":195,"unblocked_by":0,"unblocked_manager_descr":null,"unblocked_manager_login":null,"unblocked_manager_name":null,"unblocked_time":"0000-00-00 00:00:00","vg_id":213}],"total":1}`)

func Test_GetVgBlockList(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getVgBlockList success",
				out:  mustUnmarshal[*GetVgBlockListResponse](t, getVgBlockListResponse),
				method: CallMethod[GetVgBlockListRequest, *GetVgBlockListResponse](ApiCall[GetVgBlockListRequest, *GetVgBlockListResponse]{
					req:    mustUnmarshal[GetVgBlockListRequest](t, getVgBlockListRequest),
					res:    mustUnmarshal[*GetVgBlockListResponse](t, getVgBlockListResponse),
					mux:    mux,
					method: api.GetVgBlockList,
				}),
			},
			{
				name: "call getVgBlockList with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetVgBlockListRequest, *GetVgBlockListResponse](ApiCall[GetVgBlockListRequest, *GetVgBlockListResponse]{
					req:    mustUnmarshal[GetVgBlockListRequest](t, getVgBlockListRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetVgBlockList,
				}),
			},
		}
	})
}
