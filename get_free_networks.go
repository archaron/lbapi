package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetFreeNetworksRequest — запрос поиска свободных подсетей.
type GetFreeNetworksRequest struct {
	AgentID           int              `json:"agent_id"`
	Count             bool             `json:"count"`
	IncludeBroadcast  bool             `json:"include_broadcast"`
	IncludeNetaddr    bool             `json:"include_netaddr"`
	Mask              int              `json:"mask"`
	Nodata            bool             `json:"nodata,omitempty"`
	PgNum             int              `json:"pg_num,omitempty"`
	PgSize            int              `json:"pg_size,omitempty"`
	Segment           types.LBSegmentRef `json:"segment"`
	VgID              int              `json:"vg_id"`
}

func (GetFreeNetworksRequest) Method() string { return "getFreeNetworks" }

// GetFreeNetworksCount — получить количество свободных подсетей.
func (api *Client) GetFreeNetworksCount(ctx context.Context, request GetFreeNetworksRequest) (int, error) {
	request.Count = true
	request.Nodata = true
	var result int
	if err := api.client.CallResult(ctx, "getFreeNetworks", &request, &result); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}
	return result, nil
}
