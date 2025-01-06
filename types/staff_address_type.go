package types

type StaffAddressType int

const (
	StaffAddressTypeIPv4 StaffAddressType = 2
	StaffAddressTypeIPv6 StaffAddressType = 10
)

func (s StaffAddressType) String() string {
	switch s {
	case StaffAddressTypeIPv4:
		return "IPv4"
	case StaffAddressTypeIPv6:
		return "IPv6"
	default:
		return "unknown"
	}
}
