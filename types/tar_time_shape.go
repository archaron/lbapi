package types

type (

	// TarTimeShape Настройки полосы пропускания тарифа в зависимости от времени
	TarTimeShape struct {
		ID         int        `json:"id"`          // Идентификатор
		TarID      int        `json:"tar_id"`      // Идентификатор тарифа
		ShapeRate  float64    `json:"shape_rate"`  // Полоса пропускания
		TimeFrom   TimeOnly   `json:"time_from"`   // Начало периода применения
		TimeTo     TimeOnly   `json:"time_to"`     // Конец периода применения
		TimeMark   LBTimeMark `json:"time_mark"`   // Отметка времени
		UseWeekend bool       `json:"use_weekend"` // Использовать календарь праздников
		Week       struct {
			Mon   bool `json:"mon"`
			Tue   bool `json:"tue"`
			Wed   bool `json:"wed"`
			Thu   bool `json:"thu"`
			Fri   bool `json:"fri"`
			Sat   bool `json:"sat"`
			Sun   bool `json:"sun"`
			Value int  `json:"value"`
		} `json:"week"` // Расписание
	}
)
