package types

// PayCard - данные о платежной карте
type PayCard struct {
	ActTil        string     `json:"act_til"`         // Активировать до
	ActivateDate  *LBTime    `json:"activate_date"`   // Дата активации
	AgrmId        *int       `json:"agrm_id"`         // Идентификатор договора
	AgrmNumber    *string    `json:"agrm_number"`     // Номер договора
	CardKey       string     `json:"card_key"`        // Ключ карты
	CreateDate    *LBTime    `json:"create_date"`     // Дата создания
	CurId         int        `json:"cur_id"`          // Идентификатор валюты
	ExpirePeriod  int        `json:"expire_period"`   // Скор действия
	ModPerson     int        `json:"mod_person"`      // Идентификатор менеджера создавшего
	ModPersonName string     `json:"mod_person_name"` // Имя менеджера создавшего
	SerNo         int        `json:"ser_no"`          // Серийный номер карты
	SetId         int        `json:"set_id"`          // Идентификатор набора карт
	SetName       string     `json:"set_name"`        // Название набора карт
	Sum           float64    `json:"sum"`             // Сумма карты
	Symbol        string     `json:"symbol"`          // Символ валютыс
	TimeMark      LBTimeMark `json:"time_mark"`       // Временная отметка
	Uid           *int       `json:"uid"`             // Идентификатор пользователя
	Used          int        `json:"used"`            // Активирована
	UserName      *string    `json:"user_name"`       // Имя пользователя
}
