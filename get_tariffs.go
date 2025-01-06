package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetTariffsRequest struct {
		AgentID       int   `json:"agent_id"`
		TimeMarkStart int64 `json:"time_mark_start"`
		types.Pagination
	}
)

func (GetTariffsRequest) Method() string {
	return "getTariffs"
}

func (api *Client) GetTariffs(ctx context.Context, request GetTariffsRequest) ([]types.Tariff, error) {
	var result []types.Tariff
	if err := api.client.CallResult(ctx, "getTariffs", request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return result, nil
}
