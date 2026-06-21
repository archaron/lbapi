package types

// LBVgroupExt — расширенная информация об учётной записи (метод getVgroupExt).
type LBVgroupExt struct {
	AccOffDate        string        `json:"acc_off_date"`
	AccOnDate         string        `json:"acc_on_date"`
	AccOnDateFirst    string        `json:"acc_on_date_first"`
	ActBlock          int           `json:"act_block"`
	ActivationCode    *string       `json:"activation_code"`
	Addons            []interface{} `json:"addons"`
	Addresses         []LBAddress   `json:"addresses"`
	AgentID           int           `json:"agent_id"`
	AgentName         string        `json:"agent_name"`
	AgrmID            int           `json:"agrm_id"`
	AgrmNum           string        `json:"agrm_num"`
	Amount            float64       `json:"amount"`
	Archive           int           `json:"archive"`
	AuthMethod        *string       `json:"auth_method"`
	BlkReq            *int          `json:"blk_req"`
	BlockDate         string        `json:"block_date"`
	Blocked           int           `json:"blocked"`
	CDate             string        `json:"c_date"`
	ChangedTariffOn   *string       `json:"changedtariffon"`
	CreationDate      string        `json:"creation_date"`
	CurrentShape      int           `json:"current_shape"`
	DLimit            int           `json:"d_limit"`
	ID                int           `json:"id"`
	Login             string        `json:"login"`
	MaxSessions       int           `json:"max_sessions"`
	Pass              string        `json:"pass"`
	Shape             int           `json:"shape"`
	TarID             int           `json:"tar_id"`
	Template          int           `json:"template_"`
	TimeMark          LBTimeMark    `json:"time_mark"`
	UID               int           `json:"uid"`
	VgID              int           `json:"vg_id"`
}
