package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetSyncTurboShapesRequest struct {
	//AgentID       int `json:"agent_id"`
	//TimeMarkStart int `json:"time_mark_start"`

	//types.Pagination
}

func (GetSyncTurboShapesRequest) Method() string { return "getSyncTurboShapes" }

// GetSyncTurboShapes - получить информацию о турбо-кнопке для синхронизации
func (api *Client) GetSyncTurboShapes(ctx context.Context, request GetSyncTurboShapesRequest) (interface{}, error) {
	var response interface{}
	if err := api.client.CallResult(ctx, "getSyncTurboShapes", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
