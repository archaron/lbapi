package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetVgBlockListRequest — запрос истории блокировок учётной записи.
type GetVgBlockListRequest struct {
	Count  bool `json:"count"`
	Japi   int  `json:"japi"`
	PgNum  int  `json:"pg_num"`
	PgSize int  `json:"pg_size"`
	VgID   int  `json:"vg_id"`
}

func (GetVgBlockListRequest) Method() string { return "getVgBlockList" }

// GetVgBlockListResponse — ответ с историей блокировок.
type GetVgBlockListResponse struct {
	Data  []types.LBVgBlockRecord `json:"data"`
	Total *int                    `json:"total"`
}

// GetVgBlockList — получить историю блокировок учётной записи.
func (api *Client) GetVgBlockList(ctx context.Context, request GetVgBlockListRequest) (*GetVgBlockListResponse, error) {
	var result GetVgBlockListResponse
	if err := api.client.CallResult(ctx, "getVgBlockList", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
