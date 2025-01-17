// Code generated by goctl. DO NOT EDIT.
package types

type AddMemberAddressReq struct {
	MemberId      int64  `json:"memberId"`
	Name          string `json:"name"` // 收货人名称
	PhoneNumber   string `json:"phoneNumber"`
	DefaultStatus int64  `json:"defaultStatus"` // 是否为默认
	PostCode      string `json:"postCode"`      // 邮政编码
	Province      string `json:"province"`      // 省份/直辖市
	City          string `json:"city"`          // 城市
	Region        string `json:"region"`        // 区
	DetailAddress string `json:"detailAddress"` // 详细地址(街道)
}

type AddMemberAddressResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListMemberAddressReq struct {
	Current  int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}

type ListMemberAddressData struct {
	Id            int64  `json:"id"`
	MemberId      int64  `json:"memberId"`
	Name          string `json:"name"` // 收货人名称
	PhoneNumber   string `json:"phoneNumber"`
	DefaultStatus int64  `json:"defaultStatus"` // 是否为默认
	PostCode      string `json:"postCode"`      // 邮政编码
	Province      string `json:"province"`      // 省份/直辖市
	City          string `json:"city"`          // 城市
	Region        string `json:"region"`        // 区
	DetailAddress string `json:"detailAddress"` // 详细地址(街道)
}

type ListMemberAddressResp struct {
	Current  int64                   `json:"current,default=1"`
	Data     []ListMemberAddressData `json:"data"`
	PageSize int64                   `json:"pageSize,default=20"`
	Success  bool                    `json:"success"`
	Total    int64                   `json:"total"`
	Code     int64                   `json:"code"`
	Message  string                  `json:"message"`
}

type UpdateMemberAddressReq struct {
	Id            int64  `json:"id"`
	MemberId      int64  `json:"memberId"`
	Name          string `json:"name"` // 收货人名称
	PhoneNumber   string `json:"phoneNumber"`
	DefaultStatus int64  `json:"defaultStatus"` // 是否为默认
	PostCode      string `json:"postCode"`      // 邮政编码
	Province      string `json:"province"`      // 省份/直辖市
	City          string `json:"city"`          // 城市
	Region        string `json:"region"`        // 区
	DetailAddress string `json:"detailAddress"` // 详细地址(街道)
}

type UpdateMemberAddressResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DeleteMemberAddressReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteMemberAddressResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type LoginReq struct {
	Username string `json:"usernam,optional"` //会员名称
	Password string `json:"password"`         //密码
	Mobile   string `json:"mobile,optional"`  //手机号码
}

type LoginResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"` //token
}

type RegisterReq struct {
	Username string `json:"username"` //会员名称
	Password string `json:"password"` //密码
	Mobile   string `json:"mobile"`   //手机号码
}

type RegisterResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"` //token
}

type InfoReq struct {
	Username string `json:"username,optional"` //会员名称
}

type InfoResp struct {
	Code    int64      `json:"code"`
	Message string     `json:"message"`
	Data    MemberData `json:"data"`
}

type MemberData struct {
	Id                    int64  `json:"id"`
	MemberLevelId         int64  `json:"memberLevelId"`
	Username              string `json:"username"`              // 用户名
	Nickname              string `json:"nickname"`              // 昵称
	Phone                 string `json:"phone"`                 // 手机号码
	Status                int64  `json:"status"`                // 帐号启用状态:0->禁用；1->启用
	CreateTime            string `json:"createTime"`            // 注册时间
	Icon                  string `json:"icon"`                  // 头像
	Gender                int64  `json:"gender"`                // 性别：0->未知；1->男；2->女
	Birthday              string `json:"birthday"`              // 生日
	City                  string `json:"city"`                  // 所做城市
	Job                   string `json:"job"`                   // 职业
	PersonalizedSignature string `json:"personalizedSignature"` // 个性签名
	SourceType            int64  `json:"sourceType"`            // 用户来源
	Integration           int64  `json:"integration"`           // 积分
	Growth                int64  `json:"growth"`                // 成长值
	LuckeyCount           int64  `json:"luckeyCount"`           // 剩余抽奖次数
	HistoryIntegration    int64  `json:"historyIntegration"`    // 历史积分数量
}

type UpdatePasswordReq struct {
	Password string `json:"password"` //密码
}

type UpdatePasswordResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type UpdateMemberReq struct {
	Id                    int64  `json:"id"`
	MemberLevelId         int64  `json:"memberLevelId"`
	Username              string `json:"username"`              // 用户名
	Password              string `json:"password"`              // 密码
	Nickname              string `json:"nickname"`              // 昵称
	Phone                 string `json:"phone"`                 // 手机号码
	Status                int64  `json:"status"`                // 帐号启用状态:0->禁用；1->启用
	CreateTime            string `json:"createTime"`            // 注册时间
	Icon                  string `json:"icon"`                  // 头像
	Gender                int64  `json:"gender"`                // 性别：0->未知；1->男；2->女
	Birthday              string `json:"birthday"`              // 生日
	City                  string `json:"city"`                  // 所做城市
	Job                   string `json:"job"`                   // 职业
	PersonalizedSignature string `json:"personalizedSignature"` // 个性签名
	SourceType            int64  `json:"sourceType"`            // 用户来源
	Integration           int64  `json:"integration"`           // 积分
	Growth                int64  `json:"growth"`                // 成长值
	LuckeyCount           int64  `json:"luckeyCount"`           // 剩余抽奖次数
	HistoryIntegration    int64  `json:"historyIntegration"`    // 历史积分数量
}

type UpdateMemberResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CartItemAddReq struct {
	ProductId         int64   `json:"productId"`
	ProductSkuId      int64   `json:"productSkuId"`
	Quantity          int64   `json:"quantity"`          // 购买数量
	Price             float64 `json:"price"`             // 添加到购物车的价格
	ProductPic        string  `json:"productPic"`        // 商品主图
	ProductName       string  `json:"productName"`       // 商品名称
	ProductSubTitle   string  `json:"productSubTitle"`   // 商品副标题（卖点）
	ProductSkuCode    string  `json:"productSkuCode"`    // 商品sku条码
	ProductCategoryId int64   `json:"productCategoryId"` // 商品分类
	ProductBrand      string  `json:"productBrand"`
	ProductSn         string  `json:"productSn"`
	ProductAttr       string  `json:"productAttr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

type CartItemAddResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type CartItemDeleteReq struct {
	Ids []int64 `json:"Ids"`
}

type CartItemDeleteResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type CartItemUpdateQuantityReq struct {
	Id       int64 `json:"id"`
	Quantity int64 `json:"quantity"` // 购买数量
}

type CartItemUpdateAttrReq struct {
	Id                int64   `json:"id"`
	ProductId         int64   `json:"productId"`
	ProductSkuId      int64   `json:"productSkuId"`
	MemberId          int64   `json:"memberId"`
	Quantity          int64   `json:"quantity"`          // 购买数量
	Price             float64 `json:"price"`             // 添加到购物车的价格
	ProductPic        string  `json:"productPic"`        // 商品主图
	ProductName       string  `json:"productName"`       // 商品名称
	ProductSubTitle   string  `json:"productSubTitle"`   // 商品副标题（卖点）
	ProductSkuCode    string  `json:"productSkuCode"`    // 商品sku条码
	MemberNickname    string  `json:"memberNickname"`    // 会员昵称
	DeleteStatus      int64   `json:"deleteStatus"`      // 是否删除
	ProductCategoryId int64   `json:"productCategoryId"` // 商品分类
	ProductBrand      string  `json:"productBrand"`
	ProductSn         string  `json:"productSn"`
	ProductAttr       string  `json:"productAttr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

type CartItemUpdateResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type CartItemListResp struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Data    []CartListData `json:"data"`
}

