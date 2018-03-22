package market

import (
	"qzyuanmu/jtool/shop"
)

type HomeProduct struct {
	FirstCategory, TwoCategory, ThreeCategory *shop.Category
	SiblingCategorys                          []shop.Category
	Product                                   *shop.Product
	ProductTypes                              []shop.ProductType
	ProductTypes1                             []shop.ProductType
	Imgs                                      []shop.ProductImg
}

func NewHomeProduct(pId int) *HomeProduct {
	hp := new(HomeProduct)
	hp.SiblingCategorys = make([]shop.Category, 0)
	hp.Imgs = make([]shop.ProductImg, 0)
	hp.ProductTypes = make([]shop.ProductType, 0)
	hp.ProductTypes1 = make([]shop.ProductType, 0)
	c := new(shop.Category)
	hp.Product = new(shop.Product)
	hp.Product = hp.Product.Get(pId)
	hp.ThreeCategory = c.Get(hp.Product.CategoryId)
	hp.TwoCategory = c.Get(hp.ThreeCategory.ParentId)
	hp.FirstCategory = c.Get(hp.TwoCategory.ParentId)
	hp.SiblingCategorys = shop.FindCategoryByParentId(hp.TwoCategory.Id)
	hp.Imgs = shop.GetProductImgByProductId(pId)
	hp.ProductTypes = shop.GetProductTypeByProductId(pId, 1)
	hp.ProductTypes1 = shop.GetProductTypeByProductId(pId, 2)
	return hp
}
