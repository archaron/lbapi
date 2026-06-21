package lbapi

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getStatusResponse = json.RawMessage(`{"accounts":412,"agents_stop":0,"agents_work":2,"agreements":422,"build_date":"51.0 - 20250506","build_name":"LANBilling 2.0 Base","build_rev":"9d2369b1","crm_docs":0,"currencies":1,"db_build_date":"2.0-51.0-20250506","db_build_rev":"9d2369b1","devices":4,"managers":7,"operators":1,"pay_managers":3,"tickets":1,"vgroups":451}`)

func Test_GetStatus(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getStatus success",
				out:  mustUnmarshal[types.LBStatus](t, getStatusResponse),
				method: CallMethod[GetStatusRequest, types.LBStatus](ApiCall[GetStatusRequest, types.LBStatus]{
					req:    GetStatusRequest{},
					res:    mustUnmarshal[types.LBStatus](t, getStatusResponse),
					mux:    mux,
					method: func(ctx context.Context, _ GetStatusRequest) (types.LBStatus, error) {
						s, err := api.GetStatus(ctx)
						if err != nil { return types.LBStatus{}, err }
						return *s, nil
					},
				}),
			},
			{
				name: "call getStatus with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetStatusRequest, types.LBStatus](ApiCall[GetStatusRequest, types.LBStatus]{
					req:    GetStatusRequest{},
					err:    types.ErrClient,
					mux:    mux,
					method: func(ctx context.Context, _ GetStatusRequest) (types.LBStatus, error) {
						s, err := api.GetStatus(ctx)
						if err != nil { return types.LBStatus{}, err }
						return *s, nil
					},
				}),
			},
		}
	})
}
