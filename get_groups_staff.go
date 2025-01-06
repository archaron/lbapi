package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	GetGroupsStaffRequest struct {
		VgID int `json:"vg_id,omitempty"`
		types.Pagination
	}
)

func (GetGroupsStaffRequest) Method() string {
	return "getGroupsStaff"
}

// GetGroupsStaff - получить принадлежность учетной записи группам учетных записей
func (api *Client) GetGroupsStaff(ctx context.Context, request GetGroupsStaffRequest) ([]types.GroupStaff, error) {

	var result []types.GroupStaff

	if err := api.client.CallResult(ctx, "getGroupsStaff", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return result, nil
}