type CartListData struct {
	Id                int64   `json:"id"`
	ProductId         int64   `json:"productId"`
	ProductSkuId      int64   `json:"productSkuId"`
	MemberId          int64   `json:"memberId"`
	Quantity          int64   `json:"quantity"`          // 购买数量
	Price             float32 `json:"price"`             // 添加到购物车的价格
	ProductPic        string  `json:"productPic"`        // 商品主图
	ProductName       string  `json:"productName"`       // 商品名称
	ProductSubTitle   string  `json:"productSubTitle"`   // 商品副标题（卖点）
	ProductSkuCode    string  `json:"productSkuCode"`    // 商品sku条码
	MemberNickname    string  `json:"memberNickname"`    // 会员昵称
	CreateDate        string  `json:"createDate"`        // 创建时间
	ModifyDate        string  `json:"modifyDate"`        // 修改时间
	DeleteStatus      int64   `json:"deleteStatus"`      // 是否删除
	ProductCategoryId int64   `json:"productCategoryId"` // 商品分类
	ProductBrand      string  `json:"productBrand"`
	ProductSn         string  `json:"productSn"`
	ProductAttr       string  `json:"productAttr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

type CarItemListPromotionReq struct {
	Id int64 `json:"id"`
}

type CarItemtListPromotionResp struct {
	Code    int64                       `json:"code"`
	Message string                      `json:"message"`
	Data    []CarItemtPromotionListData `json:"data"`
}

type CarItemtPromotionListData struct {
	Id                int64   `json:"id"`
	ProductId         int64   `json:"productId"`
	ProductSkuId      int64   `json:"productSkuId"`
	MemberId          int64   `json:"memberId"`
	Quantity          int64   `json:"quantity"`          // 购买数量
	Price             float32 `json:"price"`             // 添加到购物车的价格
	ProductPic        string  `json:"productPic"`        // 商品主图
	ProductName       string  `json:"productName"`       // 商品名称
	ProductSubTitle   string  `json:"productSubTitle"`   // 商品副标题（卖点）
	ProductSkuCode    string  `json:"productSkuCode"`    // 商品sku条码
	MemberNickname    string  `json:"memberNickname"`    // 会员昵称
	CreateDate        string  `json:"createDate"`        // 创建时间
	ModifyDate        string  `json:"modifyDate"`        // 修改时间
	DeleteStatus      int64   `json:"deleteStatus"`      // 是否删除
	ProductCategoryId int64   `json:"productCategoryId"` // 商品分类
	ProductBrand      string  `json:"productBrand"`
	ProductSn         string  `json:"productSn"`
	ProductAttr       string  `json:"productAttr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
	PromotionMessage  string  `json:"promotionMessage"`
	ReduceAmount      string  `json:"reduceAmount"`
	RealStock         int64   `json:"realStock"`
	Integration       int64   `json:"integration"`
	Growth            int64   `json:"growth"`
}

type CartProductReq struct {
	ProductId int64 `json:"productId"`
}

type CartProductResp struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Data    CartProductData `json:"data"`
}

type CartProductData struct {
	ProductAttributeList []CartItemProductAttributeList `json:"productAttributeList"`
	SkuStockList         []CartItemSkuStockList         `json:"skuStockList"`
}

type CartItemProductAttributeList struct {
	Id                         int64  `json:"id"`
	ProductAttributeCategoryId int64  `json:"productAttributeCategoryId"`
	Name                       string `json:"name"`
	SelectType                 int64  `json:"selectType"`
	InputType                  int64  `json:"inputType"`
	InputList                  string `json:"inputList"`
	Sort                       int64  `json:"sort"`
	FilterType                 int64  `json:"filterType"`
	SearchType                 int64  `json:"searchType"`
	RelatedStatus              int64  `json:"relatedStatus"`
	HandAddStatus              int64  `json:"handAddStatus"`
	Type                       int64  `json:"type"`
}

type CartItemSkuStockList struct {
	Id             int64   `json:"id"`
	ProductId      int64   `json:"productId"`
	SkuCode        string  `json:"skuCode"` // sku编码
	Price          float32 `json:"price"`
	Stock          int64   `json:"stock"`          // 库存
	LowStock       int64   `json:"lowStock"`       // 预警库存
	Pic            string  `json:"pic"`            // 展示图片
	Sale           int64   `json:"sale"`           // 销量
	PromotionPrice float32 `json:"promotionPrice"` // 单品促销价格
	LockStock      int64   `json:"lockStock"`      // 锁定库存
	SpData         string  `json:"spData"`         // 商品销售属性，json格式
}

type CartItemClearResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type CategoryReq struct {
	ParentId int64 `path:"parentId"` // 上级分类的编号：0表示一级分类
}

type CategoryResp struct {
	Code    int64          `json:"code"`
	Message string         `json:"message"`
	Data    []CategoryData `json:"data"`
}

type CategoryData struct {
	Id           int64  `json:"id"`
	ParentId     int64  `json:"parentId"` // 上级分类的编号：0表示一级分类
	Name         string `json:"name"`
	Level        int64  `json:"level"` // 分类级别：0->1级；1->2级
	ProductCount int64  `json:"productCount"`
	ProductUnit  string `json:"productUnit"`
	NavStatus    int64  `json:"navStatus"`  // 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   int64  `json:"showStatus"` // 显示状态：0->不显示；1->显示
	Sort         int64  `json:"sort"`
	Icon         string `json:"icon"` // 图标
	Keywords     string `json:"keywords"`
	Description  string `json:"description"` // 描述
}

type HomeResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	AdvertiseList      []AdvertiseList    `json:"advertiseList"`
	BrandList          []BrandList        `json:"brandList"`
	HomeFlashPromotion HomeFlashPromotion `json:"homeFlashPromotion"`
	NewProductList     []ProductList      `json:"newProductList"`
	HotProductList     []ProductList      `json:"hotProductList"`
	SubjectList        []SubjectList      `json:"subjectList"`
}

type SubjectList struct {
	CategoryId      int64  `json:"categoryId"`
	Title           string `json:"title"`
	Pic             string `json:"pic"`          // 专题主图
	ProductCount    int64  `json:"productCount"` // 关联产品数量
	RecommendStatus int64  `json:"recommendStatus"`
	CollectCount    int64  `json:"collectCount"`
	ReadCount       int64  `json:"readCount"`
	CommentCount    int64  `json:"commentCount"`
	AlbumPics       string `json:"albumPics"` // 画册图片用逗号分割
	Description     string `json:"description"`
	ShowStatus      int64  `json:"showStatus"` // 显示状态：0->不显示；1->显示
	Content         string `json:"content"`
	ForwardCount    int64  `json:"forwardCount"` // 转发数
	CategoryName    string `json:"categoryName"` // 专题分类名称
}

type AdvertiseList struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Type       int64  `json:"type"`
	Pic        string `json:"pic"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	Status     int64  `json:"status"`
	ClickCount int64  `json:"clickCount"`
	OrderCount int64  `json:"orderCount"`
	URL        string `json:"url"`
	Sort       int64  `json:"sort"`
}

type BrandList struct {
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	FirstLetter         string `json:"firstLetter"`
	Sort                int64  `json:"sort"`
	FactoryStatus       int64  `json:"factoryStatus"`
	ShowStatus          int64  `json:"showStatus"`
	ProductCount        int64  `json:"productCount"`
	ProductCommentCount int64  `json:"productCommentCount"`
	Logo                string `json:"logo"`
	BigPic              string `json:"bigPic"`
}

type HomeFlashPromotion struct {
	StartTime     string        `json:"startTime"`
	EndTime       string        `json:"endTime"`
	NextStartTime string        `json:"nextStartTime"`
	NextEndTime   string        `json:"nextEndTime"`
	ProductList   []ProductList `json:"productList"`
}

type ProductList struct {
	Id                         int64   `json:"id"`
	BrandId                    int64   `json:"brandId"`
	ProductCategoryId          int64   `json:"productCategoryId"`
	FeightTemplateId           int64   `json:"feightTemplateId"`
	ProductAttributeCategoryId int64   `json:"productAttributeCategoryId"`
	Name                       string  `json:"name"`
	Pic                        string  `json:"pic"`
	ProductSn                  string  `json:"productSn"`
	DeleteStatus               int64   `json:"deleteStatus"`
	PublishStatus              int64   `json:"publishStatus"`
	NewStatus                  int64   `json:"newStatus"`
	RecommandStatus            int64   `json:"recommandStatus"`
	VerifyStatus               int64   `json:"verifyStatus"`
	Sort                       int64   `json:"sort"`
	Sale                       int64   `json:"sale"`
	Price                      float64 `json:"price"`
	PromotionPrice             float64 `json:"promotionPrice,omitempty"`
	GiftGrowth                 int64   `json:"giftGrowth"`
	GiftPoint                  int64   `json:"giftPoint"`
	UsePointLimit              int64   `json:"usePointLimit"`
	SubTitle                   string  `json:"subTitle"`
	OriginalPrice              float64 `json:"originalPrice"`
	Stock                      int64   `json:"stock"`
	LowStock                   int64   `json:"lowStock"`
	Unit                       string  `json:"unit"`
	Weight                     float64 `json:"weight"`
	PreviewStatus              int64   `json:"previewStatus"`
	ServiceIDS                 string  `json:"serviceIds"`
	Keywords                   string  `json:"keywords"`
	Note                       string  `json:"note"`
	AlbumPics                  string  `json:"albumPics"`
	DetailTitle                string  `json:"detailTitle"`
	PromotionStartTime         string  `json:"promotionStartTime,omitempty"`
	PromotionEndTime           string  `json:"promotionEndTime,omitempty"`
	PromotionPerLimit          int64   `json:"promotionPerLimit"`
	PromotionType              int64   `json:"promotionType"`
	BrandName                  string  `json:"brandName"`
	ProductCategoryName        string  `json:"productCategoryName"`
	Description                string  `json:"description"`
}

type RecommendProductListReq struct {
	Current  int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=10"`
}

