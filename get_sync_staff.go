package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

//getSyncStaff

type GetSyncStaffRequest struct {
	AgentID       int   `json:"agent_id"`
	TimeMarkStart int64 `json:"time_mark_start"`
	types.Pagination
}

func (GetSyncStaffRequest) Method() string {
	return "getSyncStaff"
}

// GetSyncStaff - получить статические IP для агента для синхронизации
func (api *Client) GetSyncStaff(ctx context.Context, request GetSyncStaffRequest) ([]types.LBSyncStaff, error) {
	var result []types.LBSyncStaff
	if err := api.client.CallResult(ctx, "getSyncStaff", request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
