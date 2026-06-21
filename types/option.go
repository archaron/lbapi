package types

// LBOption — опция LANBilling.
type LBOption struct {
	AvailabilityType int    `json:"availability_type"` // Тип доступности
	Descr            string `json:"descr"`             // Описание
	Name             string `json:"name"`              // Имя опции
	Value            string `json:"value"`             // Значение
}
