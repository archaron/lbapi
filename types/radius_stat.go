package types

type RadiusStat struct {
	TimeFrom        LBTime `json:"time_from"`
	TimeTo          LBTime `json:"time_to"`
	VgId            int    `json:"vg_id"`
	AgrmId          int    `json:"agrm_id"`
	ParentSessionId string `json:"parent_session_id"`
	Ani             string `json:"ani"`
	Dnis            string `json:"dnis"`
	IpAddress       string `json:"ip_address"`
	FramedPrefix    string `json:"framed_prefix"`
	DelegatedPrefix string `json:"delegated_prefix"`
	NasIp           string `json:"nas_ip"`
	Cause           int    `json:"cause"`
	AgentId         int    `json:"agent_id"`
	IsGuest         bool   `json:"is_guest"`
	SessionId       string `json:"session_id"`
	CatIdx          int    `json:"cat_idx"`
	DeltaCin        int    `json:"delta_cin"`
	DeltaCout       int    `json:"delta_cout"`
	DeltaTm         int    `json:"delta_tm"`
	Service         string `json:"service"`
}
