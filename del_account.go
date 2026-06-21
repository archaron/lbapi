package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// DelAccountRequest — запрос на удаление пользователя.
type DelAccountRequest struct {
	UID int `json:"uid"`
}

func (DelAccountRequest) Method() string { return "delAccount" }

// DelAccount — удалить пользователя.
func (api *Client) DelAccount(ctx context.Context, uid int) (bool, error) {
	var result bool
	if err := api.client.CallResult(ctx, "delAccount", &DelAccountRequest{UID: uid}, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
