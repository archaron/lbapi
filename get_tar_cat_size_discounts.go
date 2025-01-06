package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetTarCatSizeDiscountsRequest struct {
	TarID  int `json:"tar_id,omitempty"`
	CatIDX int `json:"cat_idx,omitempty"`

	TimeMarkStart int `json:"time_mark_start"`
	types.Pagination
}

func (GetTarCatSizeDiscountsRequest) Method() string { return "getTarCatSizeDiscounts" }

// GetTarCatSizeDiscounts - Получить скидки по объему тарифной категории
func (api *Client) GetTarCatSizeDiscounts(ctx context.Context, request GetTarCatSizeDiscountsRequest) ([]types.TarCatSizeDiscount, error) {
	var response []types.TarCatSizeDiscount
	if err := api.client.CallResult(ctx, "getTarCatSizeDiscounts", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
