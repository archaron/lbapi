package types

// LBSort — параметр сортировки для запросов.
type LBSort struct {
	Ascdesc int    `json:"ascdesc"` // 0=ASC, 1=DESC
	Name    string `json:"name"`   // Имя поля для сортировки
}
