package field

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type CUTime struct {
	CreateTime CustomTime `gorm:"column:create_time;autoCreateTime:milli;comment:创建时间"`
	UpdateTime CustomTime `gorm:"column:update_time;autoCreateTime:milli;comment:更新时间"`
}

type CustomTime time.Time

func (c *CustomTime) MarshalJSON() ([]byte, error) {
	//格式化毫秒
	tTime := time.Time(*c)
	seconds := tTime.UnixMilli()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}

func (c CustomTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tTime := time.Time(c)
	if tTime.UnixMilli() == zeroTime.UnixMilli() {
		return nil, nil
	}
	return tTime, nil
}

func (c *CustomTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*c = CustomTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
