package types

// LBSegmentRef — ссылка на сегмент для поиска свободных подсетей.
type LBSegmentRef struct {
	IP   string `json:"ip"`   // IP-адрес сегмента
	Mask string `json:"mask"` // Маска сегмента
}
