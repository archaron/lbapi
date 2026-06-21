package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// DhcpAcctStartRequest описывает запрос DHCP-аккаунтинга (старт сессии).
type DhcpAcctStartRequest types.DhcpAcctParams

func (DhcpAcctStartRequest) Method() string {
	return "dhcpAcctStart"
}

// DhcpAcctStart — сообщить серверу о старте DHCP-сессии.
func (api *Client) DhcpAcctStart(ctx context.Context, request DhcpAcctStartRequest) (bool, error) {
	var result bool
	if err := api.client.CallResult(ctx, "dhcpAcctStart", &request, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
