package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getSyncAgreementsExtRequest = json.RawMessage(`{"agent_id":0,"time_mark_start":0,"pg_num":1,"pg_size":1}`)

var getSyncAgreementsExtResponse = json.RawMessage(`[ { "agrm_id": 1, "payment_method": 2, "time_mark": { "is_deleted": false, "mark": 1563317248 } } ]`)

func Test_GetSyncAgreementsExt(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetSyncAgreementsExt",
				out:  mustUnmarshal[[]types.SyncAgreementExt](t, getSyncAgreementsExtResponse),
				method: CallMethod[GetSyncAgreementsExtRequest, []types.SyncAgreementExt](ApiCall[GetSyncAgreementsExtRequest, []types.SyncAgreementExt]{
					req:    mustUnmarshal[GetSyncAgreementsExtRequest](t, getSyncAgreementsExtRequest),
					res:    mustUnmarshal[[]types.SyncAgreementExt](t, getSyncAgreementsExtResponse),
					mux:    mux,
					method: api.GetSyncAgreementsExt,
				}),
			},
			{
				name: "call GetSyncAgreementsExt with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetSyncAgreementsExtRequest, []types.SyncAgreementExt](ApiCall[GetSyncAgreementsExtRequest, []types.SyncAgreementExt]{
					req:    mustUnmarshal[GetSyncAgreementsExtRequest](t, getSyncAgreementsExtRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetSyncAgreementsExt,
				}),
			},
		}
	})
}
