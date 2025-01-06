package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// getJAPIVersion
type (
	GetJAPIVersionRequest struct{}

	GetJAPIVersionResponse struct {
		Major   int    `json:"major"`
		Minor   int    `json:"minor"`
		Patch   int    `json:"patch"`
		Version string `json:"version"`
	}
)

func (GetJAPIVersionRequest) Method() string {
	return "getJAPIVersion"
}

func (api *Client) GetJAPIVersion(ctx context.Context) (*GetJAPIVersionResponse, error) {
	var result GetJAPIVersionResponse

	if err := api.client.CallResult(ctx, "getJAPIVersion", nil, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}

	return &result, nil
}
