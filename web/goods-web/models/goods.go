package models

type GoodsInfo struct {
	Id              int32    `json:"id"`
	Name            string   `json:"name"`
	GoodsBrief      string   `json:"goods_brief"` //商品简介
	ShipFree        bool     `json:"ship_free"`
	Images          []string `json:"images"`
	DescImages      []string `json:"desc_images"` //商品详情的图片
	GoodsFrontImage string   `json:"goods_front_image"`
	ShopPrice       float32  `json:"shop_price"`
	Category        `json:"category"`
	Brands          `json:"brands"`
	OnSale          bool `json:"on_sale"`
	IsNew           bool `json:"is_new"`
	IsHot           bool `json:"is_hot"`
}

type Category struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type Brands struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}
