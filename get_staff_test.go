package lbapi

//
//import (
//	"encoding/json"
//	"testing"
//
//	"github.com/archaron/lbapi/types"
//)
//
//var getStaffRequest = json.RawMessage(``)
//
//var getStaffResponse = json.RawMessage(``)
//
//func Test_GetStaff(t *testing.T) {
//	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
//		return []APITestCase{
//			{
//				name: "call GetStaff",
//				out:  mustUnmarshal[[]types.VGroupRecord](t, getStaffResponse),
//				method: CallMethod[GetStaffRequest, []types.VGroupRecord](ApiCall[GetStaffRequest, []types.VGroupRecord]{
//					req:    mustUnmarshal[GetStaffRequest](t, getStaffRequest),
//					res:    mustUnmarshal[[]types.VGroupRecord](t, getStaffResponse),
//					mux:    mux,
//					method: api.GetStaff,
//				}),
//			},
//			{
//				name: "call GetStaff with error",
//				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
//				method: CallMethod[GetStaffRequest, []types.VGroupRecord](ApiCall[GetStaffRequest, []types.VGroupRecord]{
//					req:    mustUnmarshal[GetStaffRequest](t, getStaffRequest),
//					err:    types.ErrClient,
//					mux:    mux,
//					method: api.GetStaff,
//				}),
//			},
//		}
//	})
//}
