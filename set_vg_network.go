package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type (
	SetVgNetworkRequest struct {
		AsNum     int    `json:"as_num"`
		Netmask   string `json:"netmask"`
		Network   string `json:"network"`
		RecordId  int    `json:"record_id"`
		SegmentId int    `json:"segment_id"`
		Type      int    `json:"type"`
		VgId      int    `json:"vg_id"`
	}
)

func (SetVgNetworkRequest) Method() string {
	return "setVgNetwork"
}

func (api *Client) SetVgNetwork(ctx context.Context, request SetVgNetworkRequest) (int, error) {
	var result int
	if err := api.client.CallResult(ctx, "setVgNetwork", request, &result); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}

	return result, nil
}
