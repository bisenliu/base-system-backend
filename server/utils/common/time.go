package common

import (
	"strconv"
	"time"
)

// Timestamp2Datetime
//  @Description: 时间戳转时间
//  @param timestamp 时间戳
//  @return string 转换后的时间

func Timestamp2Datetime(timestamp int64) string {
	timestampLen := len(strconv.FormatInt(timestamp, 10))
	if timestampLen != 10 {
		timestamp /= 1000
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format(time.DateOnly)
}

// TimeStr2TimeStamp
//  @Description: 时间字符串转时间戳
//  @param timeStr 时间字符串
//  @return st 时间戳
//  @return err 转换失败异常

func TimeStr2TimeStamp(timeStr string) (st int64, err error) {
	dataTime, err := time.Parse(time.DateOnly, timeStr)
	if err != nil {
		return
	}
	st = dataTime.Unix()
	return
}
