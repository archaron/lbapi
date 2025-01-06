package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getTarTimeShapesRequest = json.RawMessage(`{"tar_id":45}`)
var getTarTimeShapesResponse = json.RawMessage(`[ { "id": 1, "shape_rate": 6666, "tar_id": 45, "time_from": "00:15:00", "time_mark": null, "time_to": "05:25:00", "use_weekend": false, "week": { "fri": false, "mon": true, "sat": false, "sun": false, "thu": true, "tue": true, "value": 30, "wed": true } } ]`)

func Test_GetTarTimeShape(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetTarTimeShapes",
				out:  mustUnmarshal[[]types.TarTimeShape](t, getTarTimeShapesResponse),
				method: CallMethod[GetTarTimeShapesRequest, []types.TarTimeShape](ApiCall[GetTarTimeShapesRequest, []types.TarTimeShape]{
					req:    mustUnmarshal[GetTarTimeShapesRequest](t, getTarTimeShapesRequest),
					res:    mustUnmarshal[[]types.TarTimeShape](t, getTarTimeShapesResponse),
					mux:    mux,
					method: api.GetTarTimeShapes,
				}),
			},
			{
				name: "call GetTarTimeShapes with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetTarTimeShapesRequest, []types.TarTimeShape](ApiCall[GetTarTimeShapesRequest, []types.TarTimeShape]{
					req:    mustUnmarshal[GetTarTimeShapesRequest](t, getTarTimeShapesRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetTarTimeShapes,
				}),
			},
		}
	})
}
