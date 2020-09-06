package utils

import (
	"time"
)

const (
	LocalDateTimeFormat string = "2006-01-02 15:04:05"
)

type LocalTime time.Time

func (l LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(LocalDateTimeFormat)+2)

	b = append(b, '"')
	if time.Time(l).Unix() < 0 {
		b = time.Time(l).AppendFormat(b, ``)
	} else {
		b = time.Time(l).AppendFormat(b, LocalDateTimeFormat)
	}
	b = append(b, '"')
	return b, nil
}

func (l *LocalTime) UnmarshalJSON(b []byte) error {
	stringB := string(b)
	if stringB == `""` {
		now, err := time.ParseInLocation(``, ``, time.Local)
		*l = LocalTime(now)
		return err
	} else {
		now, err := time.ParseInLocation(`"`+LocalDateTimeFormat+`"`, stringB, time.Local)
		*l = LocalTime(now)
		return err
	}

}

func (l LocalTime) String() string {
	return time.Time(l).Format(LocalDateTimeFormat)
}
func (l LocalTime) Unix() int64 {
	return time.Time(l).Unix()
}
