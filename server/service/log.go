package service

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/table"
	"base-system-backend/global"
	"base-system-backend/model/log"
	"base-system-backend/model/log/request"
	"base-system-backend/model/log/response"
	"base-system-backend/model/user"
	"base-system-backend/utils"
	"base-system-backend/utils/common"
	"base-system-backend/utils/orm"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"strconv"
	"time"
)

type LogService struct{}

// OperateLogListService
//  @Description: 操作日志列表 service
//  @receiver l 接收者
//  @param c 上下文信息
//  @param params 查询参数
//  @return operateLogList 操作日志列表
//  @return err 查询失败异常
//  @return debugInfo 错误调试信息

func (l LogService) OperateLogListService(c *gin.Context, params *request.OperateLogFilter) (
	operateLogList *response.OperateLogList, err error, debugInfo interface{}) {
	return l.operateLogQuery(true, c, params)
}

// OperateLogDownloadService
//  @Description: 下载操作日志 service
//  @receiver l 接收者
//  @param c 上下文信息
//  @param params 查询参数
//  @return content 操作日志二进制
//  @return err 查询失败异常
//  @return debugInfo 错误调试信息

func (l LogService) OperateLogDownloadService(c *gin.Context, params *request.OperateLogFilter) (content io.ReadSeeker, err error, debugInfo interface{}) {
	var (
		operateLogList *response.OperateLogList
		res            []interface{}
		userID         string
		accessTime     string
	)
	operateLogList, err, debugInfo = l.operateLogQuery(false, c, params)
	if err != nil {
		return
	}
	for _, value := range operateLogList.Results {
		if value.UserId != nil {
			userID = strconv.FormatInt(*value.UserId, 10)
		}
		if value.AccessTime != nil {
			accessTime = value.AccessTime.String()
		}
		success := "失败"
		if value.Success == true {
			success = "成功"
		}
		res = append(res, &response.OperateLogDownload{
			Id:          value.Id,
			ActionName:  value.ActionName,
			Module:      value.Module,
			AccessUrl:   value.AccessUrl,
			RequestIp:   value.RequestIp,
			UserAgent:   value.UserAgent,
			UserId:      userID,
			UserName:    value.UserName,
			UserAccount: value.UserAccount,
			AccessTime:  accessTime,
			Success:     success,
			Message:     value.Message,
		})
	}
	content, err, debugInfo = utils.ToExcel("操作日志", []string{`Id`, `接口名称`, `模块名称`, `接口地址`, `来源IP`, `UserAgent`, `用户ID`, `用户姓名`, `用户账号`, `请求时间`, `请求状态`, `错误信息`}, res)

	return
}

// operateLogQuery
//  @Description: 操作日志查询 logic
//  @param isPage 是否需要分页
//  @param c 上下文信息
//  @param params 查询参数
//  @return operateLogList 操作日志列表
//  @return err 查询失败异常
//  @return debugInfo 错误调试信息

func (LogService) operateLogQuery(isPage bool, c *gin.Context, params *request.OperateLogFilter) (operateLogList *response.OperateLogList, err error, debugInfo interface{}) {
	tx := global.DB.Table(table.OperateLog)
	// 行为名称
	if params.ActionName != "" {
		tx = tx.Where("action_name LIKE ?", fmt.Sprintf("%%%s%%", params.ActionName))
	}
	// 模块名称
	if params.Module != "" {
		tx = tx.Where("model = ?", params.Module)
	}
	// 访问时的ip
	if params.RequestIp != "" {
		tx = tx.Where("request_ip = ?", params.RequestIp)
	}
	// 请求者的用户ID
	if params.UserId != nil {
		tx = tx.Where("user_id = ?", params.UserId)
	}
	// 操作是否成功
	if c.Query("success") != "" {
		tx = tx.Where("success = ?", c.Query("success"))
	}
	// 访问开始时间
	if params.StartAccessTime != nil {
		tx = tx.Where("access_time >= ?", common.Timestamp2Datetime(*params.StartAccessTime))
	}
	// 访问结束时间
	if params.EndAccessTime != nil {
		tx = tx.Where("access_time <= ?", common.Timestamp2Datetime(*params.EndAccessTime))
	}
	operateLogList = new(response.OperateLogList)
	if isPage {
		if err = tx.Scopes(orm.Paginate(params.Page, params.PageSize)).
			Order("id DESC").
			Find(&operateLogList.Results).
			Limit(-1).Offset(-1).Count(&operateLogList.TotalCount).Error; err != nil {
			return nil, fmt.Errorf("操作日志列表%w", errmsg.QueryFailed), err.Error()
		}
	} else {
		if err = tx.Order("id ASC").
			Find(&operateLogList.Results).Error; err != nil {
			return nil, fmt.Errorf("操作日志列表%w", errmsg.QueryFailed), err.Error()
		}
	}
	//组装响应结果
	for index, operateLog := range operateLogList.Results {
		var userInfo user.User
		if operateLog.UserId != nil {
			if err = global.DB.Table(table.User).
				Select("name", "account").
				Where("id = ?", operateLog.UserId).Take(&userInfo).Error; err != nil {
				return nil, fmt.Errorf("用户信息%w", errmsg.QueryFailed), err.Error()
			}
		}
		//用户名
		operateLogList.Results[index].UserName = userInfo.Name
		// 账号
		operateLogList.Results[index].UserAccount = userInfo.Account
		detail := operateLogList.Results[index].Detail
		//错误信息判断
		if detail != nil {
			var errMsg response.OperateLogErrMessage
			err = json.Unmarshal([]byte(detail.String()), &errMsg)
			if err != nil {
				return nil, errmsg.JsonConvertFiled, err.Error()
			}
			operateLogList.Results[index].Message = errMsg.Message
		}
	}
	operateLogList.GetPageInfo(&operateLogList.PageInfo, params.Page, params.PageSize)
	return
}

// DeleteOperateLog
//  @Description: 登录成功删除一个月以前的操作日志

func (LogService) DeleteOperateLog() {
	day := time.Now().AddDate(0, -1, 0)
	if err := global.DB.Where("access_time <=?", day).Delete(&log.OperateLog{}).Error; err != nil {
		global.LOG.Error("login success delete operation log failed: ", zap.Error(err))
	}
}
