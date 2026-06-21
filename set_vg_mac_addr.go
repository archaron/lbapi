package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// SetVgMacAddrRequest — запрос на добавление MAC-адреса УЗ.
type SetVgMacAddrRequest struct {
	Mac     string `json:"mac"`     // MAC-адрес
	Network string `json:"network"` // Привязанный IP-адрес
	VgID    int    `json:"vg_id"`   // ID учётной записи
}

func (SetVgMacAddrRequest) Method() string { return "setVgMacAddr" }

// SetVgMacAddr — добавить MAC-адрес учётной записи. Возвращает record_id.
func (api *Client) SetVgMacAddr(ctx context.Context, request SetVgMacAddrRequest) (int, error) {
	var result int
	if err := api.client.CallResult(ctx, "setVgMacAddr", &request, &result); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}
	return result, nil
}
