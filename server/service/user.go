package service

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/constants/login"
	"base-system-backend/constants/table"
	userEnum "base-system-backend/constants/user"
	"base-system-backend/global"
	"base-system-backend/model/common/field"
	"base-system-backend/model/user"
	"base-system-backend/model/user/request"
	"base-system-backend/model/user/response"
	"base-system-backend/utils"
	"base-system-backend/utils/cache"
	"base-system-backend/utils/common"
	"base-system-backend/utils/jwt"
	"base-system-backend/utils/orm"
	userUtils "base-system-backend/utils/user"
	"base-system-backend/utils/validate"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type UserService struct{}

// AccountLoginService
//
//	@Description: 登陆 Service
//	@receiver UserService
//	@param params 登陆请求参数
//	@return err 登陆失败异常
//	@return debugInfo 错误调试信息

func (UserService) AccountLoginService(params *request.UserAccountLogin) (err error, debugInfo interface{}) {
	var instance user.User
	params.Account = strings.ToLower(params.Account)
	if err = global.DB.Table(table.User).
		Select("password", "status").
		Where("account = ?", params.Account).
		First(&instance).Error; err != nil {
		return errmsg.AccPwdInvalid, nil
	}
	// 停用
	if instance.Status == userEnum.AccStop {
		return errmsg.AccStop, nil
	}
	// 冻结
	if instance.Status == userEnum.AccFreeze {
		blackList := new(user.BlackList)
		err = global.DB.Table(table.UserBlackList).
			Select("next_time", "failed_num").
			Where("account = ?", params.Account).
			First(&blackList).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) && blackList != nil && time.Now().Unix() <= time.Time(blackList.NextTime).Unix() {
			//下次登陆时间大于当前,则仍不能登陆。返回剩余时间
			var nextLoginMinute int
			nextLoginMinute = int(math.Pow(2, float64(blackList.FailedNum-login.MaxLoginFailedNum)))
			if nextLoginMinute == 0 {
				nextLoginMinute = 1
			}
			debugInfo = map[string]interface{}{
				"next_time":  time.Time(blackList.NextTime).Unix(),
				"failed_num": blackList.FailedNum,
				"login_time": nextLoginMinute,
			}
			return errmsg.AccPwdInvalid, debugInfo
		}
	}
	// 密码错误
	if ok := utils.BcryptCheck(params.Password, instance.Password); !ok {
		// 登陆失败后,更新黑名单信息
		return errmsg.AccPwdInvalid, userUtils.LoginFiled(params.Account)
	}
	return
}

// LoginSuccess
//
//	@Description: 登陆成功 Service
//	@receiver UserService
//	@param c 上下文信息
//	@param loginBase 登陆成功基础参数
//	@return loginInfo 登陆成功响应信息
//	@return err 登陆失败异常
//	@return debugInfo 错误调试信息

