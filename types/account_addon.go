package types

type AccountAddon struct {
	Descr    string      `json:"descr"`
	Idx      interface{} `json:"idx"`
	Name     string      `json:"name"`
	StrValue string      `json:"str_value"`
	Type     int         `json:"type"`
	Uid      int         `json:"uid"`
}
