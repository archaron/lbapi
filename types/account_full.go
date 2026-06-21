package types

// LBAccountFull — полная информация о пользователе (метод getAccount).
type LBAccountFull struct {
	AbonentName       string         `json:"abonent_name"`
	AbonentPatronymic string         `json:"abonent_patronymic"`
	AbonentSurname    string         `json:"abonent_surname"`
	ActOnWhat         string         `json:"act_on_what"`
	Addons            []interface{}  `json:"addons"`
	Addresses         []LBAddress    `json:"addresses"`
	Agreements        []interface{}  `json:"agreements"`
	AppID             *int           `json:"app_id"`
	BankName          string         `json:"bank_name"`
	Bik               string         `json:"bik"`
	BillDelivery      int            `json:"bill_delivery"`
	BillDeliveryName  string         `json:"bill_delivery_name"`
	BirthDate         string         `json:"birth_date"`
	BirthPlace        string         `json:"birth_place"`
	BranchBankName    string         `json:"branch_bank_name"`
	Category          int            `json:"category"`
	CorrAccount       string         `json:"corr_account"`
	Descr             string         `json:"descr"`
	Email             string         `json:"email"`
	Inn               string         `json:"inn"`
	IsArchive         bool           `json:"is_archive"`
	IsDefault         bool           `json:"is_default"`
	IsTemplate        int            `json:"is_template"`
	Kpp               string         `json:"kpp"`
	Login             string         `json:"login"`
	Mobile            string         `json:"mobile"`
	Name              string         `json:"name"`
	Ogrn              string         `json:"ogrn"`
	Okpo              string         `json:"okpo"`
	OrgUID            *int           `json:"org_uid"`
	OrganizationID    int            `json:"organization_id"`
	OrganizationName  string         `json:"organization_name"`
	PassportDate      string         `json:"passport_date"`
	PassportNum       string         `json:"passport_num"`
	PassportOrg       string         `json:"passport_org"`
	PaymentOrderNum   string         `json:"payment_order_num"`
	Phone             string         `json:"phone"`
	RS                string         `json:"rs"`
	SoleProprietor    *int           `json:"sole_proprietor"`
	Type              int            `json:"type"`
	UID               int            `json:"uid"`
}

// LBAddress — адрес пользователя.
type LBAddress struct {
	Address      string  `json:"address"`
	BuildingUUID *string `json:"building_uuid"`
	Code         string  `json:"code"`
	Type         int     `json:"type"`
}
