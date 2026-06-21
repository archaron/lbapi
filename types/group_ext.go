package types

// LBGroupExt — расширенная информация о группе учётных записей.
type LBGroupExt struct {
	AgentIDs []int  `json:"agent_ids"` // ID агентов группы
	Descr    string `json:"descr"`     // Описание группы
	GroupID  int    `json:"group_id"`  // ID группы
	Name     string `json:"name"`      // Название группы
	VgCount  int    `json:"vg_count"`  // Количество УЗ в группе
}
