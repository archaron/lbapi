package types

// LBAgentFull — полная информация об агенте (метод getAgent).
type LBAgentFull struct {
	Active      bool            `json:"active"`
	Agent       LBAgent         `json:"agent"`
	IgnoreNets  []interface{}   `json:"ignore_nets"`
	Interfaces  []interface{}   `json:"interfaces"`
	Sessions    []interface{}   `json:"sessions"`
	Vgroups     []interface{}   `json:"vgroups"`
}
