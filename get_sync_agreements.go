package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetSyncAgreementsRequest struct {
		AgentID       int `json:"agent_id"`
		TimeMarkStart int `json:"time_mark_start"`
		types.Pagination
	}
)

func (GetSyncAgreementsRequest) Method() string {
	return "getSyncAgreements"
}

// GetSyncAgreements - получение данных договоров для синхронизации
func (api *Client) GetSyncAgreements(ctx context.Context, request GetSyncAgreementsRequest) ([]types.SyncAgreement, error) {
	var result []types.SyncAgreement
	if err := api.client.CallResult(ctx, "getSyncAgreements", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