type RecommendProductListResp struct {
	Code    int64         `json:"code"`
	Message string        `json:"message"`
	Data    []ProductList `json:"data"`
}

type RecommendBrandListReq struct {
	Current  int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=10"`
}

type RecommendBrandListResp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    []BrandList `json:"data"`
}

type RecommendNewProductListReq struct {
	Current  int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=10"`
}

type RecommendNewProductListResp struct {
	Code    int64         `json:"code"`
	Message string        `json:"message"`
	Data    []ProductList `json:"data"`
}

type RecommendHotProductListReq struct {
	Current  int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=10"`
}

type RecommendHotProductListResp struct {
	Code    int64         `json:"code"`
	Message string        `json:"message"`
	Data    []ProductList `json:"data"`
}

type ReturnApplyReq struct {
	OrderID          int64   `json:"orderId"`
	ProductID        int64   `json:"productId"`
	OrderSn          string  `json:"orderSn"`
	MemberUsername   string  `json:"memberUsername"`
	ReturnName       string  `json:"returnName"`
	ReturnPhone      string  `json:"returnPhone"`
	ProductPic       string  `json:"productPic"`
	ProductName      string  `json:"productName"`
	ProductBrand     string  `json:"productBrand"`
	ProductAttr      string  `json:"productAttr"`
	ProductCount     int64   `json:"productCount"`
	ProductPrice     float32 `json:"productPrice"`
	ProductRealPrice float32 `json:"productRealPrice"`
	Reason           string  `json:"reason"`
	Description      string  `json:"description"`
	ProofPics        string  `json:"proofPics"`
}

type ReturnApplyResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type QueryProductReq struct {
	Id int64 `path:"id"`
}

type QueryProductResp struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    ProductData `json:"data"`
}

type ProductData struct {
	Product                   Product                     `json:"product"`
	Brand                     Brand                       `json:"brand"`
	ProductAttributeList      []ProductAttributeList      `json:"productAttributeList"`
	ProductAttributeValueList []ProductAttributeValueList `json:"productAttributeValueList"`
	SkuStockList              []SkuStockList              `json:"skuStockList"`
	ProductLadderList         []ProductLadderList         `json:"productLadderList"`
	ProductFullReductionList  []ProductFullReductionList  `json:"productFullReductionList"`
	CouponList                []CouponList                `json:"couponList"`
}

type Brand struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	FirstLetter         string `json:"firstLetter"`
	Sort                int64  `json:"sort"`
	FactoryStatus       int64  `json:"factoryStatus"`
	ShowStatus          int64  `json:"showStatus"`
	ProductCount        int64  `json:"productCount"`
	ProductCommentCount int64  `json:"productCommentCount"`
	Logo                string `json:"logo"`
	BigPic              string `json:"bigPic"`
	BrandStory          string `json:"brandStory"`
}

