package lbapi

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/archaron/lbapi/types"
)

var delVgMacAddrRequest = json.RawMessage(`{"record_id":412}`)
var delVgMacAddrResponse = json.RawMessage(`true`)

func Test_DelVgMacAddr(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call delVgMacAddr success",
				out:  mustUnmarshal[bool](t, delVgMacAddrResponse),
				method: CallMethod[DelVgMacAddrRequest, bool](ApiCall[DelVgMacAddrRequest, bool]{
					req:    mustUnmarshal[DelVgMacAddrRequest](t, delVgMacAddrRequest),
					res:    mustUnmarshal[bool](t, delVgMacAddrResponse),
					mux:    mux,
					method: func(ctx context.Context, req DelVgMacAddrRequest) (bool, error) {
						return api.DelVgMacAddr(ctx, req.RecordID)
					},
				}),
			},
			{
				name: "call delVgMacAddr with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[DelVgMacAddrRequest, bool](ApiCall[DelVgMacAddrRequest, bool]{
					req:    mustUnmarshal[DelVgMacAddrRequest](t, delVgMacAddrRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: func(ctx context.Context, req DelVgMacAddrRequest) (bool, error) {
						return api.DelVgMacAddr(ctx, req.RecordID)
					},
				}),
			},
		}
	})
}
