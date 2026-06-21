package types

// LBTelClass — телефонный класс (ВТС, МГ, МН и т.д.).
type LBTelClass struct {
	ClassID    int    `json:"class_id"`    // ID класса
	Color      int    `json:"color"`       // Цвет в интерфейсе
	Descr      string `json:"descr"`       // Описание
	ExternCode string `json:"extern_code"` // Внешний код
	Mgmn       bool   `json:"mgmn"`        // Междугородняя/международная связь
	Name       string `json:"name"`        // Название класса
}
