package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetVgNetworksRequest — запрос списка сетей учётной записи.
type GetVgNetworksRequest struct {
	Count  bool `json:"count"`
	Japi   int  `json:"japi"`
	PgNum  int  `json:"pg_num"`
	PgSize int  `json:"pg_size"`
	VgID   int  `json:"vg_id"`
}

func (GetVgNetworksRequest) Method() string { return "getVgNetworks" }

// GetVgNetworksResponse — ответ со списком сетей УЗ.
type GetVgNetworksResponse struct {
	Data  []types.LBVgNetwork `json:"data"`
	Total *int                `json:"total"`
}

// GetVgNetworks — получить список сетей учётной записи.
func (api *Client) GetVgNetworks(ctx context.Context, request GetVgNetworksRequest) (*GetVgNetworksResponse, error) {
	var result GetVgNetworksResponse
	if err := api.client.CallResult(ctx, "getVgNetworks", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
