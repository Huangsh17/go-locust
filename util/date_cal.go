package util

import (
	"time"
)

const TimeLayout = "2006-01-02"
const TenMinuteMilliSec = 10 * 60
const TimeFormat = "2006-01-02 15:04:05"

//获取两个日期相差几天
//t1 - t2
func TimeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}

func TimeStamp2Time(timeStamp int64) time.Time {
	return time.Unix(timeStamp/1000, 0)
}

func PresentTomorrowDate() string {
	// 往后一天
	parseDuration, _ := time.ParseDuration("24h")
	return time.Now().Add(parseDuration).Format(TimeLayout)
}

func PresentYesterday() string {
	// 往前一天
	parseDuration, _ := time.ParseDuration("-24h")
	return time.Now().Add(parseDuration).Format(TimeLayout)
}

func NowTime() time.Time {
	return time.Now()
}
func NowTimeFormat() string {
	return time.Now().Format(TimeFormat)
}

func GetDaySeconds(days int) int {
	return days * 24 * 60 * 60
}
