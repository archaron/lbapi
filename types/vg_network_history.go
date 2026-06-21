package types

// LBVgNetworkHistory — запись об истории изменения сетей УЗ.
type LBVgNetworkHistory struct {
	LastModDate string  `json:"last_mod_date"`
	Mask        string  `json:"mask"`
	RecordID    int     `json:"record_id"`
	Segment     string  `json:"segment"`
	SegmentID   int     `json:"segment_id"`
	ServiceID   *int    `json:"service_id"`
	TimeFrom    string  `json:"timefrom"`
	TimeTo      string  `json:"timeto"`
	Type        int     `json:"type"`
	VgID        int     `json:"vg_id"`
}
