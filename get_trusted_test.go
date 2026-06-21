package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getTrustedRequest = json.RawMessage(`{"count":true,"japi":13,"pg_num":1,"pg_size":100}`)
var getTrustedResponse = json.RawMessage(`{"data":[{"descr":"localhost","ip":"127.0.0.1","mask":"255.255.255.255","record_id":1},{"descr":"Офис","ip":"192.168.1.100","mask":"255.255.255.255","record_id":2}],"total":2}`)

func Test_GetTrusted(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getTrusted success",
				out:  mustUnmarshal[*GetTrustedResponse](t, getTrustedResponse),
				method: CallMethod[GetTrustedRequest, *GetTrustedResponse](ApiCall[GetTrustedRequest, *GetTrustedResponse]{
					req:    mustUnmarshal[GetTrustedRequest](t, getTrustedRequest),
					res:    mustUnmarshal[*GetTrustedResponse](t, getTrustedResponse),
					mux:    mux,
					method: api.GetTrusted,
				}),
			},
			{
				name: "call getTrusted with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetTrustedRequest, *GetTrustedResponse](ApiCall[GetTrustedRequest, *GetTrustedResponse]{
					req:    mustUnmarshal[GetTrustedRequest](t, getTrustedRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetTrusted,
				}),
			},
		}
	})
}
