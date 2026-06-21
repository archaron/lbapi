package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getVgNetworksRequest = json.RawMessage(`{"count":true,"japi":13,"pg_num":1,"pg_size":100,"vg_id":213}`)
var getVgNetworksResponse = json.RawMessage(`{"data":[{"as_num":null,"equip_id":null,"hwaddr":null,"netmask":"255.255.255.255","network":"10.0.0.100","record_id":512,"segment_id":20,"segment_type":0,"service_id":null,"time_mark":null,"type":0,"vg_id":213}],"total":1}`)

func Test_GetVgNetworks(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getVgNetworks success",
				out:  mustUnmarshal[*GetVgNetworksResponse](t, getVgNetworksResponse),
				method: CallMethod[GetVgNetworksRequest, *GetVgNetworksResponse](ApiCall[GetVgNetworksRequest, *GetVgNetworksResponse]{
					req:    mustUnmarshal[GetVgNetworksRequest](t, getVgNetworksRequest),
					res:    mustUnmarshal[*GetVgNetworksResponse](t, getVgNetworksResponse),
					mux:    mux,
					method: api.GetVgNetworks,
				}),
			},
			{
				name: "call getVgNetworks with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetVgNetworksRequest, *GetVgNetworksResponse](ApiCall[GetVgNetworksRequest, *GetVgNetworksResponse]{
					req:    mustUnmarshal[GetVgNetworksRequest](t, getVgNetworksRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetVgNetworks,
				}),
			},
		}
	})
}
