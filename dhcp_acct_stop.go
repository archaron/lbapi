package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// DhcpAcctStopRequest описывает запрос DHCP-аккаунтинга (стоп сессии).
type DhcpAcctStopRequest types.DhcpAcctParams

func (DhcpAcctStopRequest) Method() string {
	return "dhcpAcctStop"
}

// DhcpAcctStop — сообщить серверу об остановке DHCP-сессии.
func (api *Client) DhcpAcctStop(ctx context.Context, request DhcpAcctStopRequest) (bool, error) {
	var result bool
	if err := api.client.CallResult(ctx, "dhcpAcctStop", &request, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
