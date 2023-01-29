package date

import "time"

const defaultTimeLayout = "2006-01-02 15:04:05"

// 返回当前日期格式
func GetDate() string {
	timeLayout := defaultTimeLayout
	return time.Now().Format(timeLayout)
}

// 返回当前日期时间戳
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// 时间戳转日期
func Timestamp2Date(timestamp int64, timeLayout string) string {
	if timeLayout == "" {
		timeLayout = defaultTimeLayout // 转化所需模板
	}
	return time.Unix(timestamp, 0).Format(timeLayout)
}

// 日期格式化
func Date2TimeStamp(datetime string) int64 {
	// 日期转化为时间戳
	timeLayout := "2006-01-02" // 转化所需模板
	t, _ := time.Parse(timeLayout, datetime)
	return t.Unix()
}
