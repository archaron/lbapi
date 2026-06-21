package events

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/archaron/lbapi/types"
)

func TestLBEventChangeSegment_String(t *testing.T) {
	tests := []struct {
		name     string
		event    LBEventChangeSegment
		expected string
	}{
		{
			name: "Segment deleted",
			event: LBEventChangeSegment{
				SegmentID: 123,
				Deleted:   true,
			},
			expected: "Удалить сегмент: SegmentId:123",
		},
		{
			name: "Segment modified",
			event: LBEventChangeSegment{
				SegmentID: 456,
				Deleted:   false,
			},
			expected: "Изменить сегмент: SegmentId:456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := tt.event.String(); result != tt.expected {
				t.Errorf("Expected: %s, got: %s", tt.expected, result)
			}
		})
	}
}

func TestLBEventBlockVgTask_String(t *testing.T) {
	blockRequest := types.LBBlockRequest(types.BlockRequestBlock)
	event := LBEventBlockVgTask{
		BlkReq:   blockRequest,
		RecordID: 10,
		VgID:     20,
	}
	expected := "Изменить статус УЗ: VgID:20 [" + blockRequest.String() + "] №Записи: 10"
	require.Equal(t, expected, event.String())
}

func TestLBEventBlockVg_String(t *testing.T) {
	blockRequest := types.LBBlockRequest(types.BlockRequestBlock)
	now := time.Now()
	nowTime := types.LBTime(now)
	event := LBEventBlockVg{
		BlkReq:      blockRequest,
		BlockRaspID: 30,
		ChangeTime:  &nowTime,
		RequestBy:   40,
		VgID:        50,
	}
	expected := "Запланировать изменение статуса УЗ: VgID:50 [" + blockRequest.String() + "] №Записи: 30, Когда: " + nowTime.String() + ", ЗапросОт: 40, ПерваяАктивация: false"
	require.Equal(t, expected, event.String())
}

func TestAllLBEvents_ContainsUniqueEvents(t *testing.T) {
	eventMap := make(map[LBEvent]bool)
	for _, event := range AllLBEvents {
		require.Falsef(t, eventMap[event], "Duplicate event found: %s", event)
		eventMap[event] = true
	}
}

func TestLBEventChangeVgNetwork_String(t *testing.T) {
	event := LBEventChangeVgNetwork{
		AgentID:    100,
		VgID:       200,
		DeletedMAC: "00:11:22:33:44:55",
		DeletedSegments: []types.LBNetworkSegment{
			{
				Network:     "0.0.0.0",
				Netmask:     "255.255.255.255",
				SegmentType: "test_segment",
			},
		},
	}
	expected := "Изменена сеть УЗ: AgentID: 100, VgroupID: 200 Удален MAC: '00:11:22:33:44:55' Удалены сети: [ test_segment 0.0.0.0/255.255.255.255  ] "
	require.Equal(t, expected, event.String())
}

func TestLBEvent_JSONSerialization(t *testing.T) {
	event := LBEventChangeAgentOption{
		AgentID: 123,
		Name:    "TestOption",
	}
	data, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	var result LBEventChangeAgentOption

	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	require.Equal(t, event, result)
}

func TestLBEventChangeVgroup_String(t *testing.T) {
	tests := []struct {
		name     string
		event    LBEventChangeVgroup
		expected string
	}{
		{
			name: "With AgreementID",
			event: LBEventChangeVgroup{
				AgreementID: 123,
				VgID:        456,
			},
			expected: "ChangeVGroup: AgreementID:123 VgID: 456 ",
		},
		{
			name: "Without AgreementID",
			event: LBEventChangeVgroup{
				VgID: 789,
			},
			expected: "ChangeVGroup: VgID: 789 ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.event.String())
		})
	}
}

func TestLBEventChangeAgentOption_String(t *testing.T) {
	event := LBEventChangeAgentOption{
		AgentID: 101,
		Name:    "OptionName",
	}
	expected := "ChangeAgentOption: AgentID: 101, Option: OptionName"
	require.Equal(t, expected, event.String())
}

func TestLBEventCloseAgreement_String(t *testing.T) {
	event := LBEventCloseAgreement{
		AgrmID:    388,
		CloseDate: "2026-06-19 09:52:35",
	}
	expected := "Закрыть договор: AgrmID:388, ДатаЗакрытия:2026-06-19 09:52:35"
	require.Equal(t, expected, event.String())
}

func TestLBEventCloseAgreement_JSONSerialization(t *testing.T) {
	raw := `{"agrm_id":388,"close_date":"2026-06-19 09:52:35"}`
	var event LBEventCloseAgreement
	require.NoError(t, json.Unmarshal([]byte(raw), &event))
	require.Equal(t, 388, event.AgrmID)
	require.Equal(t, "2026-06-19 09:52:35", event.CloseDate)

	data, err := json.Marshal(event)
	require.NoError(t, err)
	var result LBEventCloseAgreement
	require.NoError(t, json.Unmarshal(data, &result))
	require.Equal(t, event, result)
}

func TestLBEventBlockVg_JSONSerialization(t *testing.T) {
	raw := `{"blk_req":10,"block_rasp_id":702,"change_time":"2026-06-19 09:52:35","first_activation":false,"old_block":0,"request_by":3,"vg_block_history_id":680,"vgid":424}`
	var event LBEventBlockVg
	require.NoError(t, json.Unmarshal([]byte(raw), &event))
	require.Equal(t, types.LBBlockRequest(10), event.BlkReq)
	require.Equal(t, 702, event.BlockRaspID)
	require.Equal(t, false, event.FirstActivation)
	require.Equal(t, 0, event.OldBlock)
	require.Equal(t, 3, event.RequestBy)
	require.Equal(t, 680, event.VgBlockHistoryID)
	require.Equal(t, 424, event.VgID)
}

func TestLBEventBlockVgTask_JSONSerialization(t *testing.T) {
	raw := `{"blk_req":10,"change_time":"2026-06-19 09:52:35","recordid":702,"vgid":424}`
	var event LBEventBlockVgTask
	require.NoError(t, json.Unmarshal([]byte(raw), &event))
	require.Equal(t, types.LBBlockRequest(10), event.BlkReq)
	require.Equal(t, "2026-06-19 09:52:35", event.ChangeTime)
	require.Equal(t, 702, event.RecordID)
	require.Equal(t, 424, event.VgID)
}

func TestLBEventDelRadiusVgroup_String(t *testing.T) {
	event := LBEventDelRadiusVgroup{
		VgID: 424,
	}
	expected := "Удалить RADIUS-группу: VgID:424"
	require.Equal(t, expected, event.String())
}

func TestLBEventDelRadiusVgroup_JSONSerialization(t *testing.T) {
	raw := `{"vg_id":424}`
	var event LBEventDelRadiusVgroup
	require.NoError(t, json.Unmarshal([]byte(raw), &event))
	require.Equal(t, 424, event.VgID)

	data, err := json.Marshal(event)
	require.NoError(t, err)
	var result LBEventDelRadiusVgroup
	require.NoError(t, json.Unmarshal(data, &result))
	require.Equal(t, event, result)
}

func TestLBEventPayment_String(t *testing.T) {
	event := LBEventPayment{
		PaymentID:   10282,
		AgrmID:      72,
		Amount:      45.01,
		Balance:     154.70,
		CashCode:    2,
		Payclass:    8,
		IsCancelled: false,
	}
	expected := "Платеж: PaymentID:10282, AgrmID:72, Сумма:45.01, Баланс:154.70, Касса:2, Класс:8, Отмена:false"
	require.Equal(t, expected, event.String())
}
