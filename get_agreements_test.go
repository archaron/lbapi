package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getAgreementsRequest = json.RawMessage(`{"count":true,"japi":13,"pg_num":1,"pg_size":100,"uid":370}`)
var getAgreementsResponse = json.RawMessage(`{"data":[{"addons":[],"agrm_id":378,"agrm_num":"61090019","b_check":null,"b_limit":0.0,"b_notify":0,"balance":-3000.0,"balance_acc":0.0,"balance_limit_exceeded":null,"balance_status":1,"balance_strict_limit":0.0,"block_amount":null,"block_days":null,"block_months":null,"block_orders":null,"bopos_agrm_id":"FS378","close_date":null,"create_date":"2026-01-01","credit":0.0,"cur_id":1,"curr_symbol":"руб","date_from":"2026-01-01","date_to":null,"installments":0.0,"is_closed":false,"manager_id":null,"manager_login":null,"manager_name":null,"organization_id":1,"organization_name":"Default organization","pay_class_id":0,"payment_method":0,"promise_credit":0.0,"time_mark":null,"uid":370}],"total":1}`)

func Test_GetAgreements(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getAgreements success",
				out:  mustUnmarshal[*GetAgreementsResponse](t, getAgreementsResponse),
				method: CallMethod[GetAgreementsRequest, *GetAgreementsResponse](ApiCall[GetAgreementsRequest, *GetAgreementsResponse]{
					req:    mustUnmarshal[GetAgreementsRequest](t, getAgreementsRequest),
					res:    mustUnmarshal[*GetAgreementsResponse](t, getAgreementsResponse),
					mux:    mux,
					method: api.GetAgreements,
				}),
			},
			{
				name: "call getAgreements with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetAgreementsRequest, *GetAgreementsResponse](ApiCall[GetAgreementsRequest, *GetAgreementsResponse]{
					req:    mustUnmarshal[GetAgreementsRequest](t, getAgreementsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetAgreements,
				}),
			},
		}
	})
}
