package utils

import "time"

const (
	DateTimeKST = "2006-01-02 15:04:05"
	DateKST = "2006-01-02"
)

var kst *time.Location

func init() {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic("timeutil: Asia/Seoul: " + err.Error())
	}
	kst = loc
}

func KST() *time.Location {
	return kst
}

func FormatDateTimeKST(t time.Time) string {
	return t.In(kst).Format(DateTimeKST)
}

func FormatDateKST(t time.Time) string {
	return t.In(kst).Format(DateKST)
}

func ParseDateTimeKST(s string) (time.Time, error) {
	return time.ParseInLocation(DateTimeKST, s, kst)
}

func ParseDateKST(s string) (time.Time, error) {
	return time.ParseInLocation(DateKST, s, kst)
}