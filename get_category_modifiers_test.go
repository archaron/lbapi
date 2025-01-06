package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getCategoryModifiersRequest = json.RawMessage(`{ "agent_id": 1, "pg_num": 1, "pg_size": 10000, "time_mark_start": 0 } `)

var getCategoryModifiersResponse = json.RawMessage(`[]`)

func Test_GetCategoryModifiers(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetCategoryModifiersRequest",
				out:  mustUnmarshal[[]types.CategoryModifier](t, getCategoryModifiersResponse),
				method: CallMethod[GetCategoryModifiersRequest, []types.CategoryModifier](ApiCall[GetCategoryModifiersRequest, []types.CategoryModifier]{
					req:    mustUnmarshal[GetCategoryModifiersRequest](t, getCategoryModifiersRequest),
					res:    mustUnmarshal[[]types.CategoryModifier](t, getCategoryModifiersResponse),
					mux:    mux,
					method: api.GetCategoryModifiers,
				}),
			},
			{
				name: "call GetCategoryModifiersRequest with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetCategoryModifiersRequest, []types.CategoryModifier](ApiCall[GetCategoryModifiersRequest, []types.CategoryModifier]{
					req:    mustUnmarshal[GetCategoryModifiersRequest](t, getCategoryModifiersRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetCategoryModifiers,
				}),
			},
		}
	})
}