type CouponList struct {
	Id           int64  `json:"id"`
	Type         int64  `json:"type"`
	Name         string `json:"name"`
	Platform     int64  `json:"platform"`
	Count        int64  `json:"count"`
	Amount       int64  `json:"amount"`
	PerLimit     int64  `json:"perLimit"`
	MinPoint     int64  `json:"minPoint"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	UseType      int64  `json:"useType"`
	PublishCount int64  `json:"publishCount"`
	UseCount     int64  `json:"useCount"`
	ReceiveCount int64  `json:"receiveCount"`
	EnableTime   string `json:"enableTime"`
}

type Product struct {
	Id                         int64   `json:"id"`
	BrandId                    int64   `json:"brandId"`
	ProductCategoryId          int64   `json:"productCategoryId"`
	FeightTemplateId           int64   `json:"feightTemplateId"`
	ProductAttributeCategoryId int64   `json:"productAttributeCategoryId"`
	Name                       string  `json:"name"`
	Pic                        string  `json:"pic"`
	ProductSn                  string  `json:"productSn"`
	DeleteStatus               int64   `json:"deleteStatus"`
	PublishStatus              int64   `json:"publishStatus"`
	NewStatus                  int64   `json:"newStatus"`
	RecommandStatus            int64   `json:"recommandStatus"`
	VerifyStatus               int64   `json:"verifyStatus"`
	Sort                       int64   `json:"sort"`
	Sale                       int64   `json:"sale"`
	Price                      float64 `json:"price"`
	PromotionPrice             float64 `json:"promotionPrice"`
	GiftGrowth                 int64   `json:"giftGrowth"`
	GiftPoint                  int64   `json:"giftPoint"`
	UsePointLimit              int64   `json:"usePointLimit"`
	SubTitle                   string  `json:"subTitle"`
	OriginalPrice              float64 `json:"originalPrice"`
	Stock                      int64   `json:"stock"`
	LowStock                   int64   `json:"lowStock"`
	Unit                       string  `json:"unit"`
	Weight                     float64 `json:"weight"`
	PreviewStatus              int64   `json:"previewStatus"`
	ServiceIds                 string  `json:"serviceIds"`
	Keywords                   string  `json:"keywords"`
	Note                       string  `json:"note"`
	AlbumPics                  string  `json:"albumPics"`
	DetailTitle                string  `json:"detailTitle"`
	PromotionStartTime         string  `json:"promotionStartTime"`
	PromotionEndTime           string  `json:"promotionEndTime"`
	PromotionPerLimit          int64   `json:"promotionPerLimit"`
	PromotionType              int64   `json:"promotionType"`
	BrandName                  string  `json:"brandName"`
	ProductCategoryName        string  `json:"productCategoryName"`
	Description                string  `json:"description"`
	DetailDesc                 string  `json:"detailDesc"`
	DetailHtml                 string  `json:"detailHtml"`
	DetailMobileHtml           string  `json:"detailMobileHtml"`
}

type ProductAttributeList struct {
	Id                         int64  `json:"id"`
	ProductAttributeCategoryId int64  `json:"productAttributeCategoryId"`
	Name                       string `json:"name"`
	SelectType                 int64  `json:"selectType"`
	InputType                  int64  `json:"inputType"`
	InputList                  string `json:"inputList"`
	Sort                       int64  `json:"sort"`
	FilterType                 int64  `json:"filterType"`
	SearchType                 int64  `json:"searchType"`
	RelatedStatus              int64  `json:"relatedStatus"`
	HandAddStatus              int64  `json:"handAddStatus"`
	Type                       int64  `json:"type"`
}

type ProductAttributeValueList struct {
	Id                 int64  `json:"id"`
	ProductId          int64  `json:"productId"`
	ProductAttributeId int64  `json:"productAttributeId"`
	Value              string `json:"value"`
}

type ProductLadderList struct {
	Id        int64   `json:"id"`
	ProductId int64   `json:"productId"`
	Count     int64   `json:"count"`
	Discount  int64   `json:"discount"`
	Price     float32 `json:"price"`
}

type SkuStockList struct {
	Id             int64   `json:"id"`
	ProductId      int64   `json:"productId"`
	SkuCode        string  `json:"skuCode"` // sku编码
	Price          float32 `json:"price"`
	Stock          int64   `json:"stock"`          // 库存
	LowStock       int64   `json:"lowStock"`       // 预警库存
	Pic            string  `json:"pic"`            // 展示图片
	Sale           int64   `json:"sale"`           // 销量
	PromotionPrice float32 `json:"promotionPrice"` // 单品促销价格
	LockStock      int64   `json:"lockStock"`      // 锁定库存
	SpData         string  `json:"spData"`         // 商品销售属性，json格式
}

type ProductFullReductionList struct {
	Id          int64   `json:"id"`
	ProductId   int64   `json:"productId"`
	FullPrice   float32 `json:"fullPrice"`
	ReducePrice float32 `json:"reducePrice"`
}

type QueryProductListReq struct {
	Current           int64 `json:"current,default=1"`
	PageSize          int64 `json:"pageSize,default=10"`
	BrandId           int64 `json:"brandId,default=0"`
	ProductCategoryId int64 `json:"productCategoryId,default=0"`
}

type QueryProductListResp struct {
	Code    int64                  `json:"code"`
	Message string                 `json:"message"`
	Data    []QueryProductListData `json:"data"`
}

type QueryProductListData struct {
	Id                         int64   `json:"id"`
	BrandId                    int64   `json:"brandId"`
	ProductCategoryId          int64   `json:"productCategoryId"`
	FeightTemplateId           int64   `json:"feightTemplateId"`
	ProductAttributeCategoryId int64   `json:"productAttributeCategoryId"`
	Name                       string  `json:"name"`
	Pic                        string  `json:"pic"`
	ProductSn                  string  `json:"productSn"`
	DeleteStatus               int64   `json:"deleteStatus"`
	PublishStatus              int64   `json:"publishStatus"`
	NewStatus                  int64   `json:"newStatus"`
	RecommandStatus            int64   `json:"recommandStatus"`
	VerifyStatus               int64   `json:"verifyStatus"`
	Sort                       int64   `json:"sort"`
	Sale                       int64   `json:"sale"`
	Price                      float64 `json:"price"`
	PromotionPrice             float64 `json:"promotionPrice,omitempty"`
	GiftGrowth                 int64   `json:"giftGrowth"`
	GiftPoint                  int64   `json:"giftPoint"`
	UsePointLimit              int64   `json:"usePointLimit"`
	SubTitle                   string  `json:"subTitle"`
	OriginalPrice              float64 `json:"originalPrice"`
	Stock                      int64   `json:"stock"`
	LowStock                   int64   `json:"lowStock"`
	Unit                       string  `json:"unit"`
	Weight                     float64 `json:"weight"`
	PreviewStatus              int64   `json:"previewStatus"`
	ServiceIDS                 string  `json:"serviceIds"`
	Keywords                   string  `json:"keywords"`
	Note                       string  `json:"note"`
	AlbumPics                  string  `json:"albumPics"`
	DetailTitle                string  `json:"detailTitle"`
	PromotionStartTime         string  `json:"promotionStartTime,omitempty"`
	PromotionEndTime           string  `json:"promotionEndTime,omitempty"`
	PromotionPerLimit          int64   `json:"promotionPerLimit"`
	PromotionType              int64   `json:"promotionType"`
	BrandName                  string  `json:"brandName"`
	ProductCategoryName        string  `json:"productCategoryName"`
	Description                string  `json:"description"`
}

type AddProductCollectionReq struct {
	ProductId       int64   `json:"productId"`       // 商品id
	ProductName     string  `json:"productName"`     // 商品名称
	ProductPic      string  `json:"productPic"`      // 商品图片
	ProductSubTitle string  `json:"productSubTitle"` // 商品标题
	ProductPrice    float64 `json:"productPrice"`    // 商品价格
}

type AddProductCollectionResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ProductCollectionDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type ProductCollectionDeleteResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ProductCollectionListReq struct {
	Current   int64 `json:"current,default=1"`
	PageSize  int64 `json:"pageSize,default=20"`
	ProductId int64 `json:"productId,default=0"` // 商品id
}

type ProductCollectionListResp struct {
	Code    int64                     `json:"code"`
	Message string                    `json:"message"`
	Data    ProductCollectionListData `json:"data"`
}

type ProductCollectionListData struct {
	Total int64                   `json:"total"` //总数
	Pages int64                   `json:"pages"` //总页数
	Limit int64                   `json:"limit"` //分页数量
	Page  int64                   `json:"page"`  //当前页
	List  []ProductCollectionList `json:"list"`  //地址列表
}

type ProductCollectionList struct {
	Id              int64   `json:"id"`              // 编号
	MemberId        int64   `json:"memberId"`        // 会员id
	MemberNickName  string  `json:"memberNickName"`  // 会员姓名
	MemberIcon      string  `json:"memberIcon"`      // 会员头像
	ProductId       int64   `json:"productId"`       // 商品id
	ProductName     string  `json:"productName"`     // 商品名称
	ProductPic      string  `json:"productPic"`      // 商品图片
	ProductSubTitle string  `json:"productSubTitle"` // 商品标题
	ProductPrice    float64 `json:"productPrice"`    // 商品价格
	CreateTime      string  `json:"createTime"`      // 浏览时间
}

type ProductCollectionClearResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type AddReadHistoryReq struct {
	ProductId       int64   `json:"productId"`       // 商品id
	ProductName     string  `json:"productName"`     // 商品名称
	ProductPic      string  `json:"productPic"`      // 商品图片
	ProductSubTitle string  `json:"productSubTitle"` // 商品标题
	ProductPrice    float64 `json:"productPrice"`    // 商品价格
}

type AddReadHistoryResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ReadHistoryDeleteReq struct {
	Ids []int64 `json:"ids"`
}

type ReadHistoryDeleteResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ReadHistoryListReq struct {
	Current  int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}

type ReadHistoryListResp struct {
	Code    int64               `json:"code"`
	Message string              `json:"message"`
	Data    ReadHistoryListData `json:"data"`
}

type ReadHistoryListData struct {
	Total int64             `json:"total"` //总数
	Pages int64             `json:"pages"` //总页数
	Limit int64             `json:"limit"` //分页数量
	Page  int64             `json:"page"`  //当前页
	List  []ReadHistoryList `json:"list"`  //地址列表
}

type ReadHistoryList struct {
	Id              int64   `json:"id"`              // 编号
	MemberId        int64   `json:"memberId"`        // 会员id
	MemberNickName  string  `json:"memberNickName"`  // 会员姓名
	MemberIcon      string  `json:"memberIcon"`      // 会员头像
	ProductId       int64   `json:"productId"`       // 商品id
	ProductName     string  `json:"productName"`     // 商品名称
	ProductPic      string  `json:"productPic"`      // 商品图片
	ProductSubTitle string  `json:"productSubTitle"` // 商品标题
	ProductPrice    float64 `json:"productPrice"`    // 商品价格
	CreateTime      string  `json:"createTime"`      // 浏览时间
}

type ReadHistoryClearResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type AddCouponReq struct {
	CouponId int64 `json:"couponId"`
}

type AddCouponResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type ListCouponHistoryReq struct {
	Current   int64 `json:"current,default=1"`
	PageSize  int64 `json:"pageSize,default=20"`
	UseStatus int64 `json:"useStatus,default=3"` // 使用状态：0->未使用；1->已使用；2->已过期
}

type ListCouponHistoryData struct {
	Id             int64  `json:"id"`
	CouponId       int64  `json:"couponId"`
	MemberId       int64  `json:"memberId"`
	CouponCode     string `json:"couponCode"`
	MemberNickname string `json:"memberNickName"` // 领取人昵称
	GetType        int64  `json:"getType"`        // 获取类型：0->后台赠送；1->主动获取
	CreateTime     string `json:"createTime"`
	UseStatus      int64  `json:"useStatus"` // 使用状态：0->未使用；1->已使用；2->已过期
	UseTime        string `json:"useTime"`   // 使用时间
	OrderId        int64  `json:"orderId"`   // 订单编号
	OrderSn        string `json:"orderSn"`   // 订单号码
}

type ListCouponHistoryResp struct {
	Code    int64                    `json:"code"`
	Message string                   `json:"message"`
	Data    []*ListCouponHistoryData `json:"data"`
}

type ListCouponReq struct {
	ProductId int64 `json:"productId"` // 商品id
}

type ListtCouponData struct {
	Id           int64   `json:"id"`
	Type         int64   `json:"type"` // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
	Name         string  `json:"name"`
	Platform     int64   `json:"platform"` // 使用平台：0->全部；1->移动；2->PC
	Count        int64   `json:"count"`    // 数量
	Amount       float64 `json:"amount"`   // 金额
	PerLimit     int64   `json:"perLimit"` // 每人限领张数
	MinPoint     float64 `json:"minPoint"` // 使用门槛；0表示无门槛
	StartTime    string  `json:"startTime"`
	EndTime      string  `json:"endTime"`
	UseType      int64   `json:"useType"`      // 使用类型：0->全场通用；1->指定分类；2->指定商品
	Note         string  `json:"note"`         // 备注
	PublishCount int64   `json:"publishCount"` // 发行数量
	UseCount     int64   `json:"useCount"`     // 已使用数量
	ReceiveCount int64   `json:"receiveCount"` // 领取数量
	EnableTime   string  `json:"enableTime"`   // 可以领取的日期
	Code         string  `json:"code"`         // 优惠码
	MemberLevel  int64   `json:"memberLevel"`  // 可领取的会员类型：0->无限时
}

type ListCouponResp struct {
	Code    int64              `json:"code"`
	Message string             `json:"message"`
	Data    []*ListtCouponData `json:"data"`
}
