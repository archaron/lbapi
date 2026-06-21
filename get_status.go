package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetStatusRequest — запрос статуса сервера.
type GetStatusRequest struct{}

func (GetStatusRequest) Method() string { return "getStatus" }

// GetStatus — получить статус сервера (количество аккаунтов, агентов, договоров и т.д.).
func (api *Client) GetStatus(ctx context.Context) (*types.LBStatus, error) {
	var result types.LBStatus
	if err := api.client.CallResult(ctx, "getStatus", &GetStatusRequest{}, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
