package types

type Agreement struct {
	Addons               []interface{} `json:"addons"`
	AgrmId               int           `json:"agrm_id"`
	AgrmNum              string        `json:"agrm_num"`
	BCheck               interface{}   `json:"b_check"`
	BLimit               float64       `json:"b_limit"`
	BNotify              int           `json:"b_notify"`
	Balance              float64       `json:"balance"`
	BalanceAcc           float64       `json:"balance_acc"`
	BalanceLimitExceeded interface{}   `json:"balance_limit_exceeded"`
	BalanceStatus        int           `json:"balance_status"`
	BalanceStrictLimit   float64       `json:"balance_strict_limit"`
	BlockAmount          interface{}   `json:"block_amount"`
	BlockDays            interface{}   `json:"block_days"`
	BlockMonths          interface{}   `json:"block_months"`
	BlockOrders          interface{}   `json:"block_orders"`
	BoposAgrmId          string        `json:"bopos_agrm_id"`
	CloseDate            *LBTime       `json:"close_date"`  // Дата расторжения
	CreateDate           *LBTime       `json:"create_date"` // Дата заключения
	Credit               float64       `json:"credit"`
	CurId                int           `json:"cur_id"`
	Datevalidto          interface{}   `json:"datevalidto"`
	Descr                interface{}   `json:"descr"`
	ErrorMessage         string        `json:"error_message"`
	FriendAgrmId         interface{}   `json:"friend_agrm_id"`
	FriendNumber         interface{}   `json:"friend_number"`
	Installments         float64       `json:"installments"`
	IsArchive            bool          `json:"is_archive"`
	IsAuto               bool          `json:"is_auto"`
	IsDefault            int           `json:"is_default"`
	ManagerId            interface{}   `json:"manager_id"`
	ManagerLogin         interface{}   `json:"manager_login"`
	ManagerName          interface{}   `json:"manager_name"`
	MonthBlockDay        int           `json:"month_block_day"`
	NoFinBlock           int           `json:"no_fin_block"`
	OperId               int           `json:"oper_id"`
	OperName             string        `json:"oper_name"`
	OrderPayDay          interface{}   `json:"order_pay_day"`
	OrgAgrmId            interface{}   `json:"org_agrm_id"`
	OrganizationId       int           `json:"organization_id"`
	OrganizationName     string        `json:"organization_name"`
	OwnerId              int           `json:"owner_id"`
	ParentAgrmId         interface{}   `json:"parent_agrm_id"`
	ParentNumber         interface{}   `json:"parent_number"`
	PayCode              interface{}   `json:"pay_code"`
	PaymentMethod        int           `json:"payment_method"`
	PenaltyMethod        int           `json:"penalty_method"`
	Priority             int           `json:"priority"`
	PromiseCredit        float64       `json:"promise_credit"`
	Symbol               string        `json:"symbol"`
	TimeMark             interface{}   `json:"time_mark"`
	Uid                  int           `json:"uid"`
	UserName             string        `json:"user_name"`
	Vgroups              int           `json:"vgroups"`
}
