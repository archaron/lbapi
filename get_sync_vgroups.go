package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetSyncVGroupsRequest struct {
	AgentID       int   `json:"agent_id"`
	TimeMarkStart int64 `json:"time_mark_start"`
	types.Pagination
}

func (GetSyncVGroupsRequest) Method() string {
	return "getSyncVgroups"
}

// GetSyncVGroups Получение данных учетных записей для синхронизации
func (api *Client) GetSyncVGroups(ctx context.Context, request GetSyncVGroupsRequest) ([]types.LBSyncVGroup, error) {
	var result []types.LBSyncVGroup
	if err := api.client.CallResult(ctx, "getSyncVgroups", request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return result, nil
}
