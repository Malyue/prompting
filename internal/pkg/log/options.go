package log

import "go.uber.org/zap/zapcore"

type Options struct {
	// 是否开启 caller ，如果开启会在日志中显示调用日志所在的文件和行号
	DisableCaller bool
	// 是否禁止在panic及以上级别打印堆栈信息
	DisableStacktrace bool
	// 指定日志级别
	Level string
	// 指定日志显示格式 console/json
	Format string
	// 指定日志输出位置
	OutputPaths []string
}

// NewOptions
func NewOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}
