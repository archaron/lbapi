package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getBillDeliveryMethodsRequest = json.RawMessage(`{"japi":13}`)
var getBillDeliveryMethodsResponse = json.RawMessage(`{"data":[{"is_editable":false,"name":"Курьер","record_id":0},{"is_editable":false,"name":"Почтой","record_id":1},{"is_editable":false,"name":"Email","record_id":4}],"total":null}`)

func Test_GetBillDeliveryMethods(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getBillDeliveryMethods success",
				out:  mustUnmarshal[*GetBillDeliveryMethodsResponse](t, getBillDeliveryMethodsResponse),
				method: CallMethod[GetBillDeliveryMethodsRequest, *GetBillDeliveryMethodsResponse](ApiCall[GetBillDeliveryMethodsRequest, *GetBillDeliveryMethodsResponse]{
					req:    mustUnmarshal[GetBillDeliveryMethodsRequest](t, getBillDeliveryMethodsRequest),
					res:    mustUnmarshal[*GetBillDeliveryMethodsResponse](t, getBillDeliveryMethodsResponse),
					mux:    mux,
					method: api.GetBillDeliveryMethods,
				}),
			},
			{
				name: "call getBillDeliveryMethods with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetBillDeliveryMethodsRequest, *GetBillDeliveryMethodsResponse](ApiCall[GetBillDeliveryMethodsRequest, *GetBillDeliveryMethodsResponse]{
					req:    mustUnmarshal[GetBillDeliveryMethodsRequest](t, getBillDeliveryMethodsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetBillDeliveryMethods,
				}),
			},
		}
	})
}
