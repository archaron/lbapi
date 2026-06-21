package types

// LBBillDeliveryMethod — способ доставки счета.
type LBBillDeliveryMethod struct {
	IsEditable bool   `json:"is_editable"` // Редактируемый
	Name       string `json:"name"`        // Название способа
	RecordID   int    `json:"record_id"`   // ID записи
}
