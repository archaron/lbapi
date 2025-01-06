package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// getGponPorts
type (
	GetGPONPortsRequest struct {
		PortId int `json:"port_id,omitempty"`
		types.Pagination
	}
)

func (GetGPONPortsRequest) Method() string {
	return "getGponPorts"
}

func (api *Client) GetGPONPorts(ctx context.Context, request GetGPONPortsRequest) ([]types.GPONPort, error) {

	var result []types.GPONPort

	if err := api.client.CallResult(ctx, "getGponPorts", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return result, nil
}
