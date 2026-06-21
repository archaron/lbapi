package lbapi

import (
	"encoding/json"
	"context"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getVgroupExtRequest = json.RawMessage(`{"vg_id":15}`)
var getVgroupExtResponse = json.RawMessage(`{"acc_off_date":"0000-00-00 00:00:00","acc_on_date":"2020-08-01 00:00:00","acc_on_date_first":"2020-08-01 00:00:00","act_block":0,"activation_code":null,"addons":[],"addresses":[],"agent_id":1,"agent_name":"Интернет","agrm_id":10,"agrm_num":"500001","amount":1000.0,"archive":0,"auth_method":null,"blk_req":null,"block_date":"2020-08-01 00:00:00","blocked":0,"c_date":"2026-06-01","changedtariffon":null,"creation_date":"2020-07-17 15:09:34","current_shape":102400,"d_limit":0,"id":1,"login":"testuser","max_sessions":999,"pass":"TestPass1","shape":0,"tar_id":11,"template_":0,"time_mark":{"is_deleted":false,"mark":1781816425},"uid":2,"vg_id":15}`)

func Test_GetVgroupExt(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getVgroupExt success",
				out:  mustUnmarshal[*types.LBVgroupExt](t, getVgroupExtResponse),
				method: CallMethod[GetVgroupExtRequest, *types.LBVgroupExt](ApiCall[GetVgroupExtRequest, *types.LBVgroupExt]{
					req:    mustUnmarshal[GetVgroupExtRequest](t, getVgroupExtRequest),
					res:    mustUnmarshal[*types.LBVgroupExt](t, getVgroupExtResponse),
					mux:    mux,
					method: func(ctx context.Context, req GetVgroupExtRequest) (*types.LBVgroupExt, error) {
						return api.GetVgroupExt(ctx, req.VgID)
					},
				}),
			},
		}
	})
}
