package types

// RadStopSession описывает параметры остановки Radius-сессии для метода radStopSessionsPacket.
type RadStopSession struct {
	SessionID  string `json:"session_id"`
	Nas        string `json:"nas"`
	StopTime   string `json:"stop_time"`
	AssignedIP string `json:"assigned_ip"`
	AgentID    int    `json:"agent_id"`
	Class      string `json:"class"`
}
