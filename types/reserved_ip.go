package types

// LBReservedIP описывает зарезервированный IP-адрес.
type LBReservedIP struct {
	IP        string     `json:"ip"`
	VgID      int        `json:"vg_id"`
	Mac       string     `json:"mac,omitempty"`
	Netmask   string     `json:"netmask,omitempty"`
	SegmentID int        `json:"segment_id,omitempty"`
	TimeMark  LBTimeMark `json:"time_mark"`
}
