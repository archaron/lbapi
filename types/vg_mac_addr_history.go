package types

// LBVgMacAddrHistory — запись об истории изменения MAC-адресов УЗ.
type LBVgMacAddrHistory struct {
	Mac       string `json:"mac"`
	MacID     int    `json:"mac_id"`
	Segment   string `json:"segment"`
	SegmentID int    `json:"segment_id"`
	TimeFrom  string `json:"timefrom"`
	TimeTo    string `json:"timeto"`
	VgID      int    `json:"vg_id"`
}
