package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// SetVgBlockScheduleRequest — запрос на планирование блокировки/разблокировки УЗ.
type SetVgBlockScheduleRequest struct {
	BlkReq     int    `json:"blk_req"`     // Тип блокировки (0 = разблокировать)
	ChangeTime string `json:"change_time"` // Время изменения (YYYY-MM-DD HH:MM:SS)
	VgID       int    `json:"vg_id"`       // ID учётной записи
}

func (SetVgBlockScheduleRequest) Method() string { return "setVgBlockSchedule" }

// SetVgBlockSchedule — запланировать изменение статуса блокировки УЗ. Возвращает record_id.
func (api *Client) SetVgBlockSchedule(ctx context.Context, request SetVgBlockScheduleRequest) (int, error) {
	var result int
	if err := api.client.CallResult(ctx, "setVgBlockSchedule", &request, &result); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}
	return result, nil
}
