package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func GetTimeTick64() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetTimeTick32() int32 {
	return int32(time.Now().Unix())
}

func GetFormatTime(time time.Time) string {
	return time.Format("20060102150405")
}

// 基础做法 日期20191025时间戳1571987125435+3位随机数
func GenerateCode(prefix string) string {
	date := GetFormatTime(time.Now())
	r, _ := rand.Int(rand.Reader, big.NewInt(1000))
	code := fmt.Sprintf("%s%s%03d", prefix, date, r)
	return code
}
