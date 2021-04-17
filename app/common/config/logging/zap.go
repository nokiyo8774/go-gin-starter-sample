package logging

import (
	"io"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup(options ...zap.Option) *zap.Logger {
	hook1 := getWriter("/var/log/app/access.log")
	w := zapcore.AddSync(hook1)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core)

	return logger
}

// func Setup(options ...zap.Option) (*zap.Logger, error) {
// 	return zap.Config{
// 		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
// 		Development: false,
// 		Sampling: &zap.SamplingConfig{
// 			Initial:    100,
// 			Thereafter: 100,
// 		},
// 		Encoding:         "json",
// 		EncoderConfig:    encoderConfig(),
// 		OutputPaths:      []string{"stdout", "/var/log/app/access.log"},
// 		ErrorOutputPaths: []string{"stderr"},
// 	}.Build(options...)
// }

func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
