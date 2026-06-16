package types

type AccountsAddonsSetValue struct {
	Idx   int    `json:"idx"`
	Value string `json:"value"`
}

type AccountsAddonsSet struct {
	Descr         string                   `json:"descr"`
	ForClient     int                      `json:"for_client"`
	Name          string                   `json:"name"`
	RequiredField bool                     `json:"required_field"`
	ShowOnHp      bool                     `json:"show_on_hp"`
	Type          int                      `json:"type"`
	Values        []AccountsAddonsSetValue `json:"values"`
}
