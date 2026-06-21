package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetVgroupExtRequest — запрос расширенной информации об учётной записи.
type GetVgroupExtRequest struct {
	VgID int `json:"vg_id"`
}

func (GetVgroupExtRequest) Method() string { return "getVgroupExt" }

// GetVgroupExt — получить расширенную информацию об учётной записи.
func (api *Client) GetVgroupExt(ctx context.Context, vgID int) (*types.LBVgroupExt, error) {
	var result types.LBVgroupExt
	if err := api.client.CallResult(ctx, "getVgroupExt", &GetVgroupExtRequest{VgID: vgID}, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
