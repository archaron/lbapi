package types

type RadiusAttrs struct {
	Id          int     `json:"id"`
	AgentDescr  string  `json:"agent_descr"`
	AttrId      int     `json:"attr_id"`
	CatDescr    *string `json:"cat_descr"`
	CatIdx      *int    `json:"cat_idx"`
	DaeOnEvent  int     `json:"dae_on_event"`
	Description string  `json:"description"`
	DevGroupId  *int    `json:"dev_group_id"`
	DictName    string  `json:"dict_name"`
	GroupId     *int    `json:"group_id"`
	GroupName   *string `json:"group_name"`

	NasId          *int       `json:"nas_id"`
	OwnerDescr     *string    `json:"owner_descr"`
	RadiusCode     int        `json:"radius_code"`
	RecordId       int        `json:"record_id"`
	SegmentId      *int       `json:"segment_id"`
	SegmentIp      *string    `json:"segment_ip"`
	Service        *string    `json:"service"`
	ServiceForList int        `json:"service_for_list"`
	Shape          int        `json:"shape"`
	Tag            int        `json:"tag"`
	TarId          *int       `json:"tar_id"`
	TarifDescr     *string    `json:"tarif_descr"`
	TimeMark       LBTimeMark `json:"time_mark"`
	Value          string     `json:"value"`
	VgId           *int       `json:"vg_id"`
	VgroupLogin    *string    `json:"vgroup_login"`
}
