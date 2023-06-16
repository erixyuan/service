package utils

import (
	"time"
)

// 根据传入的月字符串，获取当月第一天
func GetTimeByMonthString(dt string) time.Time {
	timeLayout := "2006-01"                  //转化所需模板
	theTime, _ := time.Parse(timeLayout, dt) //使用模板在对应时区转化为time.time类型
	return theTime
}

// GetTimeBySimpleDateString @Description: 转换日期为时间类型
func GetTimeBySimpleDateString(date string) time.Time {
	theTime, _ := time.Parse("2006-01-02", date)
	return theTime
}

// GetTimeByDateTimeString @Description: 转换日期为时间类型
func GetTimeByDateTimeString(date string) time.Time {
	theTime, _ := time.Parse("2006-01-02 15:04:05", date)
	return theTime
}

func GetMonthDateList(d time.Time) []string {
	var dtList []string

	var start time.Time
	var end time.Time
	aMonth := d.Month()
	bMonth := time.Now().AddDate(0, 0, -time.Now().Day()+1).Month()
	// 判断传入的日期 是否大于 当前日期
	// 2022-06-01
	if aMonth < bMonth || d.Year() < time.Now().Year() {
		start = d.AddDate(0, 0, -d.Day()+1).AddDate(0, 1, -1)
		end = GetLastDateTimeOfMonth(start)
	} else {
		start = GetFirstDateTimeOfMonth(time.Now())
		end = time.Now()
	}
	endDate := end.Format("2006-01-02")
	for i := 1; i < 32; i++ {
		dt := start.AddDate(0, 0, -start.Day()+i).Format("2006-01-02")
		dtList = append(dtList, dt)
		if dt == endDate {
			break
		}
	}
	return dtList
}

func GetFirstDateOfMonth(d time.Time) string {
	d = d.AddDate(0, 0, -d.Day()+1)
	return d.Format("2006-01-02")
}

func GetLastDateOfMonth(d time.Time) string {
	return d.AddDate(0, 0, -d.Day()+1).AddDate(0, 1, -1).Format("2006-01-02")
}
func GetFirstDateTimeOfMonth(d time.Time) time.Time {
	return d.AddDate(0, 0, -d.Day()+1)
}

func GetLastDateTimeOfMonth(d time.Time) time.Time {
	return d.AddDate(0, 0, -d.Day()+1).AddDate(0, 1, -1)
}

func GetTodayUnixTime() time.Time {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return addTime
}

func GetTodayUnixEndTime() time.Time {
	t := time.Now()
	addTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return addTime
}

func GetTimeAfter(t time.Time, hour int, minute int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+hour, t.Minute()+minute, 0, 0, t.Location())
}

func GetTimeAfterDays(t time.Time, days int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+days, t.Hour(), t.Minute(), 0, 0, t.Location())
}
func GetTimeBeforeDays(t time.Time, days int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()-days, t.Hour(), t.Minute(), 0, 0, t.Location())
}
func GetFormatTimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetFormatDateStr(t time.Time) string {
	return t.Format("2006-01-02")
}

func GetNowTimeFormatStr() string {
	return GetFormatTimeStr(time.Now())
}

func GetFormatTimeStrAfter(hour int, minute int) string {
	t := time.Now()
	afterTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+hour, t.Minute()+minute, 0, 0, t.Location())
	return GetFormatTimeStr(afterTime)
}

func GetMonthMinTimeByYearAndMonth(year int, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
}

func GetMonthMaxTimeByYearAndMonth(year int, month int) time.Time {
	_, m, day := GetMonthMinTimeByYearAndMonth(year, month).AddDate(0, 1, -1).Date()
	return time.Date(year, m, day, 23, 59, 59, 0, time.Local)
}
