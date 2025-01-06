package lbapi

import (
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var licenseResponse = json.RawMessage(`{
    "cdkey": "1111-2222-3333-4444-5555",
	"expire": 0,
	"gen_date": "11.03.2022 16:08:09",
	"mono_captive_portals": 10,
	"multi_captive_portals": 10,
	"payments": 20,
	"prog_name": "LANBilling 2.0 Base",
	"prog_owner": "ООО «Рога унд копыта»",
	"serial_number": "LB202017.02101",
	"use_asterisk": true,
	"use_bank_details": true,
	"use_cerber": true,
	"use_ec": true,
	"use_fidelio": true,
	"use_fiscalization": true,
	"use_inet": 99,
	"use_inventory": true,
	"use_iptvportal": true,
	"use_lbtv": true,
	"use_lifestream": true,
	"use_megogo": true,
	"use_ministra": true,
	"use_moovi": true,
	"use_nexttv": true,
	"use_operators": true,
	"use_phone": 99,
	"use_softline": true,
	"use_starcor": true,
	"use_tv24": true,
	"use_tvip": true,
	"use_usbox": 99,
	"user_limit": 99999 
}`)

func Test_GetLicense(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getLicense",
				out:  mustUnmarshal[*types.LBLicense](t, licenseResponse),
				method: CallMethod[GetLicenseRequest, *types.LBLicense](ApiCall[GetLicenseRequest, *types.LBLicense]{
					req:    GetLicenseRequest{},
					res:    mustUnmarshal[*types.LBLicense](t, licenseResponse),
					mux:    mux,
					method: APICallOneFunc[GetLicenseRequest, *types.LBLicense](api.GetLicense).call,
				}),
			},
			{
				name: "call getLicense with error",
				err:  types.NewComplexErrorf("JSON-RPC error: -32098, message: API client error"),
				method: CallMethod[GetLicenseRequest, *types.LBLicense](ApiCall[GetLicenseRequest, *types.LBLicense]{
					req:    GetLicenseRequest{},
					err:    types.ErrClient,
					mux:    mux,
					method: APICallOneFunc[GetLicenseRequest, *types.LBLicense](api.GetLicense).call,
				}),
			},
		}
	})
}
