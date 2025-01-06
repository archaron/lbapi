package types

type AgentOption struct {
	AgentID  int        `json:"agent_id"`
	Descr    string     `json:"descr"`
	Name     string     `json:"name"`
	TimeMark LBTimeMark `json:"time_mark"`
	Value    string     `json:"value"`
}
