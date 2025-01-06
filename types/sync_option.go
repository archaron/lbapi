package types

type LBSyncOption struct {
	Descr    string     `json:"descr"`
	Name     string     `json:"name"`
	TimeMark LBTimeMark `json:"time_mark"`
	Value    string     `json:"value"`
}
