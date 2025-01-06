package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type PutRadiusStatRequest []types.RadiusStat

func (PutRadiusStatRequest) Method() string { return "putRadiusStat" }

// PutRadiusStat - Отправить статистику радиуса
func (api *Client) PutRadiusStat(ctx context.Context, request PutRadiusStatRequest) (bool, error) {
	var response bool
	if err := api.client.CallResult(ctx, "putRadiusStat", &request, &response); err != nil {
		return false, types.ParseJSONRPCError(err)
	}

	return response, nil
}
