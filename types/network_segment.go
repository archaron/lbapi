package types

import (
	"fmt"
)

type LBNetworkSegment struct {
	Network     string `json:"network"`
	Netmask     string `json:"netmask"`
	SegmentType string `json:"segment_type"`
}

func (L LBNetworkSegment) String() string {
	return fmt.Sprintf("%s %s/%s", L.SegmentType, L.Network, L.Netmask)
}
