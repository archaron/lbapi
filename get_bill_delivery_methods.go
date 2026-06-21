package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetBillDeliveryMethodsRequest — запрос способов доставки счетов.
type GetBillDeliveryMethodsRequest struct {
	Japi int `json:"japi"`
}

func (GetBillDeliveryMethodsRequest) Method() string { return "getBillDeliveryMethods" }

// GetBillDeliveryMethodsResponse — ответ со списком способов доставки.
type GetBillDeliveryMethodsResponse struct {
	Data  []types.LBBillDeliveryMethod `json:"data"`
	Total *int                         `json:"total"`
}

// GetBillDeliveryMethods — получить список способов доставки счетов.
func (api *Client) GetBillDeliveryMethods(ctx context.Context, request GetBillDeliveryMethodsRequest) (*GetBillDeliveryMethodsResponse, error) {
	var result GetBillDeliveryMethodsResponse
	if err := api.client.CallResult(ctx, "getBillDeliveryMethods", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
