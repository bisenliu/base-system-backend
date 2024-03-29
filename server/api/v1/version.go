package v1

import (
	"base-system-backend/constants/code"
	"base-system-backend/constants/errmsg"
	"base-system-backend/global"
	"base-system-backend/model/common/response"
	"base-system-backend/utils/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type VersionApi struct{}

// GetVersionApi
// @Summary 获取版本号
// @Description 获取版本号
// @Tags VersionApi
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Data{data=version.Version}
// @Router /version/ [get]
func (VersionApi) GetVersionApi(c *gin.Context) {
	versionByte, err := global.FS.ReadFile("version.txt")
	if err != nil {
		response.Error(c, code.QueryFailed, fmt.Sprintf(errmsg.ReadFailed.Error(), "文件"), err.Error())
		return
	}
	submitTime, err := strconv.ParseInt(string(versionByte), 10, 64)
	if err != nil {
		response.Error(c, code.QueryFailed, fmt.Errorf("文件日期%w", errmsg.Invalid), err.Error())
		return
	}
	projectStartTime, err := common.TimeStr2TimeStamp(global.CONFIG.System.StartTime)
	if err != nil {
		response.Error(c, code.QueryFailed, errmsg.TimeCalcFiled, err.Error())
		return
	}
	str := strconv.FormatFloat(float64((submitTime-projectStartTime)/(3600*24)), 'f', 3, 64)
	version := global.CONFIG.System.Version
	info := strings.Join([]string{version, ".", str}, "")
	response.OK(c, map[string]string{"version": info})
}
