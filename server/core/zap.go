package core

import (
	"base-system-backend/global"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// Zap
//
//	@Description: 初始化Logger
//	@return logger 全局logger对象
func Zap() (logger *zap.Logger) {
	// 日志文件配置 文件位置和切割
	writeSyncer := getLogWriter(global.CONFIG.Zap.Director, global.CONFIG.Zap.MaxSize, global.CONFIG.Zap.MaxBackups, global.CONFIG.Zap.MaxAge)
	// 获取日志输出编码
	encoder := getEncoder()
	var l = new(zapcore.Level)
	if err := l.UnmarshalText([]byte(global.CONFIG.Zap.Level)); err != nil {
		panic(fmt.Errorf("zap l.UnmarshalText failed: %w", err))
	}
	var core zapcore.Core
	if global.ENV == "local" {
		//进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}
	// zap.Addcaller() 输出日志打印文件和行数如： logger/logger_test.go:33
	// 1. zap.ReplaceGlobals 函数将当前初始化的 logger 替换到全局的 logger,
	// 2. 使用 logger 的时候 直接通过 zap.S().Debugf("xxx") or zap.L().Debug("xxx")
	// 3. 使用 zap.S() 和 zap.L() 提供全局锁，保证一个全局的安全访问logger的方式
	logger = zap.New(core, zap.AddCaller())
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logger)
	zap.L().Info("init zap success")
	return
}

// getEncoder
//  @Description: 编码器(如何写入日志)
//  @return zapcore.Encoder

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// log 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	// 输出level序列化为全大写字符串，如 INFO DEBUG ERROR
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriter
//  @Description: 获取日志输出方式  日志文件 控制台
//  @param filename 日志文件路径
//  @param maxSize 单个日志文件最大多少 MB
//  @param maxBackup 日志备份数量
//  @param maxAge 日志最长保留时间
//  @return zapcore.WriteSyncer

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	// 日志只输出到日志文件
	return zapcore.AddSync(lumberJackLogger)
}
