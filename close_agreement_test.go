package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var closeAgreementRequest = json.RawMessage(`{"agrm_id":388,"close_date":"2026-06-19 00:00:00"}`)
var closeAgreementResponse = json.RawMessage(`true`)

func Test_CloseAgreement(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call closeAgreement success",
				out:  mustUnmarshal[bool](t, closeAgreementResponse),
				method: CallMethod[CloseAgreementRequest, bool](ApiCall[CloseAgreementRequest, bool]{
					req:    mustUnmarshal[CloseAgreementRequest](t, closeAgreementRequest),
					res:    mustUnmarshal[bool](t, closeAgreementResponse),
					mux:    mux,
					method: api.CloseAgreement,
				}),
			},
			{
				name: "call closeAgreement with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[CloseAgreementRequest, bool](ApiCall[CloseAgreementRequest, bool]{
					req:    mustUnmarshal[CloseAgreementRequest](t, closeAgreementRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.CloseAgreement,
				}),
			},
		}
	})
}
