package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetVgMacAddrHistoryRequest — запрос истории изменения MAC-адресов УЗ.
type GetVgMacAddrHistoryRequest struct {
	Count  bool   `json:"count"`
	Japi   int    `json:"japi"`
	PgNum  int    `json:"pg_num"`
	PgSize int    `json:"pg_size"`
	TimeTo string `json:"timeto"`
	VgID   int    `json:"vg_id"`
}

func (GetVgMacAddrHistoryRequest) Method() string { return "getVgMacAddrHistory" }

// GetVgMacAddrHistoryResponse — ответ с историей MAC-адресов.
type GetVgMacAddrHistoryResponse struct {
	Data  []types.LBVgMacAddrHistory `json:"data"`
	Total *int                       `json:"total"`
}

// GetVgMacAddrHistory — получить историю изменения MAC-адресов учётной записи.
func (api *Client) GetVgMacAddrHistory(ctx context.Context, request GetVgMacAddrHistoryRequest) (*GetVgMacAddrHistoryResponse, error) {
	var result GetVgMacAddrHistoryResponse
	if err := api.client.CallResult(ctx, "getVgMacAddrHistory", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
