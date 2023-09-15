package util

import (
	"strconv"
	"time"
)

func UnixToTime(timeStr string) time.Time {
	i, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		panic(err)
	}

	return ToBrusselsTime(time.Unix(i, 0))
}

func ToBrusselsTime(t time.Time) time.Time {
	loc, err := time.LoadLocation("Europe/Brussels")

	if err != nil {
		panic(err)
	}

	return t.In(loc)
}
