package types

import (
	"fmt"
	"net"
)

type LBSyncStaff struct {
	Netmask   string         `json:"netmask"`
	Network   string         `json:"network"`
	RecordID  int            `json:"record_id"`
	SegmentID int            `json:"segment_id"`
	TimeMark  LBTimeMark     `json:"time_mark"`
	Type      StaffBlockType `json:"type"`
	VgID      int            `json:"vg_id"`
}

func (L *LBSyncStaff) String() string {
	ip := L.ToIPNet()
	maskLen, _ := ip.Mask.Size()
	return fmt.Sprintf("%s/%d", ip.IP, maskLen)
}

func (L *LBSyncStaff) ToIPNet() net.IPNet {
	ip := net.ParseIP(L.Network).To4()
	mask := net.ParseIP(L.Netmask).To4()
	return net.IPNet{
		IP:   ip,
		Mask: net.IPMask(mask),
	}
}
