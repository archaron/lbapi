package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetUserBlockTemplateRequest — запрос шаблонов блокировок УЗ.
type GetUserBlockTemplateRequest struct {
	VgID int `json:"vg_id"`
}

func (GetUserBlockTemplateRequest) Method() string { return "getUserBlockTemplate" }

// GetUserBlockTemplate — получить шаблоны блокировок учётной записи.
func (api *Client) GetUserBlockTemplate(ctx context.Context, vgID int) ([]types.LBUserBlockTemplate, error) {
	var result []types.LBUserBlockTemplate
	if err := api.client.CallResult(ctx, "getUserBlockTemplate", &GetUserBlockTemplateRequest{VgID: vgID}, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
