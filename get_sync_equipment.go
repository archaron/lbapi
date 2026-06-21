package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetSyncEquipmentRequest — запрос на синхронизацию оборудования агента.
type GetSyncEquipmentRequest struct {
	AgentID       int `json:"agent_id"`
	TimeMarkStart int `json:"time_mark_start"`
	types.Pagination
}

func (GetSyncEquipmentRequest) Method() string {
	return "getSyncEquipment"
}

// GetSyncEquipment — получить список оборудования агента для синхронизации.
func (api *Client) GetSyncEquipment(ctx context.Context, request GetSyncEquipmentRequest) ([]types.LBSyncEquipment, error) {
	var result []types.LBSyncEquipment
	if err := api.client.CallResult(ctx, "getSyncEquipment", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
