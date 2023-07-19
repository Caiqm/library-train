package date

import "time"

type Time struct {
	TimeLayout string
}

const defaultTimeLayout = "2006-01-02 15:04:05"

// 返回当前日期格式
func (t Time) GetDate() string {
	if len(t.TimeLayout) <= 0 {
		t.TimeLayout = defaultTimeLayout // 转化所需模板
	}
	return time.Now().Format(t.TimeLayout)
}

// 返回当前日期时间戳
func (t Time) GetTimestamp() int64 {
	return time.Now().Unix()
}

// 时间戳转日期
func (t Time) Timestamp2Date(timestamp int64) string {
	if len(t.TimeLayout) <= 0 {
		t.TimeLayout = defaultTimeLayout // 转化所需模板
	}
	return time.Unix(timestamp, 0).Format(t.TimeLayout)
}

// 日期格式化
func (t Time) Date2TimeStamp(datetime string) int64 {
	if len(t.TimeLayout) <= 0 {
		t.TimeLayout = defaultTimeLayout // 转化所需模板
	}
	// 转化为时间类型
	timeTmp, _ := time.Parse(t.TimeLayout, datetime)
	// 日期转化为时间戳
	return timeTmp.Unix()
}

// 返回当天开始时间和当前时间
func (t Time) GetDateTodayStartAndNow() (start, now time.Time) {
	// 1.获取当前时区
	loc, _ := time.LoadLocation("Local")
	// 2.今日日期字符串
	date := time.Now().Format("2006-01-02")
	// 3.拼接成当天0点时间字符串
	startDate := date + " 00:00:00"
	// 得到0点日期 2021-04-24 00:00:00 +0800 CST
	start, _ = time.ParseInLocation("2006-01-02 15:04:05", startDate, loc)
	// 当前时间
	now = time.Now()
	return
}

// 返回最近7天日期
func (t Time) GetDateLatestSevenDay() (weekDate []string) {
	if len(t.TimeLayout) <= 0 {
		t.TimeLayout = defaultTimeLayout // 转化所需模板
	}
	nowTime := time.Now()
	for i := 6; i > 0; i-- {
		// 获取昨天日期信息
		tmpDate := nowTime.AddDate(0, 0, -i)
		// 转换日期
		d := tmpDate.Format(t.TimeLayout)
		weekDate = append(weekDate, d)
	}
	weekDate = append(weekDate, nowTime.Format(t.TimeLayout))
	return
}

// 返回最近7天时间戳，当日开始时间到隔天开始时间二维数组
func (t Time) GetDateLatestSevenDayUnix() (weekUnix [][]time.Time) {
	todayStart, _ := t.GetDateTodayStartAndNow()
	for i := 6; i > 0; i-- {
		// 获取日期信息
		tmpDate1 := todayStart.AddDate(0, 0, -i)
		tmpDate2 := todayStart.AddDate(0, 0, -(i - 1))
		weekUnix = append(weekUnix, []time.Time{tmpDate1, tmpDate2})
	}
	weekUnix = append(weekUnix, []time.Time{todayStart, todayStart.AddDate(0, 0, +1)})
	return
}
