package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getFreeNetworksCountResponse = json.RawMessage(`0`)

func Test_GetFreeNetworksCount(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getFreeNetworks count success",
				out:  mustUnmarshal[int](t, getFreeNetworksCountResponse),
				method: CallMethod[GetFreeNetworksRequest, int](ApiCall[GetFreeNetworksRequest, int]{
					req:    GetFreeNetworksRequest{AgentID: 1, Mask: 32, Segment: types.LBSegmentRef{IP: "10.0.0.0", Mask: "255.255.255.0"}, VgID: 213},
					res:    mustUnmarshal[int](t, getFreeNetworksCountResponse),
					mux:    mux,
					method: api.GetFreeNetworksCount,
				}),
			},
		}
	})
}
