package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

type GetLicenseRequest struct{}

func (GetLicenseRequest) Method() string { return "getLicense" }

func (api *Client) GetLicense(ctx context.Context) (*types.LBLicense, error) {

	var license *types.LBLicense
	if err := api.client.CallResult(ctx, "getLicense", nil, &license); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return license, nil
}
