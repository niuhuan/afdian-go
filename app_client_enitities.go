package afdian

type AuthToken struct {
	AuthToken string `json:"auth_token"`
}

type MyAccount struct {
	Login         MyAccountLogin `json:"login"`
	UserPrivateId string         `json:"user_private_id"`
}

type MyAccountLogin struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Plans struct {
	List                  []CreatorPlan `json:"list"`
	SaleHasMore           int           `json:"sale_has_more"`
	SaleList              []CreatorPlan `json:"sale_list"`
	LimitShowProductCount int           `json:"limit_show_product_count"`
}

type CreatorPlan struct {
	CanAliAgreement      int               `json:"can_ali_agreement"`
	PlanId               string            `json:"plan_id"`
	Rank                 int               `json:"rank"`
	UserId               string            `json:"user_id"`
	Status               int               `json:"status"`
	Name                 string            `json:"name"`
	Pic                  string            `json:"pic"`
	Desc                 string            `json:"desc"`
	Price                string            `json:"price"`
	UpdateTime           int               `json:"update_time"`
	Timing               CreatorPlanTiming `json:"timing"`
	PayMonth             int               `json:"pay_month"`
	ShowPrice            string            `json:"show_price"`
	ShowPriceAfterAdjust string            `json:"show_price_after_adjust"`
	HasCoupon            int               `json:"has_coupon"`
	FavorablePrice       int               `json:"favorable_price"`
	Independent          int               `json:"independent"`
	Permanent            int               `json:"permanent"`
	CanBuyHide           int               `json:"can_buy_hide"`
	NeedAddress          int               `json:"need_address"`
	ProductType          int               `json:"product_type"`
	SaleLimitCount       int               `json:"sale_limit_count"`
	NeedInviteCode       bool              `json:"need_invite_code"`
	BundleStock          int               `json:"bundle_stock"`
	BundleSkuSelectCount int               `json:"bundle_sku_select_count"`
	HasPlanConfig        int               `json:"has_plan_config"`
	HasBadge             int               `json:"has_badge"`
	SponsorCount         string            `json:"sponsor_count"`
	HasVip               int               `json:"has_vip"`
}

type CreatorPlanTiming struct {
	TimingOn      int `json:"timing_on"`
	TimingOff     int `json:"timing_off"`
	TimingSellOn  int `json:"timing_sell_on"`
	TimingSellOff int `json:"timing_sell_off"`
}

type PlanSkus struct {
	Plan       CreatorPlan   `json:"plan"`
	List       []Sku         `json:"list"`
	BoughtSkus []interface{} `json:"bought_skus"`
}

type Sku struct {
	SkuId               string    `json:"sku_id"`
	PlanId              string    `json:"plan_id"`
	UserId              string    `json:"user_id"`
	Status              int       `json:"status"`
	Name                string    `json:"name"`
	Pic                 string    `json:"pic"`
	Desc                string    `json:"desc"`
	Stock               string    `json:"stock"`
	SponsorCount        string    `json:"sponsor_count"`
	Price               string    `json:"price"`
	ReplyContent        string    `json:"reply_content"`
	ReplyRandomContent  string    `json:"reply_random_content"`
	ReplyRandomNum      int       `json:"reply_random_num"`
	ReplySwitch         bool      `json:"reply_switch"`
	ReplyRandomSwitch   bool      `json:"reply_random_switch"`
	PurchasedRichText   string    `json:"purchased_rich_text"`
	Redeem              SkuRedeem `json:"redeem"`
	Plan                SkuPlan   `json:"plan"`
	HasSkuPurchaseLimit int       `json:"has_sku_purchase_limit"`
	CanBuyCount         int       `json:"can_buy_count"`
}

type SkuPlan struct {
	Id                   int    `json:"id"`
	PlanId               string `json:"plan_id"`
	UserId               string `json:"user_id"`
	Status               int    `json:"status"`
	Rank                 int    `json:"rank"`
	Name                 string `json:"name"`
	Pic                  string `json:"pic"`
	Desc                 string `json:"desc"`
	Price                string `json:"price"`
	PayMonth             int    `json:"pay_month"`
	ReplySwitch          int    `json:"reply_switch"`
	ReplyContent         string `json:"reply_content"`
	ReplyRandomSwitch    int    `json:"reply_random_switch"`
	ReplyRandomContent   string `json:"reply_random_content"`
	Independent          int    `json:"independent"`
	Permanent            int    `json:"permanent"`
	ProductType          int    `json:"product_type"`
	CanBuyHide           int    `json:"can_buy_hide"`
	NeedAddress          int    `json:"need_address"`
	CreateTime           int    `json:"create_time"`
	UpdateTime           int    `json:"update_time"`
	BundleSkuSelectCount int    `json:"bundle_sku_select_count"`
	BundleStock          int    `json:"bundle_stock"`
	CanAliAgreement      int    `json:"can_ali_agreement"`
}

type SkuRedeem struct {
	TotalCount   int `json:"total_count"`
	PendingCount int `json:"pending_count"`
	UsedCount    int `json:"used_count"`
}
