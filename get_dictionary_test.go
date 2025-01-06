package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getDictionaryRequest = json.RawMessage(`{ "pg_num": 1, "pg_size": 10000, "time_mark_start": 1646922377 } `)
var getDictionaryResponse = json.RawMessage(`[ { "is_default": 1, "name": "h323-redirect-number", "nas_id": null, "radius_type": 106, "record_id": 183, "replace_id": null, "tagged": false, "time_mark": { "is_deleted": false, "mark": 1646922377 }, "to_history": false, "value_type": 1, "vendor": 9 } ] `)

func Test_GetDictionary(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetDictionary",
				out:  mustUnmarshal[[]types.Dictionary](t, getDictionaryResponse),
				method: CallMethod[GetDictionaryRequest, []types.Dictionary](ApiCall[GetDictionaryRequest, []types.Dictionary]{
					req:    mustUnmarshal[GetDictionaryRequest](t, getDictionaryRequest),
					res:    mustUnmarshal[[]types.Dictionary](t, getDictionaryResponse),
					mux:    mux,
					method: api.GetDictionary,
				}),
			},
			{
				name: "call GetDictionary with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetDictionaryRequest, []types.Dictionary](ApiCall[GetDictionaryRequest, []types.Dictionary]{
					req:    mustUnmarshal[GetDictionaryRequest](t, getDictionaryRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetDictionary,
				}),
			},
		}
	})
}
