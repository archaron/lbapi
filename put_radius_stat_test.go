package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var putRadiusStatRequest = json.RawMessage(`[{"time_from":"2025-01-05 06:48:37","time_to":"2025-01-05 07:13:37","vg_id":81,"agrm_id":83,"parent_session_id":"","ani":"b4:b0:24:08:db:0d","dnis":"dhcp-relay-agent","ip_address":"100.64.2.22","framed_prefix":"::/0","delegated_prefix":"::/0","nas_ip":"100.100.0.8","cause":0,"agent_id":1,"is_guest":false,"session_id":"cefdc083","cat_idx":0,"delta_cin":293246,"delta_cout":92463,"delta_tm":300,"service":""},{"time_from":"2024-12-09 12:48:44","time_to":"2025-01-05 07:14:16","vg_id":331,"agrm_id":302,"parent_session_id":"","ani":"00:31:92:4b:e9:9e","dnis":"dhcp-relay-agent2","ip_address":"100.100.1.7","framed_prefix":"::/0","delegated_prefix":"::/0","nas_ip":"100.100.0.8","cause":0,"agent_id":1,"is_guest":false,"session_id":"61f0c083","cat_idx":0,"delta_cin":0,"delta_cout":0,"delta_tm":300,"service":""}]`)
var putRadiusStatResponse = json.RawMessage(`true`)

func Test_PutRadiusStat(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call PutRadiusStat",
				out:  mustUnmarshal[bool](t, putRadiusStatResponse),
				method: CallMethod[PutRadiusStatRequest, bool](ApiCall[PutRadiusStatRequest, bool]{
					req:    mustUnmarshal[PutRadiusStatRequest](t, putRadiusStatRequest),
					res:    mustUnmarshal[bool](t, putRadiusStatResponse),
					mux:    mux,
					method: api.PutRadiusStat,
				}),
			},
			{
				name: "call PutRadiusStat with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[PutRadiusStatRequest, bool](ApiCall[PutRadiusStatRequest, bool]{
					req:    mustUnmarshal[PutRadiusStatRequest](t, putRadiusStatRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.PutRadiusStat,
				}),
			},
		}
	})
}
