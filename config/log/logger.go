package log

import (
	"fmt"
	"github.com/aisuosuo/letter/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Logger *zap.SugaredLogger

func init() {
	filename := fmt.Sprintf("%s/logs/letter.log", config.AppPath)
	fileRotate := &lumberjack.Logger{
		Filename:   filename,
		MaxBackups: 7,
	}
	var level zapcore.Level
	switch config.GlobalConfig.Log.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	case "warn":
		level = zap.WarnLevel
	default:
		level = zap.InfoLevel
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:       "T",
			LevelKey:      "L",
			NameKey:       "N",
			CallerKey:     "C",
			MessageKey:    "M",
			StacktraceKey: "S",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
			},
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		}),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileRotate)),
		&atomicLevel,
	)
	Logger = zap.New(core, zap.AddCaller()).Sugar()
	Logger.Info("logger init success")
}
