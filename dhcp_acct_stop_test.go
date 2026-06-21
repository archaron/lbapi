package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var dhcpAcctStopRequest = json.RawMessage(`{"session_id":"abc12345","nas":"192.168.1.100","assigned_ip":"10.10.0.55","class":"<1781823934, 332, TestDhcpClass>"}`)
var dhcpAcctStopResponse = json.RawMessage(`true`)

func Test_DhcpAcctStop(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call dhcpAcctStop success",
				out:  mustUnmarshal[bool](t, dhcpAcctStopResponse),
				method: CallMethod[DhcpAcctStopRequest, bool](ApiCall[DhcpAcctStopRequest, bool]{
					req:    mustUnmarshal[DhcpAcctStopRequest](t, dhcpAcctStopRequest),
					res:    mustUnmarshal[bool](t, dhcpAcctStopResponse),
					mux:    mux,
					method: api.DhcpAcctStop,
				}),
			},
			{
				name: "call dhcpAcctStop with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[DhcpAcctStopRequest, bool](ApiCall[DhcpAcctStopRequest, bool]{
					req:    mustUnmarshal[DhcpAcctStopRequest](t, dhcpAcctStopRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.DhcpAcctStop,
				}),
			},
		}
	})
}
