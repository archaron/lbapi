package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getSyncEquipmentRequest = json.RawMessage(`{"agent_id":1,"time_mark_start":1757537263,"pg_num":1,"pg_size":10000}`)
var getSyncEquipmentResponse = json.RawMessage(`[{"equip_id":5,"mac":"00:11:22:33:44:55","name":"Тестовое устройство","serial":"SN-001234","serial_format":0,"time_mark":{"is_deleted":false,"mark":1757537263},"vg_id":213}]`)

func Test_GetSyncEquipment(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getSyncEquipment success",
				out:  mustUnmarshal[[]types.LBSyncEquipment](t, getSyncEquipmentResponse),
				method: CallMethod[GetSyncEquipmentRequest, []types.LBSyncEquipment](ApiCall[GetSyncEquipmentRequest, []types.LBSyncEquipment]{
					req:    mustUnmarshal[GetSyncEquipmentRequest](t, getSyncEquipmentRequest),
					res:    mustUnmarshal[[]types.LBSyncEquipment](t, getSyncEquipmentResponse),
					mux:    mux,
					method: api.GetSyncEquipment,
				}),
			},
			{
				name: "call getSyncEquipment with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetSyncEquipmentRequest, []types.LBSyncEquipment](ApiCall[GetSyncEquipmentRequest, []types.LBSyncEquipment]{
					req:    mustUnmarshal[GetSyncEquipmentRequest](t, getSyncEquipmentRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetSyncEquipment,
				}),
			},
		}
	})
}