func (UserService) LoginSuccess(c *gin.Context, loginBase *request.UserLoginBase) (loginInfo *response.LoginSuccess, err error, debugInfo interface{}) {
	u := new(user.User)
	if *loginBase.LoginType == login.KeycloakLogin {
		// keycloak 登陆成功
	} else {
		// 账号密码/短信登陆成功
		if loginBase.Phone != nil {
			global.DB.Table(table.User).Where("phone = ?", *loginBase.Phone).First(&u)
		} else {
			global.DB.Table(table.User).Where("account = ? or phone = ?", strings.ToLower(*loginBase.Account), *loginBase.Account).First(&u)
		}
		currentTime := &u.CurrentTime
		currentIp := &u.CurrentIp
		// 以前有登陆记录后把上次当前登陆时间/ip改为最后一次登陆时间/ip
		if currentTime != nil && currentIp != nil {
			u.LastTime = *currentTime
			u.LastIp = *currentIp
		}
		// 当前登陆时间
		u.CurrentTime = field.CustomTime(time.Now())
		// 当前登陆IP
		u.CurrentIp = utils.GetLoginIp(c)
		u.LoginType = *loginBase.LoginType
		// 修改用户状态为正常
		u.Status = userEnum.AccNormal
		if err = global.DB.Table(table.User).Save(&u).Error; err != nil {
			return nil, fmt.Errorf("登陆信息%w", errmsg.UpdateFailed), err.Error()
		}
		// 如果黑名单有错误记录则清除记录
		if err = global.DB.Table(table.UserBlackList).
			Where("account = ?", u.Account).Delete(&user.BlackList{}).Error; err != nil {
			return nil, fmt.Errorf("黑名单记录%w", errmsg.DeleteFailed), err.Error()
		}
	}
	// 设置token
	accessToken, err := jwt.GenToken(u.Id, u.Account)
	if err != nil {
		return nil, fmt.Errorf("token%w", errmsg.SaveFailed), err.Error()
	}
	cache.SetToken(u.Id, accessToken)
	// 组装数据
	if err = global.DB.Table(table.User).Where("account = ?", u.Account).First(&loginInfo).Error; err != nil {
		return nil, fmt.Errorf("登陆信息%w", errmsg.UpdateFailed), err.Error()
	}
	privilegeKeys, userRoleIds, err, debugInfo := utils.GetPrivilegeKeysByUserId(u.Id)
	if err != nil {
		return nil, err, debugInfo
	}
	loginInfo.PrivilegeList = privilegeKeys
	loginInfo.RoleIds = userRoleIds
	loginInfo.Token.Token = accessToken
	return
}

// UserListService
//
//	@Description: 用户列表 Service
//  @receiver UserService
//	@param c 上下文信息
//	@param params 查询参数
//	@return userList 用户列表
//	@return err 查询失败异常
//	@return debugInfo 错误调试信息

func (UserService) UserListService(c *gin.Context, params *request.UserFilter) (userList *response.UserList, err error, debugInfo interface{}) {
	tx := global.DB.Table(table.User)
	// 账号/姓名
	if params.Name != "" {
		tx = tx.Where("name LIKE ? ", fmt.Sprintf("%%%s%%", params.Name)).
			Or("account LIKE ? ", fmt.Sprintf("%%%s%%", params.Name))
	}
	//状态
	if c.Query("status") != "" {
		tx = tx.Where("status = ?", params.Status)
	}
	userList = new(response.UserList)
	if err = tx.Scopes(orm.Paginate(params.Page, params.PageSize)).
		Order("id").
		Find(&userList.Results).
		Limit(-1).Offset(-1).Count(&userList.TotalCount).Error; err != nil {
		return nil, fmt.Errorf("用户列表%w", errmsg.QueryFailed), err.Error()
	}
	for index, u := range userList.Results {
		privilegeKeys, userRoleIds, err, debugInfo := utils.GetPrivilegeKeysByUserId(u.Id)
		if err != nil {
			return nil, err, debugInfo
		}
		userList.Results[index].PrivilegeList = privilegeKeys
		userList.Results[index].RoleIds = userRoleIds
	}
	userList.GetPageInfo(&userList.PageInfo, params.Page, params.PageSize)
	return

}

// UserCreateService
//  @Description: 用户创建 Service
//  @receiver UserService
//  @param params 用户创建请求参数
//  @return err 创建/参数校验失败异常
//  @return debugInfo 错误调试信息

