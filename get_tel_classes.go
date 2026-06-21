package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetTelClassesRequest — запрос списка телефонных классов.
type GetTelClassesRequest struct{}

func (GetTelClassesRequest) Method() string { return "getTelClasses" }

// GetTelClasses — получить список телефонных классов (ВТС, МГ, МН и т.д.).
func (api *Client) GetTelClasses(ctx context.Context) ([]types.LBTelClass, error) {
	var result []types.LBTelClass
	if err := api.client.CallResult(ctx, "getTelClasses", &GetTelClassesRequest{}, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
