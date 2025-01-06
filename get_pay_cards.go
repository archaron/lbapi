package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetPayCardsRequest struct {
	Activated   bool          `json:"is_activated"`
	SetID       int           `json:"set_id"`
	SerNo       string        `json:"ser_no"`
	CardKey     string        `json:"card_key"`
	DtCreated   *types.LBTime `json:"create_date"`
	DtActivated *types.LBTime `json:"activate_date"`
	types.Pagination
}

func (GetPayCardsRequest) Method() string {
	return "getPayCards"
}

func (api *Client) GetPayCards(ctx context.Context, request GetPayCardsRequest) ([]types.PayCard, error) {
	var result []types.PayCard
	if err := api.client.CallResult(ctx, "getPayCards", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
