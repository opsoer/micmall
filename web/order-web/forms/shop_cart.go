package forms

type ShopCartItemForm struct {
	//把商品加入购物车
	GoodsId int32 `json:"goods" binding:"required"`
	Nums int32 `json:"nums" binding:"required,min=1"`
}

type ShopCartItemUpdateForm struct {
	Nums int32 `json:"nums" binding:"required,min=1"`
	Checked *bool `json:"checked"`
}