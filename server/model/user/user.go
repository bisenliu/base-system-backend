package user

import (
	"base-system-backend/constants/gender"
	"base-system-backend/constants/login"
	"base-system-backend/constants/table"
	"base-system-backend/constants/user"
	"base-system-backend/model/common/field"
	"base-system-backend/model/role"
)

type User struct {
	Id          int64            `gorm:"column:id;primaryKey;autoIncrement;notNull;comment:Id"`
	Account     string           `gorm:"column:account;notNull;unique;size:20;comment:账号"`
	Password    string           `gorm:"column:password;notNull;size:70;comment:密码"`
	SecretKey   string           `gorm:"column:secret_key;notNull;size:64;comment:API秘钥"`
	Phone       string           `gorm:"column:phone;size:11;comment:手机号"`
	Email       string           `gorm:"column:email;size:100;comment:邮箱"`
	Name        string           `gorm:"column:name;notNull;size:20;comment:姓名"`
	FullName    string           `gorm:"column:full_name;size:50;comment:姓名全拼"`
	ShortName   string           `gorm:"column:short_name;size:50;comment:姓名简拼"`
	IdCard      string           `gorm:"column:id_card;size:18;comment:身份证号码"`
	Avatar      string           `gorm:"column:avatar;size:100;comment:头像"`
	CurrentIp   string           `gorm:"column:current_ip;size:50;comment:当前登陆Ip"`
	LastIp      string           `gorm:"column:last_ip;size:50;comment:最后登陆Ip"`
	IsSystem    bool             `gorm:"column:is_system;default:false;comment:是否系统账号"`
	Gender      gender.Gender    `gorm:"column:gender;notNull;comment:性别"`
	Status      user.AccStatus   `gorm:"column:status;default:0;comment:账号状态"`
	LoginType   login.Type       `gorm:"column:login_type;default:0;comment:登陆方式"`
	CurrentTime field.CustomTime `gorm:"column:current_time;comment:当前登陆时间"`
	LastTime    field.CustomTime `gorm:"column:last_time;comment:最后登出时间"`
	field.CUTime
}

func (receiver User) TableName() string {
	return table.User
}

type UserRole struct {
	UserId int64 `gorm:"column:user_id;notNull;comment:用户Id"`
	RoleId int64 `gorm:"column:role_id;notNull;comment:角色Id"`
	User   User
	Role   role.Role
}

func (receiver UserRole) TableName() string {
	return table.UserRole
}

type BlackList struct {
	Id         int64            `gorm:"column:id;primaryKey;autoIncrement;notNull;comment:Id"`
	FailedNum  int              `gorm:"column:failed_num;notNull;comment:登陆失败次数"`
	Account    string           `gorm:"column:account;notNull;unique;size:20;comment:账号"`
	FailedTime field.CustomTime `gorm:"column:failed_time;autoCreateTime:milli;comment:登陆失败时间"`
	NextTime   field.CustomTime `gorm:"column:next_time;comment:下次登陆时间"`
}

func (receiver BlackList) TableName() string {
	return table.UserBlackList
}
