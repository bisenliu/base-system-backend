package field

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// CUTime
//
//	@Description: 创建/更新数据时自动添加时间
type CUTime struct {
	CreateTime CustomTime `json:"create_time" gorm:"column:create_time;autoCreateTime;comment:创建时间" swaggertype:"integer"`
	UpdateTime CustomTime `json:"update_time" gorm:"column:update_time;autoUpdateTime;comment:更新时间" swaggertype:"integer"`
}

type CustomTime time.Time

func (c *CustomTime) MarshalJSON() ([]byte, error) {
	//格式化毫秒
	tTime := time.Time(*c)
	seconds := tTime.Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}

func (c CustomTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tTime := time.Time(c)
	if tTime.Unix() == zeroTime.Unix() {
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

func (c *CustomTime) String() string {
	return fmt.Sprintf(time.Time(*c).Format(time.DateTime))
}
