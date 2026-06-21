package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// GetEpsAgreementsParamsRequest — запрос параметров договора для ЕПС.
type GetEpsAgreementsParamsRequest struct {
	AgrmID      int  `json:"agrm_id"`
	Autopayment bool `json:"autopayment"`
}

func (GetEpsAgreementsParamsRequest) Method() string { return "getEpsAgreementsParams" }

// GetEpsAgreementsParams — получить параметры договора для ЕПС.
func (api *Client) GetEpsAgreementsParams(ctx context.Context, agrmID int) ([]types.LBEpsAgreementParam, error) {
	var result []types.LBEpsAgreementParam
	req := GetEpsAgreementsParamsRequest{AgrmID: agrmID, Autopayment: true}
	if err := api.client.CallResult(ctx, "getEpsAgreementsParams", &req, &result); err != nil {
		return nil, types.ParseJSONRPCError(err)
	}
	return result, nil
}
