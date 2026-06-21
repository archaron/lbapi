package types

// LBSyncEquipment описывает оборудование, синхронизируемое агентом.
type LBSyncEquipment struct {
	EquipID      int        `json:"equip_id"`
	Mac          string     `json:"mac"`
	Name         string     `json:"name"`
	Serial       *string    `json:"serial"`
	SerialFormat int        `json:"serial_format"`
	TimeMark     LBTimeMark `json:"time_mark"`
	VgID         int        `json:"vg_id"`
}
