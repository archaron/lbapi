package types

type (
	RNas struct {
		AgentID     int           `json:"agent_id"`
		AgentName   string        `json:"agent_name"`
		AuthMethod  string        `json:"auth_method"`
		DaePort     int           `json:"dae_port"`
		DaeServer   string        `json:"dae_server"`
		Descr       string        `json:"descr"`
		DeviceID    *int          `json:"device_id"`
		DeviceName  *string       `json:"device_name"`
		IsNew       interface{}   `json:"is_new"`
		NasID       int           `json:"nas_id"`
		Radblacklog []interface{} `json:"radblacklog"`
		Rnas        string        `json:"rnas"`
		Secret      string        `json:"secret"`
		TimeMark    LBTimeMark    `json:"time_mark"`
		Timezone    float64       `json:"timezone"`
		Type        string        `json:"type"`
	}

	LBRNasOption struct {
		Name     string     `json:"name"`
		NasID    int        `json:"nas_id"`
		OptionID int        `json:"option_id"`
		TimeMark LBTimeMark `json:"time_mark"`
		Value    string     `json:"value"`
	}
)
