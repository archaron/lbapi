package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetTarCatTimeDiscountsRequest struct {
	TarID  int `json:"tar_id,omitempty"`
	CatIDX int `json:"cat_idx,omitempty"`

	TimeMarkStart int `json:"time_mark_start"`
	types.Pagination
}

func (GetTarCatTimeDiscountsRequest) Method() string { return "getTarCatTimeDiscounts" }

// GetTarCatTimeDiscounts - Получить скидки по объему тарифной категории
func (api *Client) GetTarCatTimeDiscounts(ctx context.Context, request GetTarCatTimeDiscountsRequest) ([]types.TarCatTimeDiscount, error) {
	var response []types.TarCatTimeDiscount
	if err := api.client.CallResult(ctx, "getTarCatTimeDiscounts", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
