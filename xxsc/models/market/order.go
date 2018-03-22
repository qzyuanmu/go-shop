package market

import (
	"errors"
	"qzyuanmu/jtool"
	"qzyuanmu/jtool/shop"
	"qzyuanmu/jtool/user"
)

type OrderModel struct {
	Order    shop.Order
	Shop     shop.Shop
	OrderDet []shop.OrderDet
}

func GetOrderListByUser(uId int) []OrderModel {
	ms := make([]OrderModel, 0)
	os := shop.GetOrder(uId)
	for _, o := range os {
		m := new(OrderModel)
		m.Order = o
		m.Shop = *(shop.GetShopByShopId(o.ShopId))
		m.OrderDet = make([]shop.OrderDet, 0)
		m.OrderDet = shop.GetOrderDet(o.Id)
		ms = append(ms, *m)
	}
	return ms
}

func AddOrder(o shop.Order, uId int, z int) bool {
	cs := GetCart(uId)
	os := GetNewOrder(o, uId, len(cs.Carts))
	err := errors.New("")
	u := user.GetUserById(uId)
	session := jtool.Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	for i, c := range cs.Carts {
		no := new(shop.Order)
		no.Id = os[i].Id
		no.IsValid = 2
		if z != 2 {
			no.SumMoney = c.SumMoney + c.SumZhim*100
			no.SumPrice = c.SumMoney
			no.SumZhim = 0
			no.SumRed = c.SumRed
			no.ShopId = c.Shop.Id
		} else {
			no.SumMoney = c.SumMoney
			no.SumPrice = c.SumMoney
			no.SumZhim = c.SumZhim
			no.SumRed = c.SumZRed
			no.ShopId = c.Shop.Id
			u.JiFen = u.JiFen - c.SumZhim
			_, err = session.Id(u.Id).Update(u)
			if err != nil {
				session.Rollback()
				jtool.Log.Error("更新会员积分失败3", err)
				return false
			}

		}
		_, err = session.Id(no.Id).Update(no)
		if err != nil {
			session.Rollback()
			jtool.Log.Error("生成订单失败3", err)
			return false
		}

		for _, p := range c.Products {
			det := new(shop.OrderDet)
			det.OrderId = no.Id
			det.OriPrice = p.OriPrice
			det.Price = p.Price
			det.ProductId = p.Id
			det.ProductImg = p.MinImg
			det.ProductTypeId = p.CartTypeId
			det.ProductTypeName = p.CartTypeName
			det.Quantity = p.CartQuantity
			det.RedLevel = p.RedLevel
			det.ShopId = no.ShopId
			det.UserId = uId
			det.Zhim = p.Zhim
			det.ProductName = p.Name
			_, err = session.Insert(det)
			if err != nil {
				session.Rollback()
				jtool.Log.Error("生成订单失败3", err)
				return false
			}

		}

	}

	cart := new(shop.Cart)
	_, err = session.Where("UserId=?", uId).Delete(cart)
	if err != nil {
		session.Rollback()
		jtool.Log.Error("删除购物车失败", err)
		return false
	}

	err = session.Commit()
	if err != nil {
		jtool.Log.Error("生成订单失败4", err)
		return false
	} else {
		return true
	}

}

//根据购物中的商品，根据店铺提前生成订单
func GetNewOrder(o shop.Order, uId, row int) []shop.Order {
	os := make([]shop.Order, 0)
	err := errors.New("")
	session := jtool.Engine.NewSession()
	defer session.Close()
	err = session.Begin()

	for i := 1; i <= row; i++ {
		no := new(shop.Order)
		no.Address = o.Address
		no.BestTime = o.BestTime
		no.CityId = o.CityId
		no.DistrictId = o.DistrictId
		no.Email = o.Email
		no.Flg = 1
		no.FlgValue = "等待付款"
		no.PayFlg = 1

		no.Ip = o.Ip
		no.IsValid = 1
		no.Memo = o.Memo
		no.Mobile = o.Mobile
		no.ProvinceId = o.ProvinceId
		no.ZipCode = o.ZipCode
		no.UserId = uId
		no.RealName = o.RealName
		_, err = session.Insert(no)
		if err != nil {
			session.Rollback()
			jtool.Log.Error("生成订单失败1", err)
		}
		os = append(os, *no)
	}
	err = session.Commit()
	if err != nil {
		jtool.Log.Error("生成订单失败2", err)
	}

	return os

}
