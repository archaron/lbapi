package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetSyncTarShapePolicyRequest struct {
	AgentID       int `json:"agent_id"`
	TimeMarkStart int `json:"time_mark_start"`
	// TODO
	//types.Pagination
}

func (GetSyncTarShapePolicyRequest) Method() string { return "getSyncTarShapePolicy" }

// GetSyncTarShapePolicy - Получить правила ограничения скорости для агента для синхронизации
func (api *Client) GetSyncTarShapePolicy(ctx context.Context, request GetSyncTarShapePolicyRequest) (interface{}, error) {
	var response interface{}
	if err := api.client.CallResult(ctx, "getSyncTarShapePolicy", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
