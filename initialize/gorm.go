package initialize

import (
	"base-system-backend/global"
	"base-system-backend/model/privilege"
	"base-system-backend/model/role"
	"base-system-backend/model/user"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func GormPgSql() *gorm.DB {
	p := global.CONFIG.Pgsql
	if p.Dbname == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}

	if db, err := gorm.Open(postgres.New(pgsqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		//Logger:                                   logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		return db
	}
}

func CloseDB() {
	db, err := global.DB.DB()
	if err != nil {
		return
	}
	_ = db.Close()
}

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		// 用户表
		&user.User{},
		// 用户角色表
		&user.UserRole{},
		// 角色表
		&role.Role{},
		// 权限表
		&privilege.Privilege{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
