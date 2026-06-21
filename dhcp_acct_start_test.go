package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var dhcpAcctStartRequest = json.RawMessage(`{"session_id":"abc12345","nas":"192.168.1.100","assigned_ip":"10.10.0.55","class":"<1781823934, 332, TestDhcpClass>"}`)
var dhcpAcctStartResponse = json.RawMessage(`true`)

func Test_DhcpAcctStart(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call dhcpAcctStart success",
				out:  mustUnmarshal[bool](t, dhcpAcctStartResponse),
				method: CallMethod[DhcpAcctStartRequest, bool](ApiCall[DhcpAcctStartRequest, bool]{
					req:    mustUnmarshal[DhcpAcctStartRequest](t, dhcpAcctStartRequest),
					res:    mustUnmarshal[bool](t, dhcpAcctStartResponse),
					mux:    mux,
					method: api.DhcpAcctStart,
				}),
			},
			{
				name: "call dhcpAcctStart with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[DhcpAcctStartRequest, bool](ApiCall[DhcpAcctStartRequest, bool]{
					req:    mustUnmarshal[DhcpAcctStartRequest](t, dhcpAcctStartRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.DhcpAcctStart,
				}),
			},
		}
	})
}
