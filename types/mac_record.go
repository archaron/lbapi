package types

type LBMacRecord struct {
	Mac            string      `json:"mac"`
	Network        string      `json:"network"`
	NetworkID      int         `json:"network_id"`
	RecordID       int         `json:"record_id"`
	SegmentID      int         `json:"segment_id"`
	TimeMark       LBTimeMark  `json:"time_mark"`
	AllowOverwrite interface{} `json:"allow_overwrite"`
	VgID           int         `json:"vg_id"`
}
