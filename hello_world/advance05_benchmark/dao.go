package advance05_benchmark

type WatchBuy struct {
	Time    int `json:"time"`
	Product int `json:"product"`
	Anchor  int `json:"anchor"`
	PosX    int `json:"pos_x"`
	PoxY    int `json:"pos_y"`
}

//go:generate
type AidCache struct {
	Aid      uint32 `json:"aid"`       // 广告位
	TimeLen  uint32 `json:"time_len"`  // 视频时长
	AdBreak1 uint32 `json:"ad_break1"` // 中插点位1
	AdBreak2 uint32 `json:"ad_break2"` // 中插点位2
	AdBreak3 uint32 `json:"ad_break3"` // 中插点位3
	AdBreak4 uint32 `json:"ad_break4"` // 中插点位4
	AdBreak5 uint32 `json:"ad_break5"` // 中插点位5
	AdFlag   int    `json:"ad_flag"`   // 华纳 HBO标志
	TypeId   uint32 `json:"type_id"`   // 大分类id
	ColumnId int    `json:"column_id"` // 栏目id

	IvbBreaks []uint32   `json:"ivb_breaks"` // IVB中插
	UploadQQ  uint32     `json:"upload_qq"`  // 上传者qq号
	WatchBuys []WatchBuy `json:"watch_buys"` // 边看边买中插
	ImgTag    []uint32   `json:"img_tag"`    // 角标广告中插点

	CidStatus    uint32 `json:"cid_status"`     // 专辑的付费类型
	NoAdStatus   uint32 `json:"no_ad_status"`   // 媒资信息中配置的免广告逻辑
	VcomDealFlag uint32 `json:"vcom_deal_flag"` // 商业化免广告标记
	PayStatus    uint32 `json:"pay_status"`     // 视频付费类型
	ResourceType int32  `json:"resource_type"`  // 视频资源类型 1-NBA视频
}

func NewAidCache() *AidCache {
	return &AidCache{
		Aid:          23,
		TimeLen:      100,
		AdBreak1:     11000,
		AdBreak2:     11000,
		AdBreak3:     3,
		AdBreak4:     4,
		AdBreak5:     5,
		AdFlag:       0,
		TypeId:       0,
		ColumnId:     0,
		IvbBreaks:    []uint32{1, 2, 3, 4, 5},
		UploadQQ:     0,
		WatchBuys:    []WatchBuy{{
			Time:    1,
			Product: 2,
			Anchor:  3,
			PosX:    4,
			PoxY:    5,
		}},
		ImgTag:       []uint32{10, 20, 30, 40, 50},
		CidStatus:    1,
		NoAdStatus:   2,
		VcomDealFlag: 1,
		PayStatus:    1,
		ResourceType: 1,
	}
}
