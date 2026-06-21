package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetAccountRequest — запрос информации об одном пользователе.
type GetAccountRequest struct {
	UID int `json:"uid"`
}

func (GetAccountRequest) Method() string { return "getAccount" }

// GetAccount — получить полную информацию о пользователе.
func (api *Client) GetAccount(ctx context.Context, uid int) (*types.LBAccountFull, error) {
	var result types.LBAccountFull
	if err := api.client.CallResult(ctx, "getAccount", &GetAccountRequest{UID: uid}, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
