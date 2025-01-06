package types

type TarCatTimeDiscount struct {
	CatIdx   int        `json:"cat_idx"`
	DisID    int        `json:"dis_id"`
	Discount float64    `json:"discount"`
	TarID    int        `json:"tar_id"`
	TimeFrom TimeOnly   `json:"time_from"`
	TimeTo   TimeOnly   `json:"time_to"`
	TimeMark LBTimeMark `json:"time_mark"`

	Type       int  `json:"type"`
	UseWeekend bool `json:"use_weekend"`
	Week       struct {
		Mon   bool `json:"mon"`
		Tue   bool `json:"tue"`
		Wed   bool `json:"wed"`
		Thu   bool `json:"thu"`
		Fri   bool `json:"fri"`
		Sat   bool `json:"sat"`
		Sun   bool `json:"sun"`
		Value int  `json:"value"`
	} `json:"week"`
}
