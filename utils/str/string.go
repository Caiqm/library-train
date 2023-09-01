package str

import (
	"fmt"
	date2 "github.com/Caiqm/library-train-go/date"
	"math/rand"
	"strings"
	"time"
)

type Str struct {
	Types int
}

const (
	specialCharacter  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ&*#@{}[]?!"
	numberPlusStrings = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	plusStrings       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers           = "0123456789"
)

// 判断字符串是否存在某个指定字符串
func (s Str) Stripos(haystack string, needle string, offset ...int) int {
	off := 0
	if len(offset) > 0 {
		off = offset[0]
	}
	if off > len(haystack) || off < 0 {
		return -1
	}
	haystack = strings.ToLower(haystack[off:])
	needle = strings.ToLower(needle)
	index := strings.Index(haystack, needle)
	if index != -1 {
		return off + index
	}
	return index
}

// 生成指定长度盐值
func (s Str) Salt(length int) string {
	// 设置随机种子
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	// 截取的字符串
	letterBytes := specialCharacter
	b := make([]byte, length)
	// 生成随机串
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

// 生成随机字符串，可指定类型，10纯数字，20纯字母，默认是数字与字母混合
func (s Str) RandStr(length int) string {
	// 设置随机种子
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	// 截取的字符串
	letterBytes := numberPlusStrings
	if s.Types == 10 {
		letterBytes = numbers
	} else if s.Types == 20 {
		letterBytes = plusStrings
	}
	b := make([]byte, length)
	// 生成随机串
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

// 简易生成订单号 日期20191025时间戳1571987125435+3位随机数
func (s Str) GenerateCode() string {
	// 获取当前日期
	date := date2.Time{TimeLayout: "20060102150405"}.GetDate()
	r := rand.Intn(1000)
	timeTick64 := time.Now().UnixNano() / 1e6
	code := fmt.Sprintf("%s%d%03d", date, timeTick64, r)
	return code
}
