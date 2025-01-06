package types

type Weekend struct {
	IsHoliday bool       `json:"is_holiday"`
	Date      DateOnly   `json:"date"`
	TimeMark  LBTimeMark `json:"time_mark"`
}
