package types

// LBPassTemplate — шаблон пароля.
type LBPassTemplate struct {
	CheckPass   int    `json:"check_pass"`   // Проверка пароля
	Descr       string `json:"descr"`        // Описание
	Excluded    string `json:"excluded"`     // Исключённые символы
	MinLength   int    `json:"min_length"`   // Минимальная длина
	NeedLow     int    `json:"need_low"`     // Требовать строчные
	NeedNumbers int    `json:"need_numbers"` // Требовать цифры
	NeedSpecial int    `json:"need_special"` // Требовать спецсимволы
	NeedUpper   int    `json:"need_upper"`   // Требовать заглавные
	ObjectType  string `json:"object_type"`  // Тип объекта
	RecordID    int    `json:"record_id"`    // ID записи
}
