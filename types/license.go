package types

// LBLicense - Информация о лицензии
type LBLicense struct {
	Cdkey               string `json:"cdkey"`
	Expire              int    `json:"expire"`
	GenDate             string `json:"gen_date"`
	MonoCaptivePortals  int    `json:"mono_captive_portals"`
	MultiCaptivePortals int    `json:"multi_captive_portals"`
	Payments            int    `json:"payments"`
	ProgName            string `json:"prog_name"`
	ProgOwner           string `json:"prog_owner"`
	SerialNumber        string `json:"serial_number"`
	UseAsterisk         bool   `json:"use_asterisk"`
	UseBankDetails      bool   `json:"use_bank_details"`
	UseCerber           bool   `json:"use_cerber"`
	UseEc               bool   `json:"use_ec"`
	UseFidelio          bool   `json:"use_fidelio"`
	UseFiscalization    bool   `json:"use_fiscalization"`
	UseInet             int    `json:"use_inet"`
	UseInventory        bool   `json:"use_inventory"`
	UseIptvportal       bool   `json:"use_iptvportal"`
	UseLbtv             bool   `json:"use_lbtv"`
	UseLifestream       bool   `json:"use_lifestream"`
	UseMegogo           bool   `json:"use_megogo"`
	UseMinistra         bool   `json:"use_ministra"`
	UseMoovi            bool   `json:"use_moovi"`
	UseNexttv           bool   `json:"use_nexttv"`
	UseOperators        bool   `json:"use_operators"`
	UsePhone            int    `json:"use_phone"`
	UseSoftline         bool   `json:"use_softline"`
	UseStarcor          bool   `json:"use_starcor"`
	UseTv24             bool   `json:"use_tv24"`
	UseTvip             bool   `json:"use_tvip"`
	UseUsbox            int    `json:"use_usbox"`
	UserLimit           int    `json:"user_limit"`
}
