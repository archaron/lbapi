package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// SetAddressRequest — запрос на установку адреса пользователя.
type SetAddressRequest struct {
	Code string `json:"code"` // Код КЛАДР
	Type int    `json:"type"` // Тип адреса (0=основной, 1=дополнительный)
	UID  int    `json:"uid"`  // ID пользователя
}

func (SetAddressRequest) Method() string { return "setAddress" }

// SetAddress — установить адрес пользователя.
func (api *Client) SetAddress(ctx context.Context, code string, addrType int, uid int) (bool, error) {
	var result bool
	req := SetAddressRequest{Code: code, Type: addrType, UID: uid}
	if err := api.client.CallResult(ctx, "setAddress", &req, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
