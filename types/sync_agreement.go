package types

type (
	SyncAgreement struct {
		AgrmID        int        `json:"agrm_id"` // Идентификатор договора
		Balance       float64    `json:"balance"` // Баланс договора
		Credit        float64    `json:"credit"`  // Доступный кредит
		Installments  float64    `json:"installments"`
		Number        string     `json:"number"`         // Номер договора
		PromiseCredit float64    `json:"promise_credit"` // Кредит обещанного платежа
		TimeMark      LBTimeMark `json:"time_mark"`
	}

	SyncAgreementExt struct {
		AgrmID        int        `json:"agrm_id"`
		PaymentMethod int        `json:"payment_method"`
		TimeMark      LBTimeMark `json:"time_mark"`
	}
)
