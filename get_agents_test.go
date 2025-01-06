package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getAgentsRequest = json.RawMessage(`{"agent_id":0,"pg_num":0,"pg_size":100}`)

var getAgentsResponse = json.RawMessage(`[{ 
		"adm_notify": 0,
		"agent_id": 1, 
		"day": "2024-12-29", 
		"descr": "Тестовый агент 1", 
		"eapcertpassword": "", 
		"failed_calls": 0, 
		"flush": 60, 
		"ignorelocal": 0, 
		"keepdetail": 360, 
		"lastcontact": "2024-12-30 03:33:03", 
		"local_as_num": 0, 
		"max_radius_timeout": 86400, 
		"multiply": 1024, 
		"na_db": "billing", 
		"na_ip": "127.0.0.1", 
		"na_pass": "billing", 
		"na_username": "billing", 
		"nfhost": null, 
		"nfhost6": null, 
		"nfport": 0, 
		"oper_cat": 0, 
		"raccport": 1813, 
		"rad_stop_expired": 0, 
		"raddrpool": 0, 
		"rauthport": 1812, 
		"recalc_current": "0000-00-00", 
		"recalc_date": "2020-09-01", 
		"recalc_group": 0, 
		"recalc_rent": 0, 
		"recalc_stat": 0, 
		"remulate_on_naid": 0, 
		"restart_shape": 0, 
		"save_stat_addr": 0, 
		"service_name": "Radius IPoE", 
		"session_lifetime": 600, 
		"tel_direction_mode": 0, 
		"time_mark": null, 
		"timer": 30, 
		"type": 6, 
		"use_rtime": 0, 
		"voip_card_user": "" 
	},
	{ 
		"adm_notify": 0, 
		"agent_id": 2, 
		"day": "2020-09-15", 
		"descr": "Тестовый агент 2", 
		"eapcertpassword": "", 
		"failed_calls": 0, 
		"flush": 60, 
		"ignorelocal": 0, 
		"keepdetail": 0, 
		"lastcontact": "0000-00-00 00:00:00", 
		"local_as_num": 0, 
		"max_radius_timeout": 86400, 
		"multiply": 1024, 
		"na_db": "billing", 
		"na_ip": "127.0.0.1", 
		"na_pass": "billing", 
		"na_username": "billing", 
		"nfhost": null, 
		"nfhost6": null, 
		"nfport": 0, 
		"oper_cat": 0, 
		"raccport": 1813, 
		"rad_stop_expired": 0, 
		"raddrpool": 0, 
		"rauthport": 1812, 
		"recalc_current": "0000-00-00", 
		"recalc_date": "0000-00-00", 
		"recalc_group": 0, 
		"recalc_rent": 0, 
		"recalc_stat": 0, 
		"remulate_on_naid": 0, 
		"restart_shape": 0, 
		"save_stat_addr": 0, 
		"service_name": "Интернет",
		"session_lifetime": 600,
		"tel_direction_mode": 0,
		"time_mark": null,
		"timer": 30,
		"type": 13,
		"use_rtime": 0,
		"voip_card_user": "" 
	}
]`)

func Test_GetAgents(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getAgents",
				out:  mustUnmarshal[[]types.LBAgent](t, getAgentsResponse),
				method: CallMethod[GetAgentsRequest, []types.LBAgent](ApiCall[GetAgentsRequest, []types.LBAgent]{
					req:    mustUnmarshal[GetAgentsRequest](t, getAgentsRequest),
					res:    mustUnmarshal[[]types.LBAgent](t, getAgentsResponse),
					mux:    mux,
					method: api.GetAgents,
				}),
			},
			{
				name: "call getAgents with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetAgentsRequest, []types.LBAgent](ApiCall[GetAgentsRequest, []types.LBAgent]{
					req:    mustUnmarshal[GetAgentsRequest](t, getAgentsRequest),
					err:    types.ErrClient,
					mux:    mux,
					method: api.GetAgents,
				}),
			},
		}
	})
}
