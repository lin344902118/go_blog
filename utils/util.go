package utils

import (
	"math/rand"
	"time"
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

func FormatTime(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}