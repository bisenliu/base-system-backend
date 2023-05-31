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
	return tm.Format("2006-01-02")
}
