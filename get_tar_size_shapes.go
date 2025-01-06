package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetTarSizeShapesRequest struct {
	TarID int `json:"tar_id"`
}

func (GetTarSizeShapesRequest) Method() string { return "getTarSizeShapes" }

// GetTarSizeShapes - Получить настройки полосы пропускания тарифа в зависимости от объема
func (api *Client) GetTarSizeShapes(ctx context.Context, request GetTarSizeShapesRequest) ([]types.TarSizeShape, error) {
	var response []types.TarSizeShape
	if err := api.client.CallResult(ctx, "getTarSizeShapes", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
