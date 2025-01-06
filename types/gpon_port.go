package types

type GPONPort struct {
	ID        int        `json:"id"`
	PortID    int        `json:"port_id"`
	TimeMark  LBTimeMark `json:"time_mark"`
	VgID      int        `json:"vg_id"`
	VirtualID string     `json:"virtual_id"`
}
