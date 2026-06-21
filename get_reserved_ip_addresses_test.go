package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getReservedIpAddressesRequest = json.RawMessage(`{"agent_id":1,"pg_num":1,"pg_size":10000}`)
var getReservedIpAddressesResponse = json.RawMessage(`[{"ip":"10.10.0.55","vg_id":100,"mac":"00:11:22:33:44:55","netmask":"255.255.255.0","segment_id":1,"time_mark":{"is_deleted":false,"mark":1781851977}}]`)

func Test_GetReservedIpAddresses(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getReservedIpAddresses success",
				out:  mustUnmarshal[[]types.LBReservedIP](t, getReservedIpAddressesResponse),
				method: CallMethod[GetReservedIpAddressesRequest, []types.LBReservedIP](ApiCall[GetReservedIpAddressesRequest, []types.LBReservedIP]{
					req:    mustUnmarshal[GetReservedIpAddressesRequest](t, getReservedIpAddressesRequest),
					res:    mustUnmarshal[[]types.LBReservedIP](t, getReservedIpAddressesResponse),
					mux:    mux,
					method: api.GetReservedIpAddresses,
				}),
			},
			{
				name: "call getReservedIpAddresses with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetReservedIpAddressesRequest, []types.LBReservedIP](ApiCall[GetReservedIpAddressesRequest, []types.LBReservedIP]{
					req:    mustUnmarshal[GetReservedIpAddressesRequest](t, getReservedIpAddressesRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetReservedIpAddresses,
				}),
			},
		}
	})
}