func (UserService) UserCreateService(params *request.UserCreate) (err error, debugInfo interface{}) {
	// 身份证校验
	if params.IdCard != "" && !validate.IdCardVerify(params.IdCard) {
		return fmt.Errorf("身份证号码%w", errmsg.Invalid), nil
	}
	// 手机号校验
	if params.Phone != "" && !validate.MobileVerify(params.Phone) {
		return fmt.Errorf("手机号码%w", errmsg.Invalid), nil
	}
	// 生成 SecretKey
	secretKey, err := utils.GenerateSecretKey()
	if err != nil {
		return fmt.Errorf("secretKey%w", errmsg.SaveFailed), err.Error()
	}
	params.SecretKey = secretKey
	// 账号统一小写
	params.Account = strings.ToLower(params.Account)
	// 密码加密
	params.Password = utils.BcryptHash(params.Password)
	// 全拼简拼
	params.FullName, params.ShortName = common.ConvertCnToLetter(params.Name)
	//使用分布式ID
	params.Id = utils.GenID()
	tx := global.DB.Begin()
	if err = tx.Table(table.User).Create(&params).Error; err != nil {
		return fmt.Errorf("用户%w", errmsg.SaveFailed), err.Error()
	}
	// 角色校验
	if params.RoleIds != nil {
		// 校验角色列表
		userRoles, err, debugInfo := validate.BindRoleVerify(params.Id, params.RoleIds)
		if err != nil {
			return fmt.Errorf("角色Id列表%w", errmsg.Invalid), debugInfo
		}
		if err = tx.Table(table.UserRole).Create(&userRoles).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("用户角色%w", errmsg.SaveFailed), err.Error()
		}

	}
	tx.Commit()
	return
}

// UserDetailService
//  @Description: 用户详情 Service
//  @receiver UserService
//  @param userId 用户 ID
//  @return userDetail 用户详情
//  @return err 查询失败异常
//  @return debugInfo 错误调试信息

func (UserService) UserDetailService(userId int64) (userDetail *response.UserDetail, err error, debugInfo interface{}) {
	if err = global.DB.Table(table.User).
		Where("id = ?", userId).
		First(&userDetail).Error; err != nil {
		return nil, fmt.Errorf("用户详情%w", errmsg.QueryFailed), err.Error()
	}
	privilegeKeys, userRoleIds, err, debugInfo := utils.GetPrivilegeKeysByUserId(userId)
	if err != nil {
		return nil, err, debugInfo
	}
	userDetail.PrivilegeList = privilegeKeys
	userDetail.RoleIds = userRoleIds
	return
}

// UserUpdateService
//  @Description: 更新当前用户 Service
//  @receiver UserService
//  @param userId 用户ID
//  @param params 更新请求参数
//  @return err 参数校验/更新失败异常
//  @return debugInfo 错误调试信息

func (UserService) UserUpdateService(userId int64, params *request.UserUpdate) (err error, debugInfo interface{}) {
	//身份证校验
	if params.IdCard != "" && !validate.IdCardVerify(params.IdCard) {
		return fmt.Errorf("身份证号码%w", errmsg.Invalid), nil
	}
	// 手机号校验
	if params.Phone != "" && !validate.MobileVerify(params.Phone) {
		return fmt.Errorf("手机号码%w", errmsg.Invalid), nil
	}
	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("用户%w", errmsg.UpdateFailed)
		}
	}()

	var u user.User
	if err = global.DB.Table(table.User).Where("id = ?", userId).First(&u).Error; err != nil {
		return fmt.Errorf("用户%w", errmsg.QueryFailed), err.Error()
	}
	if params.Name != "" && params.Name != u.Name {
		// 全拼简拼
		params.FullName, params.ShortName = common.ConvertCnToLetter(params.Name)
	}
	if err = tx.Model(&u).Updates(user.User{
		IdCard:    params.IdCard,
		Phone:     params.Phone,
		Email:     params.Email,
		Name:      params.Name,
		FullName:  params.FullName,
		ShortName: params.ShortName,
		Gender:    params.Gender}).Error; err != nil {
		return fmt.Errorf("用户%w", errmsg.UpdateFailed), err.Error()
	}
	// 绑定角色
	if params.RoleIds != nil {
		// 校验角色列表
		userRoles, err, debugInfo := validate.BindRoleVerify(userId, params.RoleIds)
		if err != nil {
			return fmt.Errorf("角色Id列表%w", errmsg.Invalid), debugInfo
		}
		// 删除旧绑定
		if err = tx.Table(table.UserRole).Where("user_id = ?", userId).Delete(&user.UserRole{}).Error; err != nil {
			return fmt.Errorf("用户角色%w", errmsg.DeleteFailed), err.Error()
		}
		// 重新绑定
		if err = tx.Table(table.UserRole).Create(&userRoles).Error; err != nil {
			return fmt.Errorf("用户角色%w", errmsg.SaveFailed), err.Error()
		}
	}
	tx.Commit()
	return
}

