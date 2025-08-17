package log

import (
	"oa/internal/config"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger // 全局日志实例
	once   sync.Once   // 确保日志实例只被初始化一次
)

func Init(env string) {
	once.Do(func() {
		// 设置日志编码器配置
		encCfg := zapcore.EncoderConfig{
			TimeKey:       "ts",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			StacktraceKey: "stack",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime: func(t time.Time, arr zapcore.PrimitiveArrayEncoder) {
				arr.AppendString(t.Format("2006-01-02 15:04:05"))
			},
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		var encoder zapcore.Encoder
		if env == "debug" {
			encoder = zapcore.NewConsoleEncoder(encCfg)
		} else {
			encoder = zapcore.NewJSONEncoder(encCfg)
		}
		lv := zap.InfoLevel
		_ = lv.Set("info")
		core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lv)
		opts := []zap.Option{zap.AddCaller()}
		if env == "debug" {
			opts = append(opts, zap.Development())
		}
		logger = zap.New(core, opts...)
	})
}

func L() *zap.Logger {
	if logger == nil {
		// 如果没有初始化，使用默认配置
		Init(config.AppConfig.Server.Mode)
	}
	return logger
}

// Debug 打印调试级别日志
func Debug(msg string, fields ...zap.Field) {
	L().Debug(msg, fields...)
}

// Info 打印信息级别日志
func Info(msg string, fields ...zap.Field) {
	L().Info(msg, fields...)
}

// Warn 打印警告级别日志
func Warn(msg string, fields ...zap.Field) {
	L().Warn(msg, fields...)
}

// Error 打印错误级别日志
func Error(msg string, fields ...zap.Field) {
	L().Error(msg, fields...)
}

// Panic 打印恐慌级别日志并触发 panic
func Panic(msg string, fields ...zap.Field) {
	L().Panic(msg, fields...)
}

// Fatal 打印致命错误级别日志并终止程序
func Fatal(msg string, fields ...zap.Field) {
	L().Fatal(msg, fields...)
}
