package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetAgreementsRequest — запрос списка договоров.
type GetAgreementsRequest struct {
	Count  bool            `json:"count"`
	Japi   int             `json:"japi"`
	PgNum  int             `json:"pg_num"`
	PgSize int             `json:"pg_size"`
	Sort   []types.LBSort  `json:"sort,omitempty"`
	UID    int             `json:"uid"`
}

func (GetAgreementsRequest) Method() string { return "getAgreements" }

// GetAgreementsResponse — ответ со списком договоров.
type GetAgreementsResponse struct {
	Data  []types.LBAgreement `json:"data"`
	Total *int                `json:"total"`
}

// GetAgreements — получить список договоров пользователя.
func (api *Client) GetAgreements(ctx context.Context, request GetAgreementsRequest) (*GetAgreementsResponse, error) {
	var result GetAgreementsResponse
	if err := api.client.CallResult(ctx, "getAgreements", &request, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return &result, nil
}
