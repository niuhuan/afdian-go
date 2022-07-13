package afdian

type AfdianResponse[T any] struct {
	Ec   int    `json:"ec"`
	Em   string `json:"em"`
	Data T      `json:"data"`
}

type PageData[T any] struct {
	List       []T     `json:"list"`
	TotalCount int     `json:"total_count"`
	TotalPage  int     `json:"total_page"`
	Request    Request `json:"request"`
}

type Order struct {
	OutTradeNo     string `json:"out_trade_no"`
	UserId         string `json:"user_id"`
	PlanId         string `json:"plan_id"`
	Month          int    `json:"month"`
	TotalAmount    string `json:"total_amount"`
	ShowAmount     string `json:"show_amount"`
	Status         int    `json:"status"`
	Remark         string `json:"remark"`
	RedeemId       string `json:"redeem_id"`
	ProductType    int    `json:"product_type"`
	Discount       string `json:"discount"`
	SkuDetail      []Sku  `json:"sku_detail"`
	UserPrivateId  string `json:"user_private_id"`
	AddressPerson  string `json:"address_person"`
	AddressPhone   string `json:"address_phone"`
	AddressAddress string `json:"address_address"`
}

type Sku struct {
	SkuId   string `json:"sku_id"`
	Count   int    `json:"count"`
	Name    string `json:"name"`
	AlbumId string `json:"album_id"`
	Pic     string `json:"pic"`
}

type Request struct {
	UserId string `json:"user_id"`
	Params string `json:"params"`
	Ts     int    `json:"ts"`
	Sign   string `json:"sign"`
}

type Sponsor struct {
	AllSumAmount string `json:"all_sum_amount"`
	CurrentPlan  Plan   `json:"current_plan"`
	FirstPayTime int    `json:"first_pay_time"`
	LastPayTime  int    `json:"last_pay_time"`
	SponsorPlans []Plan `json:"sponsor_plans"`
	User         User   `json:"user"`
}

type User struct {
	Avatar        string `json:"avatar"`
	Name          string `json:"name"`
	UserId        string `json:"user_id"`
	UserPrivateId string `json:"user_private_id"`
}

type Plan struct {
	BundleSkuSelectCount int           `json:"bundle_sku_select_count"`
	BundleStock          int           `json:"bundle_stock"`
	CanBuyHide           int           `json:"can_buy_hide"`
	Coupon               []interface{} `json:"coupon"`
	Desc                 string        `json:"desc"`
	ExpireTime           int           `json:"expire_time"`
	FavorablePrice       int           `json:"favorable_price"`
	HasCoupon            int           `json:"has_coupon"`
	Independent          int           `json:"independent"`
	Name                 string        `json:"name"`
	NeedAddress          int           `json:"need_address"`
	NeedInviteCode       bool          `json:"need_invite_code"`
	PayMonth             int           `json:"pay_month"`
	Permanent            int           `json:"permanent"`
	Pic                  string        `json:"pic"`
	PlanId               string        `json:"plan_id"`
	Price                string        `json:"price"`
	ProductType          int           `json:"product_type"`
	Rank                 int           `json:"rank"`
	RankType             int           `json:"rankType"`
	SaleLimitCount       int           `json:"sale_limit_count"`
	ShowPrice            string        `json:"show_price"`
	ShowPriceAfterAdjust string        `json:"show_price_after_adjust"`
	SkuProcessed         []interface{} `json:"sku_processed"`
	Status               int           `json:"status"`
	Timing               struct {
		TimingOff int `json:"timing_off"`
		TimingOn  int `json:"timing_on"`
	} `json:"timing"`
	UpdateTime int    `json:"update_time"`
	UserId     string `json:"user_id"`
}

//////////////////// 爱发电调服务器相关 ////////////////////

type AfdianCall AfdianResponse[AfdianCallData]

type AfdianCallData struct {
	Type  string `json:"type"`
	Order Order  `json:"order"`
}

// AfdianCallResponse
// 爱发电回调返回
// {"ec":200,"em":""}
type AfdianCallResponse struct {
	Ec int    `json:"ec"`
	Em string `json:"em"`
}
