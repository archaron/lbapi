package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type SetNewMacForVGRequest struct {
	// TODO
}

func (SetNewMacForVGRequest) Method() string { return "setNewMacForVg" }

// SetNewMacForVG - установить новый мак-адрес для учетной записи
func (api *Client) SetNewMacForVG(ctx context.Context, request SetNewMacForVGRequest) (interface{}, error) {
	var response interface{}
	if err := api.client.CallResult(ctx, "setNewMacForVg", &request, &response); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return response, nil
}
