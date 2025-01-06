package types

type LBSyncStaff struct {
	Netmask   string         `json:"netmask"`
	Network   string         `json:"network"`
	RecordID  int            `json:"record_id"`
	SegmentID int            `json:"segment_id"`
	TimeMark  LBTimeMark     `json:"time_mark"`
	Type      StaffBlockType `json:"type"`
	VgID      int            `json:"vg_id"`
}
