package internal

import (
	"base-system-backend/enums/table"
	"base-system-backend/global"
	"base-system-backend/model/privilege"
	"base-system-backend/model/role"
	"base-system-backend/model/user"
	"base-system-backend/utils"
	"base-system-backend/utils/common"
	"encoding/json"
	"fmt"
	"gorm.io/datatypes"
	"os"
	"strconv"
	"strings"
)

func DefaultPrivilegeInit() {
	err := global.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table.Privilege)).Error
	if err != nil {
		panic(fmt.Errorf("clear privilege table failed: (%s)", err.Error()))
	}
	baseDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("get current dir failed: (%s)", err.Error()))
	}
	privilegeJsonPath := strings.Join([]string{baseDir, "/initialize/internal/privilege.json"}, "")
	bytePrivilege, err := os.ReadFile(privilegeJsonPath)
	if err != nil {
		panic(fmt.Errorf("read privilege json file failed: (%s)", err.Error()))
	}
	var Privileges struct {
		RECORDS []map[string]string `json:"records"`
	}
	err = json.Unmarshal(bytePrivilege, &Privileges)
	if err != nil {
		panic(fmt.Errorf("privilege json convert failed: (%s)", err.Error()))
	}
	var data []privilege.Privilege
	for _, value := range Privileges.RECORDS {
		id, _ := strconv.ParseInt(value["id"], 10, 64)
		parentId, _ := strconv.ParseInt(value["parent_id"], 10, 64)
		title, _ := value["title"]
		key, _ := value["key"]
		var dependency datatypes.JSON
		_ = json.Unmarshal([]byte(value["dependency"]), &dependency)
		data = append(data, privilege.Privilege{
			Id:         id,
			ParentId:   parentId,
			Title:      title,
			Key:        key,
			Dependency: dependency,
		})
	}
	err = global.DB.Table(table.Privilege).CreateInBatches(&data, len(data)).Error
	if err != nil {
		panic(fmt.Errorf("init privilege failed: (%s)", err.Error()))
	}
}

func DefaultRoleInit() {
	err := global.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table.Role)).Error
	if err != nil {
		panic(fmt.Errorf("clear role table failed: (%s)", err.Error()))
	}
	// 管理员角色
	var adminPrivilegeKeys []string
	err = global.DB.Table(table.Privilege).Select("key").Find(&adminPrivilegeKeys).Error
	if err != nil {
		panic(fmt.Errorf("privilege query failed: (%s)", err.Error()))
	}
	adminPrivilegeKeysByte, err := json.Marshal(adminPrivilegeKeys)
	err = global.DB.Table(table.Role).Create(&role.Role{
		Name:          "管理员",
		IsSystem:      1,
		PrivilegeKeys: adminPrivilegeKeysByte,
	}).Error
	if err != nil {
		panic(fmt.Errorf("create admin role failed: (%s)", err.Error()))
	}
	// 普通角色
	plainPrivilegeKeys, err := json.Marshal([]string{"operate_log_list", "operate_log_download", "role_list", "role_detail",
		"privilege_list", "account_list", "account_detail_other"})
	err = global.DB.Table(table.Role).Create(&role.Role{
		Name:          "普通用户",
		IsSystem:      1,
		PrivilegeKeys: plainPrivilegeKeys,
	}).Error
	if err != nil {
		panic(fmt.Errorf("create plain role failed: (%s)", err.Error()))
	}

}

func DefaultUserInit() {
	err := global.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table.User)).Error
	if err != nil {
		panic(fmt.Errorf("clear user table failed: (%s)", err.Error()))
	}
	secretKey, err := utils.GenerateSecretKey()
	if err != nil {
		panic(fmt.Errorf("generate seecret key failed: (%s)", err.Error()))
	}
	name := "管理员"
	fullName, shortName := common.ConvertCnToLetter(name)
	err = global.DB.Table(table.User).Create(&user.User{
		Id:        utils.GenID(),
		Account:   "root",
		Password:  utils.BcryptHash("123456"),
		SecretKey: secretKey,
		Name:      name,
		FullName:  fullName,
		ShortName: shortName,
		IsSystem:  true,
	}).Error
	if err != nil {
		panic(fmt.Errorf("create user failed: (%s)", err.Error()))
	}
}

func DefaultUserRoleInit() {
	err := global.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table.UserRole)).Error
	if err != nil {
		panic(fmt.Errorf("clear user_role table failed: (%s)", err.Error()))
	}
	err = global.DB.Table(table.UserRole).Create(&user.UserRole{
		UserId: 1,
		RoleId: 1,
	}).Error
	if err != nil {
		panic(fmt.Errorf("create user role failed: (%s)", err.Error()))
	}
}
