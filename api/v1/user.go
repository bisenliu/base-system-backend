package v1

import (
	"base-system-backend/enums/code"
	"base-system-backend/enums/errmsg"
	"base-system-backend/enums/login"
	userEnum "base-system-backend/enums/user"
	"base-system-backend/model/common/response"
	"base-system-backend/model/user/request"
	"base-system-backend/utils"
	"base-system-backend/utils/cache"
	"base-system-backend/utils/validate"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (UserApi) UserLoginApi(c *gin.Context) {
	loginBase := new(request.UserLoginBase)
	if ok := validate.RequestParamsVerify(c, &loginBase); !ok {
		return
	}
	// 账号密码登录
	if *loginBase.LoginType == login.AccPwdLogin {
		accLoginParams := new(request.UserAccountLogin)
		if ok := validate.RequestParamsVerify(c, accLoginParams); !ok {
			return
		}
		// 账号密码登录逻辑
		if err, debugInfo := userService.AccountLoginService(accLoginParams); err != nil {
			response.Error(c, code.InvalidLogin, err, debugInfo)
			return
		}
	} else if *loginBase.LoginType == login.KeycloakLogin {
		// todo Keycloak 登录
		panic("Keycloak login api unrealized...")
	} else {
		panic("sms login api unrealized...")
	}
	// 登录参数校验成功校验成功生成token, 记录ip ...
	loginInfo, err, debugInfo := userService.LoginSuccess(c, loginBase)
	if err != nil {
		response.Error(c, code.InvalidLogin, err, debugInfo)
		return
	}
	response.OK(c, loginInfo)
	return
}

func (UserApi) UserLogoutApi(c *gin.Context) {
	user, err, debugInfo := utils.GetCurrentUser(c)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	//todo keycloak 登录也需要退出受录
	//清除token
	cache.DeleteToken(user.Id)
	response.OK(c, nil)
	return
}

// UserListApi
// @Summary 用户列表
// @Description 用户列表
// @Tags UserApi
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object query request.UserFilter false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} response.Data{data=response.UserList}
// @Router /user/list/ [get]
func (UserApi) UserListApi(c *gin.Context) {
	params := new(request.UserFilter)
	if ok := validate.QueryParamsVerify(c, &params); !ok {
		return
	}
	userList, err, debugInfo := userService.UserListService(c, params)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	response.OK(c, userList)
	return
}

func (UserApi) UserCreateApi(c *gin.Context) {
	params := new(request.UserCreate)
	if ok := validate.RequestParamsVerify(c, params); !ok {
		return
	}
	if err, debugInfo := userService.UserCreateService(params); err != nil {
		response.Error(c, code.SaveFailed, err, debugInfo)
		return

	}
	response.OK(c, map[string]int64{"id": params.Id})
	return
}

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
	return
}

func (UserApi) UserUpdateApi(c *gin.Context) {
	params := new(request.UserUpdate)
	if ok := validate.RequestParamsVerify(c, params); !ok {
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
	return
}

func (UserApi) UserChangePwdApi(c *gin.Context) {
	params := new(request.UserChangePwdBase)
	if ok := validate.RequestParamsVerify(c, params); !ok {
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
		if ok := validate.RequestParamsVerify(c, pwdChangeParams); !ok {
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
	return
}

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
	// 获取当前登录用户
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
	return
}

func (UserApi) UserResetPwdByIdApi(c *gin.Context) {
	params := new(request.PwdChangeById)
	if ok := validate.RequestParamsVerify(c, &params); !ok {
		return
	}
	userId := c.Param("user_id")
	if err, debugInfo := userService.UserResetPwdByIdService(userId, params); err != nil {
		response.Error(c, code.UpdateFailed, err, debugInfo)
		return
	}
	response.OK(c, nil)
	return
}

func (UserApi) UserStatusChangeByIdApi(c *gin.Context) {
	params := new(request.StatusChangeById)
	if ok := validate.RequestParamsVerify(c, &params); !ok {
		return
	}

	userId := c.Param("user_id")
	if err, debugInfo := userService.UserStatusChangeByIdService(userId, params); err != nil {
		response.Error(c, code.UpdateFailed, err, debugInfo)
		return
	}
	response.OK(c, nil)
	return
}

func (UserApi) UserDetailByIdApi(c *gin.Context) {
	userId := c.Param("user_id")
	userDetail, err, debugInfo := userService.UserDetailByIdService(userId)
	if err != nil {
		response.Error(c, code.QueryFailed, err, debugInfo)
		return
	}
	response.OK(c, userDetail)
	return
}

func (UserApi) UserUpdateByIdApi(c *gin.Context) {
	params := new(request.UserUpdateById)
	if ok := validate.RequestParamsVerify(c, params); !ok {
		return
	}
	userId := c.Param("user_id")
	err, debugInfo := userService.UserUpdateByIdService(userId, params)
	if err != nil {
		response.Error(c, code.UpdateFailed, err, debugInfo)
		return
	}
	response.OK(c, nil)
	return
}
