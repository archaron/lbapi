package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetSyncVGroupsAccessOverridesRequest struct {
	AgentID       int `json:"agent_id"`
	TimeMarkStart int `json:"time_mark_start"`
	// TODO
	//types.Pagination
}

func (GetSyncVGroupsAccessOverridesRequest) Method() string { return "getSyncVgroupsAccessOverrides" }

// GetSyncVGroupsAccessOverrides - получить особые условия доступа для учетных записей в блоке для синхронизации
func (api *Client) GetSyncVGroupsAccessOverrides(ctx context.Context, request GetSyncVGroupsAccessOverridesRequest) (interface{}, error) {
	var response interface{}
	if err := api.client.CallResult(ctx, "getSyncVgroupsAccessOverrides", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
