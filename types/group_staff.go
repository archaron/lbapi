package types

// GroupStaff - принадлежность учетной записи группам учетных записей
type GroupStaff struct {
	GroupID  int        `json:"group_id"`  // Идентификатор группы учетных записей
	TimeMark LBTimeMark `json:"time_mark"` // Отметка времени
	VgID     int        `json:"vg_id"`     // Идентификатор учетной записи в группе
}
