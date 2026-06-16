package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// getJAPIVersion
type (
	PaymentRequest struct {
		AgrmID int     `json:"agrm_id"`
		Amount float64 `json:"amount"`
		/*
			Тип платежа при печати ККМ-чека:
			1 — Наличный
			2 — Безналичный
			3 — Продажа без чека
			4 — Возврат без чека
		*/
		CashCode        int           `json:"cash_code"`
		ClassID         int           `json:"class_id"`
		CurrID          int           `json:"curr_id"`
		FromAgrmID      int           `json:"from_agrm_id"`
		ModPerson       int           `json:"mod_person"`
		PayDate         *types.LBTime `json:"pay_date"`
		PaymentOrderNum string        `json:"payment_order_num"`
		Receipt         string        `json:"receipt"`
		RecordID        int           `json:"record_id"`
		OrderIDs        []int         `json:"order_ids"`
		Comment         string        `json:"comment"`
	}
	// agrm_id: 82; amount: 650; from_agrm_id: 0; mod_person: 0; pay_date: 2025-08-11 00:15:48; receipt: 20250815121548-9240; record_id: 0; orders_ids: ; cash_code: 1
	// { "agrm_id": 198, "amount": 10.000000, "cash_code": 2, "class_id": 2, "comment": "Comment", "curr_id": 1, "mod_person": 0, "pay_date": "2025-08-23 00:58:41", "payment_order_num": "Recieps",
	// "receipt": "20250858125841-7731" } }

	PaymentResponse int
)

func (PaymentRequest) Method() string {
	return "Payment"
}

func (api *Client) Payment(ctx context.Context, request PaymentRequest) (*PaymentResponse, error) {
	var result PaymentResponse

	if err := api.client.CallResult(ctx, "Payment", request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return &result, nil
}
