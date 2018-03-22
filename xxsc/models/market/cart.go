package market

import (
	"strconv"
)

import (
	"qzyuanmu/jtool/shop"
	"qzyuanmu/xxsc/models"
)

type ShopCart struct {
	Shop     shop.Shop
	Products []shop.Product
	SumMoney int
	SumRed   int
	SumZhim  int
	SumZRed  int
}

type CartModel struct {
	SumMoney int
	SumRed   int
	SumZhim  int

	Carts []ShopCart
}

func GetCart(uId int) *CartModel {
	cm := new(CartModel)
	cm.SumMoney = 0
	cm.SumRed = 0
	cm.Carts = make([]ShopCart, 0)
	cs := shop.GetCartByUser(uId)
	shops := shop.GetShopByUserIdFormCart(uId)
	ps := shop.GetProductByUserIdFormCart(uId)
	np := FormatCart(cs, ps)
	for _, s := range shops {
		scart := new(ShopCart)
		scart.Shop = s
		scart.Products = FindPorductByshopId(s.Id, np)
		scart.SumMoney = SumMoney(scart.Products)
		scart.SumRed = SumRed(scart.Products)
		scart.SumZhim = SumZhim(scart.Products)
		scart.SumZRed = SumZRed(scart.Products)
		cm.Carts = append(cm.Carts, *scart)
	}
	cm.SumMoney = SumMoney(np)
	cm.SumRed = SumRed(np)
	cm.SumZhim = SumZhim(np)
	return cm
}

func SumMoney(ps []shop.Product) int {
	m := 0
	for _, p := range ps {
		m += p.CartQuantity * p.Price
	}
	return m
}
func SumZhim(ps []shop.Product) int {
	m := 0
	for _, p := range ps {
		m += p.CartQuantity * p.Zhim
	}
	return m
}

func SumRed(ps []shop.Product) int {
	m := 0
	for _, p := range ps {
		m += p.CartQuantity * p.CartRed
	}
	return m
}

func SumZRed(ps []shop.Product) int {
	m := 0
	for _, p := range ps {
		m += p.CartQuantity * p.CartZRed
	}

	return m
}

func FindPorductByshopId(sId int, ps []shop.Product) []shop.Product {
	np := make([]shop.Product, 0)
	for _, p := range ps {
		if p.ShopId == sId {
			np = append(np, p)
		}
	}
	return np
}

//格式化购物车内的商品
func FormatCart(cs []shop.Cart, ps []shop.Product) []shop.Product {
	nps := make([]shop.Product, 0)
	for _, c := range cs {
		for _, p := range ps {
			if p.Id == c.ProductId {
				np := &(shop.Product{})
				np = &p
				np.CartQuantity = c.Quantity
				np.CartRed, _ = strconv.Atoi(models.FormatRed(p))
				np.CartZRed, _ = strconv.Atoi(models.FormatZRed(p))
				np.CartTypeId = c.ProductTypeId
				np.CartId = c.Id
				np.CartTypeName = c.ProductTypeName
				np.CartMoney = c.Quantity * p.Price
				if len(c.ProductImg) > 0 {
					np.MinImg = c.ProductImg
				}
				nps = append(nps, *np)
			}
		}

	}
	return nps
}
