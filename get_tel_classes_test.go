package lbapi

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getTelClassesResponse = json.RawMessage(`[{"class_id":0,"color":14286828,"descr":"Внутренняя связь","extern_code":"--","mgmn":false,"name":"ВТС"},{"class_id":1,"color":14024703,"descr":"Междугородняя связь","extern_code":"11","mgmn":true,"name":"МГ"},{"class_id":2,"color":13816575,"descr":"Международная связь","extern_code":"12","mgmn":true,"name":"МН"}]`)

func Test_GetTelClasses(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getTelClasses success",
				out:  mustUnmarshal[[]types.LBTelClass](t, getTelClassesResponse),
				method: CallMethod[GetTelClassesRequest, []types.LBTelClass](ApiCall[GetTelClassesRequest, []types.LBTelClass]{
					req:    GetTelClassesRequest{},
					res:    mustUnmarshal[[]types.LBTelClass](t, getTelClassesResponse),
					mux:    mux,
					method: func(ctx context.Context, _ GetTelClassesRequest) ([]types.LBTelClass, error) {
						return api.GetTelClasses(ctx)
					},
				}),
			},
			{
				name: "call getTelClasses with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetTelClassesRequest, []types.LBTelClass](ApiCall[GetTelClassesRequest, []types.LBTelClass]{
					req:    GetTelClassesRequest{},
					err:    types.ErrClient,
					mux:    mux,
					method: func(ctx context.Context, _ GetTelClassesRequest) ([]types.LBTelClass, error) {
						return api.GetTelClasses(ctx)
					},
				}),
			},
		}
	})
}
