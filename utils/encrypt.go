package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const secret  = "beego"

func Md5Encrypted(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum([]byte(secret))
	return hex.EncodeToString(cipherStr)
}
