package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetReservedIpAddressesRequest — запрос на получение зарезервированных IP-адресов.
type GetReservedIpAddressesRequest struct {
	AgentID int `json:"agent_id"`
	types.Pagination
}

func (GetReservedIpAddressesRequest) Method() string {
	return "getReservedIpAddresses"
}

// GetReservedIpAddresses — получить список зарезервированных IP-адресов агента.
func (api *Client) GetReservedIpAddresses(ctx context.Context, request GetReservedIpAddressesRequest) ([]types.LBReservedIP, error) {
	var result []types.LBReservedIP
	if err := api.client.CallResult(ctx, "getReservedIpAddresses", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
