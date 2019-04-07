package date

import (
	"time"
)

// TodayTimeStamp 获取今天0点的时间戳
func TodayTimeStamp() int64 {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local).Unix()

}

// TodayTimeStamp 获取今天24点的时间戳
func TodayEndTimeStamp() int64 {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 23, 59, 59, 999, time.Local).Unix()

}

// YesterdayTimeStamp  获取昨天0点的时间戳
func YesterdayTimeStamp() int64 {
	return TodayTimeStamp() - 86400
}

// YesterdayTimeStamp  获取昨天24点的时间戳
func YesterdayEndTimeStamp() int64 {
	return TodayEndTimeStamp() - 86400
}

// TodayFormat 获取今天0点的日期
func TodayFormat() string {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local).Format("2006-01-02 15:04:05")
}

// TodayEndFormat 获取今天24点的日期
func TodayEndFormat() string {
	n := time.Now()
	return time.Date(n.Year(), n.Month(), n.Day(), 23, 59, 59, 999, time.Local).Format("2006-01-02 15:04:05")
}

// NowFormat 获取现在的日期
func NowFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TimeStampToDate 时间戳转换为日期格式
func TimeStampToDate(timeStamp int64) string {
	tm := time.Unix(timeStamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func DateToTimeStamp(date string) int64 {
	return str2Time(date).Unix()
}

/**字符串->时间对象*/
func str2Time(formatTimeStr string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc)

	return theTime

}

/**字符串->时间戳*/
func Str2Stamp(formatTimeStr string) int64 {
	timeStruct := str2Time(formatTimeStr)
	millisecond := timeStruct.UnixNano() / 1e6
	return millisecond
}

/*时间戳->字符串*/
func stamp2Str(stamp int64) string {
	timeLayout := "2006-01-02 15:04:05"
	str := time.Unix(stamp/1000, 0).Format(timeLayout)
	return str
}

/*时间戳->时间对象*/
func stamp2Time(stamp int64) time.Time {
	stampStr := stamp2Str(stamp)
	timer := str2Time(stampStr)
	return timer
}