// UserChangePwdByPwdService
//  @Description: 修改当前用户密码 Service
//  @receiver UserService
//  @param u 当前登陆用户
//  @param params 修改密码请求参数
//  @return err 参数校验/修改失败异常
//  @return debugInfo 错误调试信息

func (UserService) UserChangePwdByPwdService(u *user.User, params *request.PwdChangeByPwd) (err error, debugInfo interface{}) {
	if ok := utils.BcryptCheck(params.OldPassword, u.Password); !ok {
		return fmt.Errorf("原密码%w", errmsg.Incorrect), nil
	}
	if err = global.DB.Model(&u).Update("password", utils.BcryptHash(params.NewPassword)).Error; err != nil {
		return fmt.Errorf("密码%w", errmsg.UpdateFailed), err.Error()
	}
	// 修改成功删除 token
	cache.DeleteToken(u.Id)
	return
}

// UserUploadAvatarService
//  @Description: 上传头像 Service
//  @receiver UserService
//  @param c 上下文信息
//  @param user 当前用户
//  @param fileHeader 文件请求头
//  @return err 头像校验/保存失败异常
//  @return debugInfo 错误调试信息

func (UserService) UserUploadAvatarService(c *gin.Context, user *user.User, fileHeader *multipart.FileHeader) (err error, debugInfo interface{}) {
	// 拼接路径
	savePath := strings.Join(global.CONFIG.Static.Avatar, "")
	if err, debugInfo = common.FileCheck(savePath); err != nil {
		return
	}
	avatarAbsPath := strings.Join([]string{savePath, "/", strconv.FormatInt(user.Id, 10), ".jpg"}, "")
	if err = c.SaveUploadedFile(fileHeader, avatarAbsPath); err != nil {
		return fmt.Errorf("头像%w", errmsg.UpdateFailed), err.Error()
	}
	pathSlice := strings.Split(avatarAbsPath, "static")
	avatarPath := strings.Join([]string{"/static", pathSlice[1]}, "")
	// 保存到数据库
	if err = global.DB.Model(&user).Update("avatar", avatarPath).Error; err != nil {
		return fmt.Errorf("头像%w", errmsg.UpdateFailed), err.Error()
	}
	return
}

// UserResetPwdByIdService
//  @Description: 重置指定账号密码 Service
//  @receiver UserService
//  @param userId 用户 ID
//  @param params 重置密码请求参数
//  @return err 校验/重置失败异常
//  @return debugInfo 错误调试信息

func (UserService) UserResetPwdByIdService(userId string, params *request.PwdChangeById) (err error, debugInfo interface{}) {
	var u user.User
	if err = global.DB.Table(table.User).Where("id = ?", userId).First(&u).Error; err != nil {
		return fmt.Errorf("用户%w", errmsg.QueryFailed), err.Error()
	}
	//不能重置管理员账号密码
	if u.IsSystem == true {
		return fmt.Errorf(errmsg.NotPrivilege.Error(), "修改管理员密码"), nil
	}
	//禁用/冻结
	if u.Status == userEnum.AccStop || u.Status == userEnum.AccFreeze {
		return fmt.Errorf(errmsg.ResetPwdFailed.Error(), u.Status.Choices()), nil
	}
	//更新密码
	if err = global.DB.Model(&u).
		Updates(user.User{Password: utils.BcryptHash(params.Password), Status: userEnum.AccChangePwd}).Error; err != nil {
		return fmt.Errorf("用户密码%w", errmsg.UpdateFailed), err.Error()
	}
	return
}

// UserStatusChangeByIdService
//  @Description: 修改指定用户状态 Service
//  @receiver UserService
//  @param userId 用户 ID
//  @param params 状态
//  @return err 参数校验/修改失败异常
//  @return debugInfo 错误调试信息

