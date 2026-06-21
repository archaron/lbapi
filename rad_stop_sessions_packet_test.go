package lbapi

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var radStopSessionsPacketRequest = json.RawMessage(`[{"session_id":"abc12345","nas":"192.168.1.100","stop_time":"2026-06-19 02:02:09","assigned_ip":"10.10.0.55","agent_id":1,"class":"<1781779301, 461, TestRadClass>"}]`)
var radStopSessionsPacketResponse = json.RawMessage(`null`)

func Test_RadStopSessionsPacket(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call radStopSessionsPacket success",
				out:  nil,
				method: CallMethod[RadStopSessionsPacketRequest, interface{}](ApiCall[RadStopSessionsPacketRequest, interface{}]{
					req:    mustUnmarshal[RadStopSessionsPacketRequest](t, radStopSessionsPacketRequest),
					res:    nil,
					mux:    mux,
					method: func(ctx context.Context, req RadStopSessionsPacketRequest) (interface{}, error) {
						return nil, api.RadStopSessionsPacket(ctx, req)
					},
				}),
			},
			{
				name: "call radStopSessionsPacket with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[RadStopSessionsPacketRequest, interface{}](ApiCall[RadStopSessionsPacketRequest, interface{}]{
					req:    mustUnmarshal[RadStopSessionsPacketRequest](t, radStopSessionsPacketRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: func(ctx context.Context, req RadStopSessionsPacketRequest) (interface{}, error) {
						return nil, api.RadStopSessionsPacket(ctx, req)
					},
				}),
			},
		}
	})
}
