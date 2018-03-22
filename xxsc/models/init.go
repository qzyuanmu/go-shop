package models

import (
	"time"
)

import (
	"fmt"
	"qzyuanmu/jtool/shop"
)

func FormatPrice(price int) string {

	n := float32(price)
	n = n / 100
	return fmt.Sprintf("%.2f", n)

}

func FormatRed(p shop.Product) string {

	n := float32(p.Price)
	n = n / 100
	s := int(n * float32(p.RedLevel) / 100 / 2)
	return fmt.Sprintf("%d", s)

}
func FormatZRed(p shop.Product) string {

	n := float32(p.Price + p.Zhim*100)
	n = n / 100
	s := int(n * float32(p.RedLevel) / 100 / 2)
	return fmt.Sprintf("%d", s)

}

func FormatRedBox(p, RedLevel int) int {

	n := float32(p)
	n = n / 100
	s := int(n * float32(RedLevel) / 100 / 2)
	return s

}
func FormatDateTime(d time.Time, i int) string {
	d = d.Add(-8 * time.Hour)
	if i == 1 {
		return d.Format("2006-01-02 15:04:05")
	} else {
		return d.Format("2006-01-02")
	}

}

func FormatRow(i int) int {
	return i + 1
}
