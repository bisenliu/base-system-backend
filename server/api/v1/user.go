package v1

import (
	"base-system-backend/constants/code"
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/login"
	userEnum "base-system-backend/constants/user"
	"base-system-backend/model/common/response"
	"base-system-backend/model/user/request"
	"base-system-backend/utils"
	"base-system-backend/utils/cache"
	"base-system-backend/utils/validate"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// UserLoginApi
// @Summary 登陆
// @Description 登陆
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object body request.UserLoginBase true "登陆参数"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.LoginSuccess}
// @Router /user/login/ [post]
func (UserApi) UserLoginApi(c *gin.Context) {
	loginBase := new(request.UserLoginBase)
	if !validate.RequestParamsVerify(c, &loginBase) {
		return
	}
	// 账号密码登陆
	if *loginBase.LoginType == login.AccPwdLogin {
		accLoginParams := new(request.UserAccountLogin)
		if !validate.RequestParamsVerify(c, accLoginParams) {
			return
		}
		// 账号密码登陆逻辑
		if err, debugInfo := userService.AccountLoginService(accLoginParams); err != nil {
			response.Error(c, code.InvalidLogin, err, debugInfo)
			return
		}
	} else if *loginBase.LoginType == login.KeycloakLogin {
		// todo Keycloak 登陆
		panic("Keycloak login api unrealized...")
	} else {
		panic("sms login api unrealized...")
	}
	// 登陆参数校验成功校验成功生成token, 记录ip ...
	loginInfo, err, debugInfo := userService.LoginSuccess(c, loginBase)
	if err != nil {
		response.Error(c, code.InvalidLogin, err, debugInfo)
		return
	}
	// 删除一个月以前的操作日志
	go logService.DeleteOperateLog()
	response.OK(c, loginInfo)
}

// UserLogoutApi
// @Summary 登出
// @Description 登出
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /user/logout/ [post]
func (UserApi) UserLogoutApi(c *gin.Context) {
	user, err, debugInfo := utils.GetCurrentUser(c)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	//todo keycloak 也需要退出登陆
	//清除token
	cache.DeleteToken(user.Id)
	response.OK(c, nil)
}

// UserListApi
// @Summary 用户列表
// @Description 用户列表
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object query request.UserFilter false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.UserList}
// @Router /user/list/ [get]
func (UserApi) UserListApi(c *gin.Context) {
	params := new(request.UserFilter)
	if !validate.QueryParamsVerify(c, &params) {
		return
	}
	userList, err, debugInfo := userService.UserListService(c, params)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	response.OK(c, userList)
}

// UserCreateApi
// @Summary 用户创建
// @Description 用户创建
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object body request.UserCreate true "用户信息"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.Create}
// @Router /user/ [post]
func (UserApi) UserCreateApi(c *gin.Context) {
	params := new(request.UserCreate)
	if !validate.RequestParamsVerify(c, params) {
		return
	}
	if err, debugInfo := userService.UserCreateService(params); err != nil {
		response.Error(c, code.SaveFailed, err, debugInfo)
		return

	}
	response.OK(c, map[string]int64{"id": params.Id})
}

// UserDetailApi
// @Summary 用户详情
// @Description 用户详情
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.UserDetail}
// @Router /user/detail/ [get]
func (UserApi) UserDetailApi(c *gin.Context) {
	u, err, debugInfo := utils.GetCurrentUser(c)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	userDetail, err, debugInfo := userService.UserDetailService(u.Id)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	response.OK(c, *userDetail)
}

// UserUpdateApi
// @Summary 用户修改
// @Description 用户修改
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object body request.UserUpdate true "用户修改信息"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /user/detail/ [put]
func (UserApi) UserUpdateApi(c *gin.Context) {
	params := new(request.UserUpdate)
	if !validate.RequestParamsVerify(c, params) {
		return
	}
	u, err, debugInfo := utils.GetCurrentUser(c)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	if err, debugInfo = userService.UserUpdateService(u.Id, params); err != nil {
		response.Error(c, code.UpdateFailed, err, debugInfo)
		return
	}
	response.OK(c, nil)
}

