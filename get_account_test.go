package lbapi

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/archaron/lbapi/types"
)

var getAccountRequest = json.RawMessage(`{"uid":420}`)
var getAccountResponse = json.RawMessage(`{"abonent_name":"Иванов","abonent_patronymic":"Иванович","abonent_surname":"Иван","act_on_what":"устава","addons":[],"addresses":[{"address":"Россия,обл Московская,,г Подольск,,ул Парковая,дом 10,,,,","building_uuid":null,"code":"1,50,0,1000,0,200000,300,0,0,0","type":0}],"agreements":[],"app_id":null,"bank_name":"ПАО «Тестовый Банк»","bik":"040000000","bill_delivery":5,"bill_delivery_name":"Email","birth_date":"0000-00-00","birth_place":"","branch_bank_name":"","category":0,"corr_account":"30101810000000000000","descr":"","email":"test@example.com","inn":"","is_archive":false,"is_default":false,"is_template":0,"kpp":"","login":"testuser","mobile":"+79991234567","name":"Тестовый Пользователь","ogrn":"","okpo":"","org_uid":null,"organization_id":1,"organization_name":"Default organization","passport_date":"0000-00-00","passport_num":"","passport_org":"","payment_order_num":"","phone":"+79991234567","rs":"40702810000000012345","sole_proprietor":null,"type":1,"uid":420}`)

func Test_GetAccount(t *testing.T) {
	runAPITests(t, func(api *Client, mux *APIMux) []APITestCase {
		return []APITestCase{
			{
				name: "call getAccount success",
				out:  mustUnmarshal[*types.LBAccountFull](t, getAccountResponse),
				method: CallMethod[GetAccountRequest, *types.LBAccountFull](ApiCall[GetAccountRequest, *types.LBAccountFull]{
					req:    mustUnmarshal[GetAccountRequest](t, getAccountRequest),
					res:    mustUnmarshal[*types.LBAccountFull](t, getAccountResponse),
					mux:    mux,
					method: func(ctx context.Context, req GetAccountRequest) (*types.LBAccountFull, error) {
						return api.GetAccount(ctx, req.UID)
					},
				}),
			},
		}
	})
}
