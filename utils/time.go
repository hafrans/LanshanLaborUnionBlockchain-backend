package utils

import (
	"database/sql/driver"
	"fmt"
	"log"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

const dateFormat = "2006-01"

const dateDayFormat = "2006-01-02"

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	tx, err := time.Parse("\""+timeFormat+"\"", string(b))
	if err != nil {
		log.Println(err)
		return err
	}
	*t = Time(tx)
	return nil
}

func NowTime() *Time {
	time := Time(time.Now())
	return &time
}

func GetTime(t time.Time) *Time {
	time := Time(t)
	return &time
}

func (t *Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) FormattedString() string {
	return time.Time(t).Format(timeFormat)
}

func CurrentTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (t Time) Value() (driver.Value, error) {
	var xt time.Time
	var tt time.Time = time.Time(t)
	if tt.UnixNano() == xt.UnixNano() {
		return nil, nil
	} else {
		return tt, nil
	}
}

// Scan valueof time.Time
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type DateMonth time.Time

func (t *DateMonth) UnmarshalJSON(b []byte) error {
	tx, err := time.Parse("\""+dateFormat+"\"", string(b))
	if err != nil {
		log.Println(err)
		return err
	}
	*t = DateMonth(tx)
	return nil
}

func NowDate() *DateMonth {
	time := DateMonth(time.Now())
	return &time
}

func GetDate(t time.Time) *DateMonth {
	time := DateMonth(t)
	return &time
}

func (t *DateMonth) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormat)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, dateFormat)
	b = append(b, '"')
	return b, nil
}

func (t DateMonth) FormattedString() string {
	return time.Time(t).Format(dateFormat)
}

func CurrentDateString() string {
	return time.Now().Format(dateFormat)
}

func (t DateMonth) Value() (driver.Value, error) {
	var xt time.Time
	var tt time.Time = time.Time(t)
	if tt.UnixNano() == xt.UnixNano() {
		return nil, nil
	} else {
		return tt, nil
	}
}

// Scan valueof time.Time
func (t *DateMonth) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DateMonth(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type Date time.Time

func (t *Date) UnmarshalJSON(b []byte) error {
	tx, err := time.Parse("\""+dateDayFormat+"\"", string(b))
	if err != nil {
		log.Println(err)
		return err
	}

	*t = Date(tx)
	return nil
}

func NowDateDay() *Date {
	time := Date(time.Now())
	return &time
}

func GetDateDay(t time.Time) *Date {
	time := Date(t)
	return &time
}

func (t *Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateDayFormat)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, dateDayFormat)
	b = append(b, '"')
	return b, nil
}

func (t Date) FormattedString() string {
	return time.Time(t).Format(dateDayFormat)
}

func CurrentDateDayString() string {
	return time.Now().Format(dateDayFormat)
}

func (t Date) Value() (driver.Value, error) {
	var xt time.Time
	var tt time.Time = time.Time(t)
	if tt.UnixNano() == xt.UnixNano() {
		return nil, nil
	} else {
		return tt, nil
	}
}

// Scan valueof time.Time
func (t *Date) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Date(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
