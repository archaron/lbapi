package types

// LBUserBlockTemplate — шаблон блокировки учётной записи.
type LBUserBlockTemplate struct {
	BlkReq     int    `json:"blk_req"`
	ChangeTime string `json:"change_time"`
	Comment    string `json:"comment"`
	RecordID   int    `json:"record_id"`
	VgID       int    `json:"vg_id"`
}
