package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getRadiusAttrsRequest = json.RawMessage(`{ "agent_id": 1, "pg_num": 1, "pg_size": 10000, "time_mark_start": 1646922115 }`)
var getRadiusAttrsResponse = json.RawMessage(`[{ "agent_descr": "Интернет", "attr_id": 187, "cat_descr": null, "cat_idx": null, "dae_on_event": 0, "description": "Скорость 5120","dev_group_id": null, "dict_name": "Mikrotik-Rate-Limit", "group_id": null, "group_name": null, "id": 1, "nas_id": null, "owner_descr": null, "radius_code": 2,"record_id": 1, "segment_id": null, "segment_ip": null, "service": null, "service_for_list": 0, "shape": 5120, "tag": 0, "tar_id": null, "tarif_descr": null, "time_mark": { "is_deleted": false, "mark": 1646922115 }, "value": "5120K", "vg_id": null, "vgroup_login": null }, { "agent_descr": "Интернет", "attr_id": 187, "cat_descr": null, "cat_idx": null, "dae_on_event": 0, "description": "Скорость 8192", "dev_group_id": null, "dict_name": "Mikrotik-Rate-Limit", "group_id": null, "group_name": null, "id": 1, "nas_id": null, "owner_descr": null, "radius_code": 2, "record_id": 2, "segment_id": null, "segment_ip": null, "service": null, "service_for_list": 0, "shape": 8192, "tag": 0, "tar_id": null, "tarif_descr": null, "time_mark": { "is_deleted": false, "mark": 1646922115 }, "value": "8192K", "vg_id": null, "vgroup_login": null }]`)

func Test_GetRadiusAttrs(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetRadiusAttrs",
				out:  mustUnmarshal[[]types.RadiusAttrs](t, getRadiusAttrsResponse),
				method: CallMethod[GetRadiusAttrsRequest, []types.RadiusAttrs](ApiCall[GetRadiusAttrsRequest, []types.RadiusAttrs]{
					req:    mustUnmarshal[GetRadiusAttrsRequest](t, getRadiusAttrsRequest),
					res:    mustUnmarshal[[]types.RadiusAttrs](t, getRadiusAttrsResponse),
					mux:    mux,
					method: api.GetRadiusAttrs,
				}),
			},
			{
				name: "call GetRadiusAttrs with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetRadiusAttrsRequest, []types.RadiusAttrs](ApiCall[GetRadiusAttrsRequest, []types.RadiusAttrs]{
					req:    mustUnmarshal[GetRadiusAttrsRequest](t, getRadiusAttrsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetRadiusAttrs,
				}),
			},
		}
	})
}
