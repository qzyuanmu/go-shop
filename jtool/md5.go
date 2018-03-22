package jtool

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	bt := []byte(str + Key)
	b := md5.Sum(bt)
	return hex.EncodeToString(b[:])
}
