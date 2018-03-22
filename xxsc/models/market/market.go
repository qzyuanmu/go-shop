package market

import (
	"qzyuanmu/jtool/shop"
)

type HomeCategory struct {
	FirstCategory shop.Category
	TwoCategory   []shop.Category
	Products      []shop.Product
}

func NewHomeCategory() HomeCategory {
	home := new(HomeCategory)
	home.FirstCategory = *(new(shop.Category))
	home.TwoCategory = make([]shop.Category, 0)
	home.Products = make([]shop.Product, 0)
	return *home
}

func Home() []HomeCategory {
	homes := make([]HomeCategory, 0)
	cs := shop.FindCategoryRoot()
	for _, c := range cs {
		h := NewHomeCategory()
		h.FirstCategory = c
		h.TwoCategory = shop.FindCategoryByParentId(c.Id)
		h.Products = shop.GetProductByCategoryId(c.Id)

		homes = append(homes, h)
	}
	return homes
}
