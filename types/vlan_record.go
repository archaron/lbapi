package types

type VlanRecord struct {
	VlanId    int      `json:"vlan_id,omitempty"`
	InnerVlan int      `json:"inner_vlan"`
	OuterVlan int      `json:"outer_vlan"`
	Name      string   `json:"name"`
	Type      VlanType `json:"type"`
}
