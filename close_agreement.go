package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// CloseAgreementRequest — запрос на закрытие договора.
type CloseAgreementRequest struct {
	AgrmID    int    `json:"agrm_id"`    // ID договора
	CloseDate string `json:"close_date"` // Дата закрытия (YYYY-MM-DD)
}

func (CloseAgreementRequest) Method() string { return "closeAgreement" }

// CloseAgreement — закрыть договор.
func (api *Client) CloseAgreement(ctx context.Context, request CloseAgreementRequest) (bool, error) {
	var result bool
	if err := api.client.CallResult(ctx, "closeAgreement", &request, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
