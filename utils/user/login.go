package user

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/login"
	"base-system-backend/constants/table"
	userEnum "base-system-backend/constants/user"
	"base-system-backend/global"
	"base-system-backend/model/user"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"math"
	"time"
)

func LoginFiled(account string) (debugInfo interface{}) {
	blackList, err, debugInfo := addLoginFailedNum(account)
	if err != nil {
		return
	}
	if blackList.FailedNum >= login.LoginFailedMaxNum {
		nextLoginMinute := int(math.Pow(2, float64(blackList.FailedNum-login.LoginFailedMaxNum)))
		debugInfo = map[string]interface{}{
			"next_time":  time.Time(blackList.NextTime).Unix() * 1000,
			"failed_num": blackList.FailedNum,
			"login_time": nextLoginMinute,
		}
		return
	}
	return debugInfo
}

func addLoginFailedNum(account string) (blackList *user.BlackList, err error, debugInfo interface{}) {
	var userExists bool
	u := new(user.User)
	// 获取用户实例
	if err = global.DB.Table(table.User).Where("account = ? or phone = ?", account, account).First(&u).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		userExists = false
	} else if err != nil {
		return nil, fmt.Errorf("用户%w", errmsg.QueryFailed), err.Error()
	} else {
		userExists = true
	}
	//登录失败次数
	failedNum := getUserLoginFailedNum(account)
	// 获取黑名单实例
	err = global.DB.Table(table.UserBlackList).Where("account = ? ", account).First(&blackList).Error
	// 不存在添加到黑名单
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newBlacklist := user.BlackList{Account: account, FailedNum: failedNum}
		if err = global.DB.Table(table.UserBlackList).Create(&newBlacklist).Error; err != nil {
			return nil, fmt.Errorf("黑名单%w", errmsg.QueryFailed), err.Error()
		}
		return
	}
	if err != nil {
		return nil, fmt.Errorf("黑名单%w", errmsg.QueryFailed), err.Error()
	}
	// 登录失败五次
	if failedNum == login.LoginFailedMaxNum {
		// 账号存在则修改状态为冻结
		if userExists {
			if err = global.DB.Model(&u).Update("status", userEnum.AccFreeze).Error; err != nil {
				return nil, fmt.Errorf("用户%w", errmsg.UpdateFailed), err.Error()
			}
		}
		// 失败五次,下次登录时间为1分钟后
		if err, debugInfo = blackListTimeAddMinute(blackList, (failedNum-login.LoginFailedMaxNum)+1); err != nil {
			return nil, err, debugInfo
		}
	} else if failedNum > login.LoginFailedMaxNum {
		// 大于五次登录失败时间翻倍
		if err, debugInfo = blackListTimeAddMinute(blackList, ((failedNum-login.LoginFailedMaxNum)+1)*2); err != nil {
			return nil, err, debugInfo
		}
	}
	if err = global.DB.Model(&blackList).Update("failed_num", failedNum).Error; err != nil {
		return nil, fmt.Errorf("黑名单失败次数%w", errmsg.UpdateFailed), err.Error()
	}
	return

}

func getUserLoginFailedNum(account string) (failedNum int) {
	if err := global.DB.Table(table.UserBlackList).
		Select("failed_num").Where("account = ?", account).Take(&failedNum).
		Error; errors.Is(err, gorm.ErrRecordNotFound) {
		failedNum = 1
		return
	}
	failedNum += 1
	return
}

func blackListTimeAddMinute(blacklist *user.BlackList, nextLoginMinute int) (err error, debugInfo interface{}) {
	nowTime := time.Now()
	parseTime, err := time.ParseDuration(fmt.Sprintf("%dm", nextLoginMinute))
	if err != nil {
		return errmsg.TimeCalcFiled, err.Error()
	}
	if err = global.DB.Model(&blacklist).Update("next_time", nowTime.Add(parseTime)).Error; err != nil {
		return fmt.Errorf("下次登录时间%w", errmsg.UpdateFailed), err.Error()
	}
	return
}
