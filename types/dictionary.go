package types

type Dictionary struct {
	IsDefault  int        `json:"is_default"`
	Name       string     `json:"name"`
	NasID      *int       `json:"nas_id"`
	RadiusType int        `json:"radius_type"`
	RecordID   int        `json:"record_id"`
	ReplaceID  *int       `json:"replace_id"`
	Tagged     bool       `json:"tagged"`
	TimeMark   LBTimeMark `json:"time_mark"`
	ToHistory  bool       `json:"to_history"`
	ValueType  int        `json:"value_type"`
	Vendor     int        `json:"vendor"`
}
