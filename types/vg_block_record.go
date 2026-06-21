package types

// LBVgBlockRecord — запись об истории блокировок учётной записи.
type LBVgBlockRecord struct {
	AgrmID              int     `json:"agrm_id"`
	AgrmNum             string  `json:"agrm_num"`
	BlockType           int     `json:"block_type"`
	ChangeTime          string  `json:"change_time"`
	Comment             string  `json:"comment"`
	CurID               int     `json:"cur_id"`
	CurrSymbol          string  `json:"curr_symbol"`
	IsFreewill          bool    `json:"is_freewill"`
	IsHistory           bool    `json:"is_history"`
	Login               string  `json:"login"`
	ManagerLogin        *string `json:"manager_login"`
	MgrDescr            *string `json:"mgr_descr"`
	MgrName             *string `json:"mgr_name"`
	RecordID            int     `json:"record_id"`
	RequestBy           *int    `json:"request_by"`
	TimeTo              string  `json:"time_to"`
	UID                 int     `json:"uid"`
	UnblockedBy         int     `json:"unblocked_by"`
	UnblockedManagerDescr *string `json:"unblocked_manager_descr"`
	UnblockedManagerLogin *string `json:"unblocked_manager_login"`
	UnblockedManagerName  *string `json:"unblocked_manager_name"`
	UnblockedTime        string  `json:"unblocked_time"`
	VgID                int     `json:"vg_id"`
}
