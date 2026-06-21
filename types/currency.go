package types

// LBCurrency — валюта в LANBilling.
type LBCurrency struct {
	CodeOkv *int   `json:"code_okv"` // Код ОКВ
	CurID   int    `json:"cur_id"`   // ID валюты
	IsDef   bool   `json:"is_def"`   // Валюта по умолчанию
	Name    string `json:"name"`     // Название валюты
	Symbol  string `json:"symbol"`   // Символ валюты
}
