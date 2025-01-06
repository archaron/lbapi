package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getTariffsRequest = json.RawMessage(`{ "agent_id": 1, "pg_num": 1, "pg_size": 10000, "time_mark_start": 1646917621 }`)

var getTariffsResponse = json.RawMessage(`[
 { "act_block": 0, "additional": 0, "addons": [  ], "adm_block_rent": 150.000000, "archive": 0, "assigned": 1, "available_fl": 0, "available_ul": 1, "block_rent": 0.000000, "block_rent_duration": 180, "charge_incoming": 0, "check_active_hours": 0, "coef_high": 2.000000, "coef_low": 0.000000, "common_includes": 0, "cur_id": 1, "cur_name": "RUR", "daily_rent": 0, "descr": "Услуга доступа в сеть Интернет, 50 Мбит/сек", "descr_full": "", "dyn_route": 0, "dynamic_rent": 0, "last_mod_date": 1600166489, "link": "", "nds_above": -1, "price_plan": 0, "rent": 13000.000000, "rent_as_service": 0, "rent_multiply": 0, "sale_dictionary_id": 4, "service_type": 0, "shape": 51200, "shape_prior": 0, "symbol": "руб", "tar_id": 9, "time_mark": { "is_deleted": false, "mark": 1646917621 }, "traff_limit": 0, "traff_limit_per": 0, "traff_type": 0, "type": 1, "unavailable": 0, "use_common_includes": 0, "used": 1, "usr_block_rent": 150.000000, "uuid": "", "voip_block_local": 0 } ] `)

func Test_GetTariffs(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call GetTariffs",
				out:  mustUnmarshal[[]types.Tariff](t, getTariffsResponse),
				method: CallMethod[GetTariffsRequest, []types.Tariff](ApiCall[GetTariffsRequest, []types.Tariff]{
					req:    mustUnmarshal[GetTariffsRequest](t, getTariffsRequest),
					res:    mustUnmarshal[[]types.Tariff](t, getTariffsResponse),
					mux:    mux,
					method: api.GetTariffs,
				}),
			},
			{
				name: "call GetTariffs with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetTariffsRequest, []types.Tariff](ApiCall[GetTariffsRequest, []types.Tariff]{
					req:    mustUnmarshal[GetTariffsRequest](t, getTariffsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetTariffs,
				}),
			},
		}
	})
}
