package types

type VGroupRecord struct {
	AccOffDate     string      `json:"acc_off_date"`
	AccOnDate      string      `json:"acc_on_date"`
	ActivationCode interface{} `json:"activation_code"`
	Addresses      []struct {
		Address string      `json:"address"`
		Code    string      `json:"code"`
		Type    interface{} `json:"type"`
	} `json:"addresses"`
	AgentID          int         `json:"agent_id"`
	AgentName        string      `json:"agent_name"`
	AgentType        AgentType   `json:"agent_type"`
	AgrmID           int         `json:"agrm_id"`
	AgrmNum          string      `json:"agrm_num"`
	Archive          int         `json:"archive"`
	AuthMethod       string      `json:"auth_method"`
	Balance          float32     `json:"balance"`
	BlkReq           int         `json:"blk_req"`
	BlockDate        string      `json:"block_date"`
	Blocked          int         `json:"blocked"`
	ChangedTariffOn  interface{} `json:"changedtariffon"`
	CreationDate     string      `json:"creation_date"`
	CuID             interface{} `json:"cu_id"`
	DClear           string      `json:"d_clear"`
	DLimit           int         `json:"d_limit"`
	Descr            string      `json:"descr"`
	Dirty            int         `json:"dirty"`
	Login            string      `json:"login"`
	MaxSessions      int         `json:"max_sessions"`
	MonthlyRent      interface{} `json:"monthlyrent"`
	ParentVgID       interface{} `json:"parent_vg_id"`
	ParentVgLogin    interface{} `json:"parent_vg_login"`
	Pass             string      `json:"pass"`
	PayCode          interface{} `json:"pay_code"`
	PpDebt           interface{} `json:"pp_debt"`
	RenewalState     int         `json:"renewal_state"`
	Rent             interface{} `json:"rent"`
	SetID            interface{} `json:"set_id"`
	Symbol           string      `json:"symbol"`
	TarID            int         `json:"tar_id"`
	TarName          string      `json:"tar_name"`
	Template         int         `json:"template"`
	TimeMark         LBTimeMark  `json:"time_mark"`
	TotalMonthlyRent interface{} `json:"totalmonthlyrent"`
	UID              int         `json:"uid"`
	UserCategory     int         `json:"user_category"`
	UserName         string      `json:"user_name"`
	UserType         interface{} `json:"user_type"`
	VgID             int         `json:"vg_id"`
}
