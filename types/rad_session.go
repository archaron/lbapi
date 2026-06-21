package types

// RadSession описывает Radius-сессию для методов radStartSessions/radStopSessionsPacket.
type RadSession struct {
	VgID             int    `json:"vg_id"`
	Class            string `json:"class"`
	Nas              string `json:"nas"`
	SessionID        string `json:"session_id"`
	ParentSessionID  string `json:"parent_session_id"`
	StartTime        string `json:"start_time,omitempty"`
	SessAni          string `json:"sess_ani,omitempty"`
	ID               int    `json:"id,omitempty"`
	StopReq          int    `json:"stop_req,omitempty"`
	Authmoment       string `json:"authmoment,omitempty"`
	Direction        int    `json:"direction,omitempty"`
	Updatetime       string `json:"updatetime,omitempty"`
	Shape            int    `json:"shape,omitempty"`
	VlanID           int    `json:"vlan_id,omitempty"`
	PortName         string `json:"port_name,omitempty"`
	Synchronized     int    `json:"synchronized,omitempty"`
	AssignedIP       string `json:"assigned_ip"`
	IPv4Static       int    `json:"ipv4_static,omitempty"`
}
