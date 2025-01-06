package types

type TarCatSizeDiscount struct {
	Type     int        `json:"type"`      // Тип скидки (%/руб)
	Amount   float32    `json:"amount"`    // Объем, МБ
	Bonus    float32    `json:"bonus"`     // Бонус, руб
	CatIDX   int        `json:"cat_idx"`   // Идентификатор каталога
	DisID    int        `json:"dis_id"`    // Идентификатор скидки
	Discount float32    `json:"discount"`  // Скидка
	TarID    int        `json:"tar_id"`    // Идентификатор тарифа
	TimeMark LBTimeMark `json:"time_mark"` // Отметка времени
}
