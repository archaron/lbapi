package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetDictionaryRequest struct {
		TimeMarkStart int `json:"time_mark_start"`
		types.Pagination
	}
)

func (GetDictionaryRequest) Method() string {
	return "getDictionary"
}

func (api *Client) GetDictionary(ctx context.Context, request GetDictionaryRequest) ([]types.Dictionary, error) {
	var result []types.Dictionary
	if err := api.client.CallResult(ctx, "getDictionary", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
