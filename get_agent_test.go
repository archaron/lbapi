package lbapi

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getAgentResponse = json.RawMessage(`{"active":true,"agent":{"adm_notify":0,"agent_id":1,"day":"2026-06-20","descr":"Интернет","eapcertpassword":"","failed_calls":0,"flush":60,"ignorelocal":1,"keepdetail":360,"lastcontact":"2026-06-21 03:35:14","local_as_num":0,"max_radius_timeout":86400,"multiply":1024,"na_db":"billing","na_ip":"127.0.0.1","na_pass":"billing","na_username":"billing","nfhost":null,"nfhost6":null,"nfport":0,"oper_cat":0,"org_agent_id":null,"organization_id":1,"organization_name":"Default organization","raccport":1813,"rad_stop_expired":0,"raddrpool":1,"rauthport":1812,"recalc_current":"0000-00-00","recalc_date":"2020-09-01","recalc_group":0,"recalc_rent":0,"recalc_stat":0,"remulate_on_naid":0,"restart_shape":0,"save_stat_addr":1,"service_name":"Radius IPoE","session_lifetime":600,"sorm_id":null,"tel_direction_mode":0,"time_mark":null,"timer":30,"type":6,"unload_to_sorm":true,"use_rtime":0,"vendor_sorm_name":null,"voip_card_user":""},"ignore_nets":[],"interfaces":[],"sessions":[],"vgroups":[]}`)

func Test_GetAgent(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getAgent success",
				out:  mustUnmarshal[*types.LBAgentFull](t, getAgentResponse),
				method: CallMethod[GetAgentRequest, *types.LBAgentFull](ApiCall[GetAgentRequest, *types.LBAgentFull]{
					req:    mustUnmarshal[GetAgentRequest](t, json.RawMessage(`{"agent_id":1}`)),
					res:    mustUnmarshal[*types.LBAgentFull](t, getAgentResponse),
					mux:    mux,
					method: func(ctx context.Context, req GetAgentRequest) (*types.LBAgentFull, error) {
						return api.GetAgent(ctx, req.AgentID)
					},
				}),
			},
		}
	})
}
