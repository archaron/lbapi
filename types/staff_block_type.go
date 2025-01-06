package types

type StaffBlockType int

const (
	StaffBlockTypeIPv4          StaffBlockType = 0
	StaffBlockTypeIPv6Delegated StaffBlockType = 1
	StaffBlockTypeIPv6Framed    StaffBlockType = 2
)

func (s StaffBlockType) String() string {
	switch s {
	case StaffBlockTypeIPv4:
		return "IPv4 block"
	case StaffBlockTypeIPv6Delegated:
		return "IPv6 delegated block"
	case StaffBlockTypeIPv6Framed:
		return "IPv6 framed block"
	default:
		return "unknown"
	}
}