func (UserService) UserStatusChangeByIdService(userId string, params *request.StatusChangeById) (err error, debugInfo interface{}) {
	var u user.User
	if err = global.DB.Table(table.User).Where("id =?", userId).First(&u).Error; err != nil {
		return fmt.Errorf("用户%w", errmsg.QueryFailed), err.Error()
	}
	// 不能修改管理员账号状态
	if u.IsSystem == true {
		return fmt.Errorf(errmsg.NotPrivilege.Error(), "修改管理员状态"), nil
	}
	// 只能启动或停用
	if params.Status == userEnum.AccFreeze || params.Status == userEnum.AccChangePwd {
		return errmsg.OnlyStopOrEnable, nil
	}
	//更新状态
	if err = global.DB.Model(&u).Update("status", params.Status).Error; err != nil {
		return fmt.Errorf("用户状态%w", errmsg.UpdateFailed), err.Error()
	}
	//修改成功清除token
	cache.DeleteToken(u.Id)
	return
}

// UserDetailByIdService
//  @Description: 查询指定用户详情 Service
//  @receiver UserService
//  @param userId 用户 ID
//  @return userDetail 用户详情信息
//  @return err 用户不存在异常
//  @return debugInfo 错误调试信息

func (UserService) UserDetailByIdService(userId string) (userDetail *response.UserDetail, err error, debugInfo interface{}) {
	if err = global.DB.Table(table.User).Where("id =?", userId).First(&userDetail).Error; err != nil {
		return nil, fmt.Errorf("用户%w", errmsg.QueryFailed), err.Error()
	}
	privilegeKeys, userRoleIds, err, debugInfo := utils.GetPrivilegeKeysByUserId(userDetail.Id)
	if err != nil {
		return nil, err, debugInfo
	}
	userDetail.PrivilegeList = privilegeKeys
	userDetail.RoleIds = userRoleIds
	return
}

// UserUpdateByIdService
//  @Description: 更新指定用户信息 Service
//  @receiver UserService
//  @param userId 用户 ID
//  @param params 用户信息请求参数
//  @return err 参数校验/更新失败异常
//  @return debugInfo 错误调试信息

func (UserService) UserUpdateByIdService(userId string, params *request.UserUpdateById) (err error, debugInfo interface{}) {
	// 身份证校验
	if params.IdCard != "" && !validate.IdCardVerify(params.IdCard) {
		return fmt.Errorf("身份证号码%w", errmsg.Invalid), nil
	}
	var u user.User
	if err = global.DB.Table(table.User).Where("id = ?", userId).First(&u).Error; err != nil {
		return fmt.Errorf("用户%w", errmsg.QueryFailed), err.Error()
	}

	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = fmt.Errorf("用户%w", errmsg.UpdateFailed)
		}
	}()

	if params.Name != "" && params.Name != u.Name {
		//全拼简拼
		params.FullName, params.ShortName = common.ConvertCnToLetter(params.Name)
	}
	if err = tx.Model(&u).Updates(user.User{
		IdCard:    params.IdCard,
		Email:     params.Email,
		Name:      params.Name,
		FullName:  params.FullName,
		ShortName: params.ShortName,
		Gender:    params.Gender,
	}).Error; err != nil {
		return fmt.Errorf("用户%w", errmsg.UpdateFailed), err.Error()
	}
	//绑定角色
	if params.RoleIds != nil {
		//校验角色列表
		userRoles, err, debugInfo := validate.BindRoleVerify(u.Id, params.RoleIds)
		if err != nil {
			return fmt.Errorf("角色Id列表%w", errmsg.Invalid), debugInfo
		}
		//删除旧绑定
		if err = tx.Table(table.UserRole).Where("user_id =?", u.Id).Delete(&user.UserRole{}).Error; err != nil {
			return fmt.Errorf("用户绑定角色%w", errmsg.DeleteFailed), err.Error()
		}
		//重新绑定
		if err = tx.Table(table.UserRole).Create(&userRoles).Error; err != nil {
			return fmt.Errorf("用户角色%w", errmsg.SaveFailed), err.Error()
		}
	}
	tx.Commit()
	return
}
