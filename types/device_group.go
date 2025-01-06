package types

// LBDeviceGroup Группа операторского оборудования
type LBDeviceGroup struct {
	GroupID   int    `json:"group_id"` //Идентификатор группы
	Name      string `json:"name"`     // Название группы
	Desc      string `json:"desc"`     //  Краткая характеристика о группе устройств.
	AgentID   int    `json:"agent_id"` // Общий сетевой агент SNMP, который осуществляет мониторинг работы оборудования.
	AgentName string `json:"agent_name"`

	TimeMark LBTimeMark `json:"time_mark"`
}
