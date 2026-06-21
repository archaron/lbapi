package lbapi

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var radStartSessionsRequest = json.RawMessage(`[{"vg_id":100,"class":"<1781822260, 100, TestRadClass>","nas":"192.168.1.100","session_id":"abc12345","parent_session_id":"","start_time":"2026-06-17 17:49:11","sess_ani":"00:11:22:33:44:55","id":1,"stop_req":0,"authmoment":"0000-00-00 00:00:00","direction":0,"updatetime":"2026-06-19 01:49:13","shape":30720,"vlan_id":0,"port_name":"","synchronized":0,"assigned_ip":"10.10.0.55","ipv4_static":0}]`)
var radStartSessionsResponse = json.RawMessage(`null`)

func Test_RadStartSessions(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call radStartSessions success",
				out:  nil,
				method: CallMethod[RadStartSessionsRequest, interface{}](ApiCall[RadStartSessionsRequest, interface{}]{
					req:    mustUnmarshal[RadStartSessionsRequest](t, radStartSessionsRequest),
					res:    nil,
					mux:    mux,
					method: func(ctx context.Context, req RadStartSessionsRequest) (interface{}, error) {
						return nil, api.RadStartSessions(ctx, req)
					},
				}),
			},
			{
				name: "call radStartSessions with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[RadStartSessionsRequest, interface{}](ApiCall[RadStartSessionsRequest, interface{}]{
					req:    mustUnmarshal[RadStartSessionsRequest](t, radStartSessionsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: func(ctx context.Context, req RadStartSessionsRequest) (interface{}, error) {
						return nil, api.RadStartSessions(ctx, req)
					},
				}),
			},
		}
	})
}
