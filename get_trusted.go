package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetTrustedRequest — запрос списка доверенных сетей.
type GetTrustedRequest struct {
	Count  bool `json:"count"`
	Japi   int  `json:"japi"`
	PgNum  int  `json:"pg_num"`
	PgSize int  `json:"pg_size"`
}

func (GetTrustedRequest) Method() string { return "getTrusted" }

// GetTrustedResponse — ответ со списком доверенных сетей.
type GetTrustedResponse struct {
	Data  []types.LBTrustedNetwork `json:"data"`
	Total *int                     `json:"total"`
}

// GetTrusted — получить список доверенных сетей (IP-адреса, с которых разрешён доступ).
func (api *Client) GetTrusted(ctx context.Context, request GetTrustedRequest) (*GetTrustedResponse, error) {
	var result GetTrustedResponse
	if err := api.client.CallResult(ctx, "getTrusted", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
