package types

import (
	"encoding/json"
)

type LBAPINotification struct {
	Message string          `json:"message"`
	Params  json.RawMessage `json:"params"`
}

type Pagination struct {
	PgNum  int `json:"pg_num"`
	PgSize int `json:"pg_size"`
}

type LanbillingRequest interface {
	Method() string
}
type LanbillingResponse interface{}
