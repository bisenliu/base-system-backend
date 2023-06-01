package internal

import (
	"base-system-backend/enums/table"
	"base-system-backend/global"
	"base-system-backend/model/privilege"
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
	bytePrivliege, err := os.ReadFile(privilegeJsonPath)
	if err != nil {
		panic(fmt.Errorf("read privilege json file failed: (%s)", err.Error()))
	}
	var Privileges struct {
		RECORDS []map[string]string `json:"records"`
	}
	err = json.Unmarshal(bytePrivliege, &Privileges)
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

func defaultRoleInit() {
	global.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table.Role))
}

func defaultUserInit() {
	global.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table.User))
}

func defaultUserRoleInit() {
	global.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table.UserRole))
}
