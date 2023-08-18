package initialize

import (
	"base-system-backend/global"
	"fmt"
	config2 "github.com/TestsLing/aj-captcha-go/config"
	"github.com/TestsLing/aj-captcha-go/const"
	"github.com/TestsLing/aj-captcha-go/service"
	"image/color"
)

func Captcha() *service.CaptchaServiceFactory {
	// **********************默认配置***************************************************
	// 默认配置，可以根据项目自行配置，将其他类型配置序列化上去
	//var config = config2.NewConfig()

	// *********************自定义配置**************************************************
	// 水印配置（参数可从业务系统自定义）
	watermarkConfig := &config2.WatermarkConfig{
		FontSize: 12,
		Color:    color.RGBA{R: 255, G: 255, B: 255, A: 255},
		Text:     "我的水印",
	}

	// 点击文字配置（参数可从业务系统自定义）
	clickWordConfig := &config2.ClickWordConfig{
		FontSize: 25,
		FontNum:  4,
	}

	// 滑动模块配置（参数可从业务系统自定义）
	blockPuzzleConfig := &config2.BlockPuzzleConfig{Offset: 10}

	// 行为校验配置模块（具体参数可从业务系统配置文件自定义）
	config := config2.BuildConfig(constant.RedisCacheKey, constant.DefaultResourceRoot, watermarkConfig,
		clickWordConfig, blockPuzzleConfig, 60)

	// 服务工厂，主要用户注册 获取 缓存和验证服务
	factory := service.NewCaptchaServiceFactory(config)

	// 内存缓存，但是不足以应对生产环境，希望自行注册缓存驱动 实现缓存接口即可替换（CacheType就是注册进去的 key）
	// factory.RegisterCache(constant.MemCacheKey, service.NewMemCacheService(20)) // 这里20指的是缓存阈值
	// 注册使用默认redis数据库
	// factory.RegisterCache(constant.RedisCacheKey, service.NewDftRedisCacheService())
	// 注册自定义配置redis数据库
	factory.RegisterCache(constant.RedisCacheKey, service.NewConfigRedisCacheService(
		[]string{fmt.Sprintf("%s:%s", global.CONFIG.Redis.Host, global.CONFIG.Redis.Port)},
		"",
		global.CONFIG.Redis.Password,
		false,
		global.CONFIG.Redis.VerifyCodDb))

	// 注册了两种验证码服务 可以自行实现更多的验证
	factory.RegisterService(constant.ClickWordCaptcha, service.NewClickWordCaptchaService(factory))
	factory.RegisterService(constant.BlockPuzzleCaptcha, service.NewBlockPuzzleCaptchaService(factory))

	return factory
}
