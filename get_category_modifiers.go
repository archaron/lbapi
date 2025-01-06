package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetCategoryModifiersRequest struct {
		RecordID      int `json:"record_id,omitempty"`
		TimeMarkStart int `json:"time_mark_start"`
		types.Pagination
	}
)

func (GetCategoryModifiersRequest) Method() string {
	return "getCategoryModifiers"
}

func (api *Client) GetCategoryModifiers(ctx context.Context, request GetCategoryModifiersRequest) ([]types.CategoryModifier, error) {
	var result []types.CategoryModifier
	if err := api.client.CallResult(ctx, "getCategoryModifiers", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
