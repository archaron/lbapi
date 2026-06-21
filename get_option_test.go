package lbapi

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getOptionRequest = json.RawMessage(`{"name":"tax_value"}`)
var getOptionResponse = json.RawMessage(`{"availability_type":0,"descr":"Налоговый процент","name":"tax_value","value":"18"}`)

func Test_GetOption(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getOption success",
				out:  mustUnmarshal[*types.LBOption](t, getOptionResponse),
				method: CallMethod[GetOptionRequest, *types.LBOption](ApiCall[GetOptionRequest, *types.LBOption]{
					req:    mustUnmarshal[GetOptionRequest](t, getOptionRequest),
					res:    mustUnmarshal[*types.LBOption](t, getOptionResponse),
					mux:    mux,
					method: func(ctx context.Context, req GetOptionRequest) (*types.LBOption, error) {
						return api.GetOption(ctx, req.Name)
					},
				}),
			},
		}
	})
}
