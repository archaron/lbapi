package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetGroupsExtRequest — запрос расширенной информации о группах УЗ.
type GetGroupsExtRequest struct {
	Count  bool `json:"count"`
	Japi   int  `json:"japi"`
	PgNum  int  `json:"pg_num"`
	PgSize int  `json:"pg_size"`
}

func (GetGroupsExtRequest) Method() string { return "getGroupsExt" }

// GetGroupsExtResponse — ответ с расширенной информацией о группах УЗ.
type GetGroupsExtResponse struct {
	Data  []types.LBGroupExt `json:"data"`
	Total *int               `json:"total"`
}

// GetGroupsExt — получить список групп учётных записей с расширенной информацией.
func (api *Client) GetGroupsExt(ctx context.Context, request GetGroupsExtRequest) (*GetGroupsExtResponse, error) {
	var result GetGroupsExtResponse
	if err := api.client.CallResult(ctx, "getGroupsExt", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
