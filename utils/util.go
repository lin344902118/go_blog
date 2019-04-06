package utils

import (
	"math/rand"
)

const letters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var EnglishToChinese = map[string]string{"blog":"博文", "category": "分类"}

func GetRandomString(length int) string{
	randStr := make([]byte, length)
	for i := range randStr {
		randStr[i] = letters[rand.Intn(length)]
	}
	return string(randStr)
}

func Translates(str string) string{
	for key, value := range EnglishToChinese {
		if key == str {
			return value
		}
	}
	return str
}