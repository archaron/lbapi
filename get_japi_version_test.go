package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var japiVersionResponse = json.RawMessage(`{ "id": 64, "jsonrpc": "2.0", "result": { "major": 20, "minor": 1, "patch": 0, "version": "20.1.0" } }`)

func Test_GetJAPIVersion(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getJAPIVersion",
				out:  mustUnmarshal[*GetJAPIVersionResponse](t, japiVersionResponse),
				method: CallMethod[GetJAPIVersionRequest, *GetJAPIVersionResponse](ApiCall[GetJAPIVersionRequest, *GetJAPIVersionResponse]{
					req:    GetJAPIVersionRequest{},
					res:    mustUnmarshal[*GetJAPIVersionResponse](t, japiVersionResponse),
					mux:    mux,
					method: APICallOneFunc[GetJAPIVersionRequest, *GetJAPIVersionResponse](api.GetJAPIVersion).call,
				}),
			},
			{
				name: "call getJAPIVersion with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetJAPIVersionRequest, *GetJAPIVersionResponse](ApiCall[GetJAPIVersionRequest, *GetJAPIVersionResponse]{
					req:    GetJAPIVersionRequest{},
					err:    types.ErrClient,
					mux:    mux,
					method: APICallOneFunc[GetJAPIVersionRequest, *GetJAPIVersionResponse](api.GetJAPIVersion).call,
				}),
			},
		}
	})
}
