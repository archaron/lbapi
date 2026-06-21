package lbapi

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getPassTemplatesResponse = json.RawMessage(`[{"check_pass":0,"descr":"Пароль пользователя","excluded":"","min_length":8,"need_low":0,"need_numbers":0,"need_special":0,"need_upper":0,"object_type":"account","record_id":2}]`)

func Test_GetPassTemplates(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getPassTemplates success",
				out:  mustUnmarshal[[]types.LBPassTemplate](t, getPassTemplatesResponse),
				method: CallMethod[GetPassTemplatesRequest, []types.LBPassTemplate](ApiCall[GetPassTemplatesRequest, []types.LBPassTemplate]{
					req:    mustUnmarshal[GetPassTemplatesRequest](t, json.RawMessage(`{"object_type":"account"}`)),
					res:    mustUnmarshal[[]types.LBPassTemplate](t, getPassTemplatesResponse),
					mux:    mux,
					method: func(ctx context.Context, req GetPassTemplatesRequest) ([]types.LBPassTemplate, error) {
						return api.GetPassTemplates(ctx, req.ObjectType)
					},
				}),
			},
		}
	})
}
