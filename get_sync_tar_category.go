package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetSyncTarCategoryRequest struct {
	// TODO

	//types.Pagination
}

func (GetSyncTarCategoryRequest) Method() string { return "getSyncTarCategory" }

// GetSyncTarCategory - получить категории радиус-агента для синхронизации
func (api *Client) GetSyncTarCategory(ctx context.Context, request GetSyncTarCategoryRequest) (interface{}, error) {
	var response interface{}
	if err := api.client.CallResult(ctx, "getSyncTarCategory", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
