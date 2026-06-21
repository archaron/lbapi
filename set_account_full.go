package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// SetAccountFullRequest — запрос на создание/обновление пользователя.
type SetAccountFullRequest types.LBAccountFull

func (SetAccountFullRequest) Method() string { return "setAccountFull" }

// SetAccountFull — создать или обновить пользователя. Возвращает uid.
func (api *Client) SetAccountFull(ctx context.Context, request SetAccountFullRequest) (int, error) {
	var result int
	if err := api.client.CallResult(ctx, "setAccountFull", &request, &result); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}
	return result, nil
}
