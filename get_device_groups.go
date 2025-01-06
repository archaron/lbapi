package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetDeviceGroupsRequest struct {
	TimeMarkStart int `json:"time_mark_start"`
	types.Pagination
}

func (GetDeviceGroupsRequest) Method() string {
	return "getDeviceGroups"
}

// GetDeviceGroups Возвращает информацию о группах операторского оборудования
func (api *Client) GetDeviceGroups(ctx context.Context, request GetDeviceGroupsRequest) ([]types.LBDeviceGroup, error) {
	var result []types.LBDeviceGroup
	if err := api.client.CallResult(ctx, "getDeviceGroups", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
