package types

// LBStatus — статус сервера LANBilling.
type LBStatus struct {
	Accounts    int    `json:"accounts"`
	AgentsStop  int    `json:"agents_stop"`
	AgentsWork  int    `json:"agents_work"`
	Agreements  int    `json:"agreements"`
	BuildDate   string `json:"build_date"`
	BuildName   string `json:"build_name"`
	BuildRev    string `json:"build_rev"`
	CrmDocs     int    `json:"crm_docs"`
	Currencies  int    `json:"currencies"`
	DbBuildDate string `json:"db_build_date"`
	DbBuildRev  string `json:"db_build_rev"`
	Devices     int    `json:"devices"`
	Managers    int    `json:"managers"`
	Operators   int    `json:"operators"`
	PayManagers int    `json:"pay_managers"`
	Tickets     int    `json:"tickets"`
	Vgroups     int    `json:"vgroups"`
}
