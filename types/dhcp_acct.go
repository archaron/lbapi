package types

// DhcpAcctParams описывает параметры DHCP-аккаунтинга (start/stop).
type DhcpAcctParams struct {
	SessionID  string `json:"session_id"`  // ID сессии
	Nas        string `json:"nas"`         // Адрес NAS-сервера
	AssignedIP string `json:"assigned_ip"` // Назначенный IP-адрес
	Class      string `json:"class"`       // Класс DHCP-сессии
}
