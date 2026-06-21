package lbapi

import (
	"encoding/json"
	"context"
	"testing"
)

var getPromisePaymentsRequest = json.RawMessage(`{"agrm_id":378,"count":true,"nodata":true,"payed":3}`)
var getPromisePaymentsResponse = json.RawMessage(`0`)

func Test_GetPromisePaymentsCount(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getPromisePayments count success",
				out:  mustUnmarshal[int](t, getPromisePaymentsResponse),
				method: CallMethod[GetPromisePaymentsRequest, int](ApiCall[GetPromisePaymentsRequest, int]{
					req:    mustUnmarshal[GetPromisePaymentsRequest](t, getPromisePaymentsRequest),
					res:    mustUnmarshal[int](t, getPromisePaymentsResponse),
					mux:    mux,
					method: func(ctx context.Context, req GetPromisePaymentsRequest) (int, error) {
						return api.GetPromisePaymentsCount(ctx, req.AgrmID)
					},
				}),
			},
		}
	})
}
