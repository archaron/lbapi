package types

// LBVgNetwork — сеть учётной записи.
type LBVgNetwork struct {
	AsNum       *int   `json:"as_num"`
	EquipID     *int   `json:"equip_id"`
	Hwaddr      *string `json:"hwaddr"`
	Netmask     string `json:"netmask"`
	Network     string `json:"network"`
	RecordID    int    `json:"record_id"`
	SegmentID   int    `json:"segment_id"`
	SegmentType int    `json:"segment_type"`
	ServiceID   *int   `json:"service_id"`
	TimeMark    *LBTimeMark `json:"time_mark"`
	Type        int    `json:"type"`
	VgID        int    `json:"vg_id"`
}
