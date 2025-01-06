package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetWeekendsRequest struct {
	DateFrom types.DateOnly `json:"date_from,omitempty"`
	DateTo   types.DateOnly `json:"date_to,omitempty"`
	types.Pagination
}

func (GetWeekendsRequest) Method() string { return "getWeekends" }

func (api *Client) GetWeekends(ctx context.Context, request GetWeekendsRequest) ([]types.Weekend, error) {
	var response []types.Weekend
	if err := api.client.CallResult(ctx, "getWeekends", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
