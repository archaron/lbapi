package types

type AccountForRadius struct {
	Archive  bool   `json:"archive"`
	Name     string `json:"name"`
	Template int    `json:"template"`
	TimeMark int    `json:"time_mark"`
	UID      int    `json:"uid"`
}
