package lbapi

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/archaron/lbapi/types"
)

var setAddressRequest = json.RawMessage(`{"code":"1,61,0,3254,0,293868,468,0,0,0","type":0,"uid":420}`)
var setAddressResponse = json.RawMessage(`true`)

func Test_SetAddress(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call setAddress success",
				out:  mustUnmarshal[bool](t, setAddressResponse),
				method: CallMethod[SetAddressRequest, bool](ApiCall[SetAddressRequest, bool]{
					req:    mustUnmarshal[SetAddressRequest](t, setAddressRequest),
					res:    mustUnmarshal[bool](t, setAddressResponse),
					mux:    mux,
					method: func(ctx context.Context, req SetAddressRequest) (bool, error) {
						return api.SetAddress(ctx, req.Code, req.Type, req.UID)
					},
				}),
			},
			{
				name: "call setAddress with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[SetAddressRequest, bool](ApiCall[SetAddressRequest, bool]{
					req:    mustUnmarshal[SetAddressRequest](t, setAddressRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: func(ctx context.Context, req SetAddressRequest) (bool, error) {
						return api.SetAddress(ctx, req.Code, req.Type, req.UID)
					},
				}),
			},
		}
	})
}
