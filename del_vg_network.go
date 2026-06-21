package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// DelVgNetworkRequest — запрос на удаление сети УЗ.
type DelVgNetworkRequest struct {
	RecordID int `json:"record_id"`
}

func (DelVgNetworkRequest) Method() string { return "delVgNetwork" }

// DelVgNetwork — удалить сеть учётной записи.
func (api *Client) DelVgNetwork(ctx context.Context, recordID int) (bool, error) {
	var result bool
	if err := api.client.CallResult(ctx, "delVgNetwork", &DelVgNetworkRequest{RecordID: recordID}, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
