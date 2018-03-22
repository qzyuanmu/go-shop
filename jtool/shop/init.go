package shop

import (
	"qzyuanmu/jtool"
)

func init() {
	jtool.Engine.Sync2(new(Shop), new(Category), new(ShopCategory),
		new(Product), new(ProductImg), new(CategoryType), new(ProductType),
		new(Cart), new(ShopType), new(Order), new(OrderDet))
}
