package types

// LBTrustedNetwork — доверенная сеть (IP, с которого разрешён доступ к биллингу).
type LBTrustedNetwork struct {
	Descr    string `json:"descr"`     // Описание
	IP       string `json:"ip"`        // IP-адрес
	Mask     string `json:"mask"`      // Маска сети
	RecordID int    `json:"record_id"` // ID записи
}
