package lbapi

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/archaron/lbapi/types"
)

var delAccountRequest = json.RawMessage(`{"uid":377}`)
var delAccountResponse = json.RawMessage(`true`)

func Test_DelAccount(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call delAccount success",
				out:  mustUnmarshal[bool](t, delAccountResponse),
				method: CallMethod[DelAccountRequest, bool](ApiCall[DelAccountRequest, bool]{
					req:    mustUnmarshal[DelAccountRequest](t, delAccountRequest),
					res:    mustUnmarshal[bool](t, delAccountResponse),
					mux:    mux,
					method: func(ctx context.Context, req DelAccountRequest) (bool, error) {
						return api.DelAccount(ctx, req.UID)
					},
				}),
			},
			{
				name: "call delAccount with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[DelAccountRequest, bool](ApiCall[DelAccountRequest, bool]{
					req:    mustUnmarshal[DelAccountRequest](t, delAccountRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: func(ctx context.Context, req DelAccountRequest) (bool, error) {
						return api.DelAccount(ctx, req.UID)
					},
				}),
			},
		}
	})
}
