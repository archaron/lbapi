package types

// LBAgreement — информация о договоре (метод getAgreements).
type LBAgreement struct {
	Addons               []interface{} `json:"addons"`
	AgrmID               int           `json:"agrm_id"`
	AgrmNum              string        `json:"agrm_num"`
	BCheck               *int          `json:"b_check"`
	BLimit               float64       `json:"b_limit"`
	BNotify              int           `json:"b_notify"`
	Balance              float64       `json:"balance"`
	BalanceAcc           float64       `json:"balance_acc"`
	BalanceLimitExceeded *bool         `json:"balance_limit_exceeded"`
	BalanceStatus        int           `json:"balance_status"`
	BalanceStrictLimit   float64       `json:"balance_strict_limit"`
	BlockAmount          *float64      `json:"block_amount"`
	BlockDays            *int          `json:"block_days"`
	BlockMonths          *int          `json:"block_months"`
	BlockOrders          *int          `json:"block_orders"`
	BoposAgrmID          string        `json:"bopos_agrm_id"`
	CloseDate            *string       `json:"close_date"`
	CreateDate           string        `json:"create_date"`
	Credit               float64       `json:"credit"`
	CurID                int           `json:"cur_id"`
	CurrSymbol           string        `json:"curr_symbol"`
	DateFrom             string        `json:"date_from"`
	DateTo               *string       `json:"date_to"`
	Installments         float64       `json:"installments"`
	IsClosed             bool          `json:"is_closed"`
	ManagerID            *int          `json:"manager_id"`
	ManagerLogin         *string       `json:"manager_login"`
	ManagerName          *string       `json:"manager_name"`
	OrganizationID       int           `json:"organization_id"`
	OrganizationName     string        `json:"organization_name"`
	PayClassID           int           `json:"pay_class_id"`
	PaymentMethod        int           `json:"payment_method"`
	PromiseCredit        float64       `json:"promise_credit"`
	TimeMark             *LBTimeMark   `json:"time_mark"`
	UID                  int           `json:"uid"`
}
