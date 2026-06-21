package types

// LBEpsAgreementParam — параметр договора для ЕПС.
type LBEpsAgreementParam struct {
	AgrmID int    `json:"agrm_id"`
	Param  string `json:"param"`
	Value  string `json:"value"`
}
