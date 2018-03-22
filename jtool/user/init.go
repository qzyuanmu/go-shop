package user

import (
	"math/rand"
	"qzyuanmu/jtool"
	"strconv"
)

func init() {
	jtool.Engine.Sync2(new(User), new(UserRole), new(UserAddress),
		new(UserMoneyLog), new(TiXian))

}

func Rand(n int) string {
	str := ""
	for i := 1; i <= n; i++ {
		str += strconv.Itoa(rand.Intn(10))
	}
	return str
}
