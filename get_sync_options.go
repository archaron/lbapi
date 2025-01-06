package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetSyncOptionsRequest struct {
		AgentID       int `json:"agent_id"`
		TimeMarkStart int `json:"time_mark_start"`
		types.Pagination
	}
)

func (GetSyncOptionsRequest) Method() string {
	return "getRnasOptions"
}

// GetSyncOptions - Получить все опции для агента для синхронизации
func (api *Client) GetSyncOptions(ctx context.Context, request GetSyncOptionsRequest) ([]types.LBSyncOption, error) {
	var result []types.LBSyncOption
	if err := api.client.CallResult(ctx, "getRnasOptions", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
