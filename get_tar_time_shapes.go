package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetTarTimeShapesRequest struct {
	TarID int `json:"tar_id"`
}

func (GetTarTimeShapesRequest) Method() string { return "getTarTimeShapes" }

// GetTarTimeShapes Получить настройки полосы пропускания тарифа в зависимости от времени
func (api *Client) GetTarTimeShapes(ctx context.Context, request GetTarTimeShapesRequest) ([]types.TarTimeShape, error) {
	var response []types.TarTimeShape
	if err := api.client.CallResult(ctx, "getTarTimeShapes", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
