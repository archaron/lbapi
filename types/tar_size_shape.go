package types

type (
	TarSizeShapePeriod int

	// TarSizeShape Настройки полосы пропускания тарифа в зависимости от объема
	TarSizeShape struct {
		ID        int                `json:"id"`         // Идентификатор
		TarID     int                `json:"tar_id"`     // Идентификатор тарифа
		Amount    int                `json:"amount"`     // Объем, МБ
		ShapeRate float64            `json:"shape_rate"` // Полоса пропускания
		Period    TarSizeShapePeriod `json:"period"`     // Период
		TimeMark  LBTimeMark         `json:"time_mark"`  // Отметка времени
	}
)

func (t TarSizeShapePeriod) String() string {
	switch t {
	case TarSizeShapePeriodDay:
		return "day"
	case TarSizeShapePeriodWeek:
		return "week"
	case TarSizeShapePeriodMonth:
		return "month"
	default:
		return "unknown"
	}
}

const (
	TarSizeShapePeriodMonth TarSizeShapePeriod = iota
	TarSizeShapePeriodWeek
	TarSizeShapePeriodDay
)
