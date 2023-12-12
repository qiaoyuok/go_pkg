package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const DateTimeFormat = "2006-01-02 15:04:05"
const TimeFormat = "15:04:05"
const DateFormat = "2006-01-02"

type LocalTime time.Time

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "null" || string(data) == `""` {
		return
	}

	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+DateTimeFormat+`"`, string(data))
	*t = LocalTime(now)
	return
}

func (t *LocalTime) IsZero() bool {
	return time.Time(*t).IsZero()
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	// 如果时间值是空或者0值 返回为null 如果写空字符串会报错
	if &t == nil || t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" || t.String() == "0000-00-00 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(DateTimeFormat)), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = LocalTime(tTime)
	return nil
}

func (t LocalTime) String() string {
	return time.Time(t).Format(DateTimeFormat)
}

func (t LocalTime) GetTimeString() string {
	return time.Time(t).Format(TimeFormat)
}

func (t LocalTime) GetDateString() string {
	return time.Time(t).Format(DateFormat)
}

func (t LocalTime) GetDateTimeString() string {
	return time.Time(t).Format(DateTimeFormat)
}

func (t LocalTime) GetTime() time.Time {
	loc, _ := time.LoadLocation("Local")
	dt, _ := time.ParseInLocation(DateTimeFormat, t.String(), loc)
	return dt
}

func (t LocalTime) Before(a time.Time) bool {
	return t.GetTime().Before(a)
}

func (t LocalTime) After(a time.Time) bool {
	return t.GetTime().After(a)
}

func (t LocalTime) Equal(a time.Time) bool {
	return t.GetTime().Equal(a)
}

func (t LocalTime) DayStart() time.Time {
	y, m, d := t.GetTime().Date()

	return time.Date(y, m, d, 0, 0, 0, 0, t.GetTime().Location())
}

func (t LocalTime) Tomorrow() time.Time {
	y, m, d := t.GetTime().Date()

	newDate := time.Date(y, m, d+1, 0, 0, 0, 0, t.GetTime().Location())

	return newDate.AddDate(0, 0, 1)
}

// GetNowDay 获取当前日期
func (t LocalTime) GetNowDay() string {
	return t.GetTime().Format(DateTimeFormat)
}
