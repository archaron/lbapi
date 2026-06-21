package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetOptionRequest — запрос значения одной опции.
type GetOptionRequest struct {
	Name string `json:"name"` // Имя опции (например, "tax_value")
}

func (GetOptionRequest) Method() string { return "getOption" }

// GetOption — получить значение опции по имени.
func (api *Client) GetOption(ctx context.Context, name string) (*types.LBOption, error) {
	var result types.LBOption
	if err := api.client.CallResult(ctx, "getOption", &GetOptionRequest{Name: name}, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
