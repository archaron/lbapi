package lbapi

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/archaron/lbapi/types"
)

var delVgNetworkRequest = json.RawMessage(`{"record_id":512}`)
var delVgNetworkResponse = json.RawMessage(`true`)

func Test_DelVgNetwork(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call delVgNetwork success",
				out:  mustUnmarshal[bool](t, delVgNetworkResponse),
				method: CallMethod[DelVgNetworkRequest, bool](ApiCall[DelVgNetworkRequest, bool]{
					req:    mustUnmarshal[DelVgNetworkRequest](t, delVgNetworkRequest),
					res:    mustUnmarshal[bool](t, delVgNetworkResponse),
					mux:    mux,
					method: func(ctx context.Context, req DelVgNetworkRequest) (bool, error) {
						return api.DelVgNetwork(ctx, req.RecordID)
					},
				}),
			},
			{
				name: "call delVgNetwork with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[DelVgNetworkRequest, bool](ApiCall[DelVgNetworkRequest, bool]{
					req:    mustUnmarshal[DelVgNetworkRequest](t, delVgNetworkRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: func(ctx context.Context, req DelVgNetworkRequest) (bool, error) {
						return api.DelVgNetwork(ctx, req.RecordID)
					},
				}),
			},
		}
	})
}
