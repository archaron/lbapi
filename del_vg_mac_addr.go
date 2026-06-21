package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// DelVgMacAddrRequest — запрос на удаление MAC-адреса УЗ.
type DelVgMacAddrRequest struct {
	RecordID int `json:"record_id"`
}

func (DelVgMacAddrRequest) Method() string { return "delVgMacAddr" }

// DelVgMacAddr — удалить MAC-адрес учётной записи.
func (api *Client) DelVgMacAddr(ctx context.Context, recordID int) (bool, error) {
	var result bool
	if err := api.client.CallResult(ctx, "delVgMacAddr", &DelVgMacAddrRequest{RecordID: recordID}, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
