package types

type VlanType int

const (
	VlanTypeClient VlanType = 1
	VlanTypeDevice VlanType = 2
	VlanTypeGuest  VlanType = 3

	VlanTypeClientName = "client"
	VlanTypeDeviceName = "device"
	VlanTypeGuestName  = "guest"
)

func (v VlanType) String() string {
	switch v {
	case VlanTypeClient:
		return VlanTypeClientName
	case VlanTypeDevice:
		return VlanTypeDeviceName
	case VlanTypeGuest:
		return VlanTypeGuestName
	default:
		return "Unknown"
	}
}
