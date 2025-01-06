package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetSegmentsRequest struct {
	TimeMarkStart int `json:"time_mark_start"`
	types.Pagination
}

func (GetSegmentsRequest) Method() string {
	return "getSegments"
}

func (api *Client) GetSegments(ctx context.Context, request GetSegmentsRequest) ([]types.Segment, error) {
	var result []types.Segment
	if err := api.client.CallResult(ctx, "getSegments", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
