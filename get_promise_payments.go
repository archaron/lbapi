package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetPromisePaymentsRequest — запрос количества обещанных платежей.
type GetPromisePaymentsRequest struct {
	AgrmID int  `json:"agrm_id"`
	Count  bool `json:"count"`
	Nodata bool `json:"nodata"`
	Payed  int  `json:"payed"`
}

func (GetPromisePaymentsRequest) Method() string { return "getPromisePayments" }

// GetPromisePaymentsCount — получить количество обещанных платежей по договору.
func (api *Client) GetPromisePaymentsCount(ctx context.Context, agrmID int) (int, error) {
	var result int
	req := GetPromisePaymentsRequest{AgrmID: agrmID, Count: true, Nodata: true, Payed: 3}
	if err := api.client.CallResult(ctx, "getPromisePayments", &req, &result); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}
	return result, nil
}
