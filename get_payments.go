package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetPaymentsRequestParams struct {
		AccType           int           `json:"acc_type,omitempty"`
		AgentId           int           `json:"agent_id,omitempty"`
		AgrmId            int           `json:"agrm_id,omitempty"`
		AgrmNum           string        `json:"agrm_num,omitempty"`
		BankAccountIncome string        `json:"bank_account_income,omitempty"`
		ClassId           int           `json:"class_id,omitempty"`
		Email             string        `json:"email,omitempty"`
		FromAgrmId        int           `json:"from_agrm_id,omitempty"`
		GroupId           int           `json:"group_id,omitempty"`
		HasRegistry       int           `json:"has_registry,omitempty"`
		LastPayment       bool          `json:"last_payment,omitempty"`
		Login             string        `json:"login,omitempty"`
		ModPerson         int           `json:"mod_person,omitempty"`
		Nodata            bool          `json:"nodata,omitempty"`
		Nodetails         bool          `json:"nodetails,omitempty"`
		OperId            int           `json:"oper_id,omitempty"`
		PayDateFrom       string        `json:"pay_date_from,omitempty"`
		PayDateTo         string        `json:"pay_date_to,omitempty"`
		PayHistory        int           `json:"pay_history,omitempty"`
		PgNum             int           `json:"pg_num,omitempty"`
		PgSize            int           `json:"pg_size,omitempty"`
		Phone             string        `json:"phone,omitempty"`
		Receipt           string        `json:"receipt,omitempty"`
		RecordId          int           `json:"record_id,omitempty"`
		SearchTemplate    []interface{} `json:"search_template,omitempty"`
		Sort              []interface{} `json:"sort,omitempty"`
		Uid               int           `json:"uid,omitempty"`
		UserName          string        `json:"user_name,omitempty"`
		UsrGroupId        int           `json:"usr_group_id,omitempty"`
		Uuid              string        `json:"uuid,omitempty"`
		VgId              int           `json:"vg_id,omitempty"`
		VgLogin           string        `json:"vg_login,omitempty"`
	}

	GetPaymentsRequest struct {
		GetPaymentsRequestParams
	}
)

func (GetPaymentsRequest) Method() string {
	return "getPayments"
}

// GetPayments Возвращает список дополнительных полей пользователя
func (api *Client) GetPayments(ctx context.Context, params GetPaymentsRequestParams) ([]types.Payment, error) {
	var result []types.Payment
	request := GetPaymentsRequest{GetPaymentsRequestParams: params}
	if err := api.client.CallResult(ctx, "getPayments", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}

// record_id,agrm_id,amount,comment,receipt,pay_date,local_date,cancel_date,period_date,status,mod_person,order_id,class_id,agrm_number,amount_cur,amount_cur_id,from_agrm_id,bso_id,uuid,script_executed,order_number,card_number,cash_code,parent_record_id,organization_id,org_payment_id
// 8714,121,1500.000000,"ЗА 24/07/2025;ГОРЕТСКАЯ ТАТЬЯНА АНДРЕЕВНА;ОПЛАТА ЗА ИП ГОРЕТСКИЙ А.Е.ПО ДОГОВОРУ N 3П0026 ОТ 01.08.2019 ПЛ-К ШАУМЯНА, Д 8 КВ 63//",1250725ROCO#DS4013578,2025-07-25 01:30:17,2025-08-23 02:54:53,,,0,6,,0,,1500.000000,1,,,,0,14205,,2,,1

// record_id,payment_id,agrm_id,order_id,amount,last_mod_date
// 4043,8714,121,,1500.000000,2025-08-23 02:54:53

// record_id,payment_id,agrm_id,sale_id,amount,last_mod_date
// 8724,8714,121,,1500.000000,2025-08-23 02:54:53
