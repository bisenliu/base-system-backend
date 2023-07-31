package common

import (
	"strconv"
	"time"
)

func Timestamp2Datetime(timestamp int64) string {
	timestampLen := len(strconv.FormatInt(timestamp, 10))
	if timestampLen != 10 {
		timestamp /= 1000
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(time.DateOnly)
}

func TimeStr2TimeStamp(timeStr string) (st int64, err error) {
	dataTime, err := time.Parse(time.DateOnly, timeStr)
	if err != nil {
		return
	}
	st = dataTime.Unix()
	return
}
