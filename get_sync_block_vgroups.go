package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetSyncBlockVGroupsRequest struct {
		AgentID       int   `json:"agent_id"`
		TimeMarkStart int64 `json:"time_mark_start"`
		types.Pagination
	}
)

func (GetSyncBlockVGroupsRequest) Method() string {
	return "getSyncBlockVgroups"
}

// GetSyncBlockVGroups - Получить данные об истории блокировок учетных записей
func (api *Client) GetSyncBlockVGroups(ctx context.Context, request GetSyncBlockVGroupsRequest) ([]types.LBSyncBlockVgroup, error) {
	var result []types.LBSyncBlockVgroup
	if err := api.client.CallResult(ctx, "getSyncBlockVgroups", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return result, nil
}
