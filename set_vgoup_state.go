package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// SetVgroupStateRequest — установить статус учетной записи.
// Вызывается после обработки блокировки/разблокировки УЗ агентом.
type SetVgroupStateRequest struct {
	ID       int `json:"id"`        // Внутренний ID записи
	VgID     int `json:"vg_id"`     // ID виртуальной группы
	Changed  int `json:"changed"`   // Флаг изменений
	TimeMark int `json:"time_mark"` // Временная метка
}

func (SetVgroupStateRequest) Method() string {
	return "setVgroupState"
}

// SetVgroupState — установить статус учетной записи.
func (api *Client) SetVgroupState(ctx context.Context, request SetVgroupStateRequest) (bool, error) {
	var result bool
	if err := api.client.CallResult(ctx, "setVgroupState", &request, &result); err != nil {
		return false, types.ParseJSONRPCError(err)
	}
	return result, nil
}
