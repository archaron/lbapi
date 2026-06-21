package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// VersionRequest — запрос версии сервера LANBilling.
type VersionRequest struct{}

func (VersionRequest) Method() string { return "version" }

// VersionInfo — информация о версии сервера.
type VersionInfo struct {
	Build            string `json:"build"`
	ReleaseDate      string `json:"release_date"`
	Revision         string `json:"revision"`
	Version          string `json:"version"`
	VersionSupported bool   `json:"version_supported"`
}

// Version — получить версию сервера LANBilling.
func (api *Client) Version(ctx context.Context) (*VersionInfo, error) {
	var result VersionInfo
	if err := api.client.CallResult(ctx, "version", &VersionRequest{}, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
