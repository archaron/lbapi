package types

type LBSyncVGroup struct {
	AgrmID       int        `json:"agrm_id"`
	Amount       float32    `json:"amount"`
	Archive      int        `json:"archive"`
	BlockDate    string     `json:"block_date"`
	Blocked      int        `json:"blocked"`
	CDate        string     `json:"c_date"`
	Changed      int        `json:"changed"`
	CreationDate string     `json:"creation_date"`
	CurrentShape int        `json:"current_shape"`
	DLimit       int        `json:"d_limit"`
	ID           int        `json:"id"`
	Login        string     `json:"login"`
	MaxSessions  int        `json:"max_sessions"`
	Pass         string     `json:"pass"`
	Shape        int        `json:"shape"`
	TarID        int        `json:"tar_id"`
	Template     int        `json:"template_"`
	TimeMark     LBTimeMark `json:"time_mark"`
	UID          int        `json:"uid"`
	VgID         int        `json:"vg_id"`
}
