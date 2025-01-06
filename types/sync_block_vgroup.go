package types

type LBSyncBlockVgroup struct {
	BlockType int        `json:"block_type"`
	RecordId  int        `json:"record_id"`
	TimeMark  LBTimeMark `json:"time_mark"`
	Timefrom  LBTime     `json:"timefrom"`
	Timeto    LBTime     `json:"timeto"`
	VgId      int        `json:"vg_id"`
}
