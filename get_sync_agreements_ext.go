package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetSyncAgreementsExtRequest struct {
		AgentID       int `json:"agent_id"`
		TimeMarkStart int `json:"time_mark_start"`
		types.Pagination
	}
)

func (GetSyncAgreementsExtRequest) Method() string {
	return "getSyncAgreementsExt"
}

// GetSyncAgreementsExt - получение дополнительных полей для договоров
func (api *Client) GetSyncAgreementsExt(ctx context.Context, request GetSyncAgreementsExtRequest) ([]types.SyncAgreementExt, error) {
	var result []types.SyncAgreementExt
	if err := api.client.CallResult(ctx, "getSyncAgreementsExt", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
