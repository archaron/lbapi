package types

type LBBlockRequest int

const (
	BlockRequestActivate = 0
	BlockRequestBlock    = 3
	BlockRequestOff      = 10
)

func (L LBBlockRequest) String() string {
	switch L {
	case BlockRequestActivate:
		return "Включить"
	case BlockRequestBlock:
		return "Заблокировать"
	case BlockRequestOff:
		return "Отключить"
	default:
		return "Unknown"
	}
}
