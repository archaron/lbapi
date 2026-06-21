package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetPassTemplatesRequest — запрос шаблонов паролей.
type GetPassTemplatesRequest struct {
	ObjectType string `json:"object_type"` // Тип объекта: "account"
}

func (GetPassTemplatesRequest) Method() string { return "getPassTemplates" }

// GetPassTemplates — получить список шаблонов паролей.
func (api *Client) GetPassTemplates(ctx context.Context, objectType string) ([]types.LBPassTemplate, error) {
	var result []types.LBPassTemplate
	if err := api.client.CallResult(ctx, "getPassTemplates", &GetPassTemplatesRequest{ObjectType: objectType}, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
