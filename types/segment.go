package types

import (
	"fmt"
)

type Segment struct {
	AgentID         int         `json:"agent_id"`
	DelegatedPrefix interface{} `json:"delegated_prefix"`
	Descr           string      `json:"descr"`
	DeviceGroupID   interface{} `json:"device_group_id"`
	DeviceGroupName interface{} `json:"device_group_name"`
	Gateway         interface{} `json:"gateway"`
	Guest           int         `json:"guest"`
	IgnoreLocal     int         `json:"ignore_local"`
	IP              string      `json:"ip"`
	Mask            string      `json:"mask"`
	NasID           int         `json:"nas_id"`
	Nat             int         `json:"nat"`
	OuterVlan       *int        `json:"outer_vlan"`
	Priority        int         `json:"priority"`
	Rnas            interface{} `json:"rnas"`
	SegmentID       int         `json:"segment_id"`
	TimeMark        LBTimeMark  `json:"time_mark"`
	Type            int         `json:"type"`
	Unavailable     int         `json:"unavailable"`
	VlanID          *int        `json:"vlan_id"`
	VlanName        *string     `json:"vlan_name"`
}

func (s Segment) String() string {
	return fmt.Sprintf("[%3d]\t%-15s\t%-15s\t%s", s.SegmentID, s.IP, s.Mask, s.Descr)
}
