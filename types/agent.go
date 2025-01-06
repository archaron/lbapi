package types

type LBAgent struct {
	AdmNotify        int         `json:"adm_notify"`
	AgentID          int         `json:"agent_id"`
	Day              string      `json:"day"`
	Descr            string      `json:"descr"`
	EAPCertPassword  string      `json:"eapcertpassword"`
	FailedCalls      int         `json:"failed_calls"`
	Flush            int         `json:"flush"`
	IgnoreLocal      int         `json:"ignorelocal"`
	KeepDetail       int         `json:"keepdetail"`
	LastContact      string      `json:"lastcontact"`
	LocalASNum       int         `json:"local_as_num"`
	MaxRadiusTimeout int         `json:"max_radius_timeout"`
	Multiply         int         `json:"multiply"`
	NaDB             string      `json:"na_db"`
	NaIP             string      `json:"na_ip"`
	NaPass           string      `json:"na_pass"`
	NaUsername       string      `json:"na_username"`
	Nfhost           interface{} `json:"nfhost"`
	Nfhost6          interface{} `json:"nfhost6"`
	Nfport           int         `json:"nfport"`
	OperCat          int         `json:"oper_cat"`
	Raccport         int         `json:"raccport"`
	RadStopExpired   int         `json:"rad_stop_expired"`
	Raddrpool        int         `json:"raddrpool"`
	Rauthport        int         `json:"rauthport"`
	RecalcCurrent    string      `json:"recalc_current"`
	RecalcDate       string      `json:"recalc_date"`
	RecalcGroup      int         `json:"recalc_group"`
	RecalcRent       int         `json:"recalc_rent"`
	RecalcStat       int         `json:"recalc_stat"`
	RemulateOnNaid   int         `json:"remulate_on_naid"`
	RestartShape     int         `json:"restart_shape"`
	SaveStatAddr     int         `json:"save_stat_addr"`
	ServiceName      string      `json:"service_name"`
	SessionLifetime  int         `json:"session_lifetime"`
	TelDirectionMode int         `json:"tel_direction_mode"`
	TimeMark         LBTimeMark  `json:"time_mark"`
	Timer            int         `json:"timer"`
	Type             int         `json:"type"`
	UseRtime         int         `json:"use_rtime"`
	VoipCardUser     string      `json:"voip_card_user"`
}
