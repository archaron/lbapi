package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getSyncAgreementsRequest = json.RawMessage(`{"agent_id":0,"time_mark_start":0,"pg_num":1,"pg_size":1}`)

var getSyncAgreementsResponse = json.RawMessage(`[ { "agrm_id": 1, "balance": 0.000000, "credit": 0.000000, "installments": 0.000000, "number": "1", "promise_credit": 0.000000, "time_mark": { "is_deleted": false, "mark": 1600681217 } } ] `)

func Test_GetSyncAgreements(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetSyncAgreements",
				out:  mustUnmarshal[[]types.SyncAgreement](t, getSyncAgreementsResponse),
				method: CallMethod[GetSyncAgreementsRequest, []types.SyncAgreement](ApiCall[GetSyncAgreementsRequest, []types.SyncAgreement]{
					req:    mustUnmarshal[GetSyncAgreementsRequest](t, getSyncAgreementsRequest),
					res:    mustUnmarshal[[]types.SyncAgreement](t, getSyncAgreementsResponse),
					mux:    mux,
					method: api.GetSyncAgreements,
				}),
			},
			{
				name: "call GetSyncAgreements with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetSyncAgreementsRequest, []types.SyncAgreement](ApiCall[GetSyncAgreementsRequest, []types.SyncAgreement]{
					req:    mustUnmarshal[GetSyncAgreementsRequest](t, getSyncAgreementsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetSyncAgreements,
				}),
			},
		}
	})
}
