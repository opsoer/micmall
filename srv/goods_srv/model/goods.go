package model

//Category 商品分类表结构
type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20);not null" json:"name"`
	ParentCategoryID int32       `json:"parent"`
	ParentCategory   *Category   `json:"-"` //自己指向自己要用指针
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
}

//Brands 品牌表
type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"` //名称一样  就是联合索引
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

//TableName 自定义表名
func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

//Banner 轮播图
type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`  //跳转到指定的商品详情页
	Index int32  `gorm:"type:int;default:1;not null"` //轮播图顺序
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null"`
	ShipFree bool `gorm:"default:false;not null"` //免运费
	IsNew    bool `gorm:"default:false;not null"`
	IsHot    bool `gorm:"default:false;not null"`

	Name            string   `gorm:"type:varchar(50);not null"`
	GoodsSn         string   `gorm:"type:varchar(50);not null"`   //订单编号
	ClickNum        int32    `gorm:"type:int;default:0;not null"` //商品点击量
	SoldNum         int32    `gorm:"type:int;default:0;not null"` //商品售卖量
	FavNum          int32    `gorm:"type:int;default:0;not null"` //收藏
	MarketPrice     float32  `gorm:"not null"`                    //市场价格
	ShopPrice       float32  `gorm:"not null"`                    //本地价格
	GoodsBrief      string   `gorm:"type:varchar(100);not null"`  //商品简介
	Images          GormList `gorm:"type:varchar(1000);not null"` //GormList 为自定义切片类型 gorm没有内置
	DescImages      GormList `gorm:"type:varchar(1000);not null"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null"` //商品封面图
}
