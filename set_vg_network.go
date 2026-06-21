package lbapi

import (
	"context"

	"github.com/archaron/lbapi/types"
)

// SetVgNetworkRequest — запрос на добавление/изменение сети УЗ.
type SetVgNetworkRequest struct {
	AsNum     int    `json:"as_num"`     // Номер автономной системы
	Netmask   string `json:"netmask"`    // Маска сети
	Network   string `json:"network"`    // IP-адрес сети
	RecordID  int    `json:"record_id"`  // 0 = создать новую, >0 = изменить существующую
	SegmentID int    `json:"segment_id"` // ID сегмента
	Type      int    `json:"type"`       // Тип сети (0 = обычная)
	VgID      int    `json:"vg_id"`      // ID учётной записи
}

func (SetVgNetworkRequest) Method() string { return "setVgNetwork" }

// SetVgNetwork — добавить или изменить сеть учётной записи. Возвращает record_id.
func (api *Client) SetVgNetwork(ctx context.Context, request SetVgNetworkRequest) (int, error) {
	var result int
	if err := api.client.CallResult(ctx, "setVgNetwork", &request, &result); err != nil {
		return 0, types.ParseJSONRPCError(err)
	}
	return result, nil
}