// UserChangePwdApi
// @Summary 用户密码修改
// @Description 用户密码修改
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object body request.UserChangePwdBase true "用户密码信息"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /user/password/ [patch]
func (UserApi) UserChangePwdApi(c *gin.Context) {
	params := new(request.UserChangePwdBase)
	if !validate.RequestParamsVerify(c, params) {
		return
	}
	u, err, debugInfo := utils.GetCurrentUser(c)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	// 密码修改
	if params.Type == userEnum.PwdChange {
		pwdChangeParams := new(request.PwdChangeByPwd)
		if !validate.RequestParamsVerify(c, pwdChangeParams) {
			return
		}
		if err, debugInfo = userService.UserChangePwdByPwdService(u, pwdChangeParams); err != nil {
			response.Error(c, code.UpdateFailed, err, debugInfo)
			return
		}

	} else {
		panic("sms change password api unrealized...")
	}
	response.OK(c, nil)
}

// UserUploadAvatarApi
// @Summary 用户头像修改
// @Description 用户头像修改
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param avatar formData file true "头像"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /user/avatar/ [patch]
func (UserApi) UserUploadAvatarApi(c *gin.Context) {
	fileHeader, err := c.FormFile("avatar")
	// 文件不存在
	if fileHeader == nil {
		response.Error(c, code.SaveFailed, fmt.Errorf("头像文件%w", errmsg.Required), nil)
		return
	}
	// 读取失败
	if err != nil {
		response.Error(c, code.SaveFailed, fmt.Errorf("头像文件%w", errmsg.Invalid), err.Error())
		return
	}
	// 头像文件校验
	if err, debugInfo := validate.ImageVerify(fileHeader); err != nil {
		response.Error(c, code.SaveFailed, err, debugInfo)
		return
	}
	// 获取当前登陆用户
	u, err, debugInfo := utils.GetCurrentUser(c)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	// 上传头像
	if err, debugInfo = userService.UserUploadAvatarService(c, u, fileHeader); err != nil {
		response.Error(c, code.SaveFailed, err, debugInfo)
		return
	}
	response.OK(c, nil)
}

// UserResetPwdByIdApi
// @Summary 重置指定账号密码
// @Description 重置指定账号密码
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object body request.PwdChangeById true "用户密码信息"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /user/:user_id/password/ [put]
func (UserApi) UserResetPwdByIdApi(c *gin.Context) {
	params := new(request.PwdChangeById)
	if !validate.RequestParamsVerify(c, &params) {
		return
	}
	userId := c.Param("user_id")
	if err, debugInfo := userService.UserResetPwdByIdService(userId, params); err != nil {
		response.Error(c, code.UpdateFailed, err, debugInfo)
		return
	}
	response.OK(c, nil)
}

// UserStatusChangeByIdApi
// @Summary 修改指定账户状态
// @Description 修改指定账户状态
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object body request.StatusChangeById true "用户状态"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /user/:user_id/status/ [put]
func (UserApi) UserStatusChangeByIdApi(c *gin.Context) {
	params := new(request.StatusChangeById)
	if !validate.RequestParamsVerify(c, &params) {
		return
	}

	userId := c.Param("user_id")
	if err, debugInfo := userService.UserStatusChangeByIdService(userId, params); err != nil {
		response.Error(c, code.UpdateFailed, err, debugInfo)
		return
	}
	response.OK(c, nil)
}

// UserDetailByIdApi
// @Summary 查询指定用户信息
// @Description 查询指定用户信息
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.UserDetail}
// @Router /user/:user_id/ [get]
func (UserApi) UserDetailByIdApi(c *gin.Context) {
	userId := c.Param("user_id")
	userDetail, err, debugInfo := userService.UserDetailByIdService(userId)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	response.OK(c, userDetail)
}

// UserUpdateByIdApi
// @Summary 修改指定用户信息
// @Description 修改指定用户信息
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Identification header string true "Token 令牌"
// @Param object body request.UserUpdateById true "用户信息"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data
// @Router /user/:user_id/ [put]
func (UserApi) UserUpdateByIdApi(c *gin.Context) {
	params := new(request.UserUpdateById)
	if !validate.RequestParamsVerify(c, params) {
		return
	}
	userId := c.Param("user_id")
	err, debugInfo := userService.UserUpdateByIdService(userId, params)
	if err != nil {
		response.Error(c, code.UpdateFailed, err, debugInfo)
		return
	}
	response.OK(c, nil)
}
