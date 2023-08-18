package global

import (
	"base-system-backend/config"
	"embed"
	"github.com/TestsLing/aj-captcha-go/service"
	sf "github.com/bwmarrin/snowflake"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ENV        string
	SystemInit bool
	CONFIG     config.Service
	LOG        *zap.Logger
	DB         *gorm.DB
	VP         *viper.Viper
	REDIS      *redis.Client
	TRANS      ut.Translator
	Node       *sf.Node
	FS         embed.FS
	CAPTCHA    *service.CaptchaServiceFactory
)
