package main

type RequestMsg struct {
	PreVid string
}

// ResponseMsg 应答参数
type ResponseMsg struct {
	// 网吧免广告(是否)
	IsAdKeyNoAd bool
	// 政治会议免广告 (是否)
	IsConferenceNoAid bool
	// 是否是商业化免广告
	IsCommercialAdFree bool
	// 体育会员
	IsSportVip bool
	// TV会员
	NoAid bool
}
