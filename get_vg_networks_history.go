package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetVgNetworksHistoryRequest — запрос истории изменения сетей УЗ.
type GetVgNetworksHistoryRequest struct {
	Count       bool   `json:"count"`
	Japi        int    `json:"japi"`
	PgNum       int    `json:"pg_num"`
	PgSize      int    `json:"pg_size"`
	SegmentType int    `json:"segment_type"`
	TimeTo      string `json:"timeto"`
	VgID        int    `json:"vg_id"`
}

func (GetVgNetworksHistoryRequest) Method() string { return "getVgNetworksHistory" }

// GetVgNetworksHistoryResponse — ответ с историей сетей.
type GetVgNetworksHistoryResponse struct {
	Data  []types.LBVgNetworkHistory `json:"data"`
	Total *int                       `json:"total"`
}

// GetVgNetworksHistory — получить историю изменения сетей учётной записи.
func (api *Client) GetVgNetworksHistory(ctx context.Context, request GetVgNetworksHistoryRequest) (*GetVgNetworksHistoryResponse, error) {
	var result GetVgNetworksHistoryResponse
	if err := api.client.CallResult(ctx, "getVgNetworksHistory", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
