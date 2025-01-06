package types

type LBAccount struct {
	UID            int           `json:"uid"`
	Category       int           `json:"category"`
	Agreements     []interface{} `json:"agreements"`
	IsDefault      bool          `json:"is_default"`
	Name           string        `json:"name"`
	Login          string        `json:"login"`
	Email          string        `json:"email"`
	IsTemplate     int           `json:"is_template"`
	Mobile         string        `json:"mobile"`
	Phone          string        `json:"phone"`
	AppID          int           `json:"app_id"`
	Descr          string        `json:"descr"`
	IsArchive      bool          `json:"is_archive"`
	SoleProprietor bool          `json:"sole_proprietor"`
	Type           int           `json:"type"`
	Addresses      interface{}   `json:"address"`
}
