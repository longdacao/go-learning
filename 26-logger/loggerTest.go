package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	var logger *zap.Logger
	var writes = []zapcore.WriteSyncer{}
	writes = append(writes, zapcore.AddSync(os.Stdout))

	// Setup Log Level
	atomicLevel := zap.NewAtomicLevel()

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	caller := zap.AddCaller()
	field := zap.Fields(zap.String("appName", "jolycao"))
	development := zap.Development()
	logger = zap.New(core, caller, development, field)

	logger.Info("log initialized successfully")

	// logFileLocation, _ := os.OpenFile("/Users/jolycao/temp/go-learning/26-logger/test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	// logger.SetOutput(logFileLocation)
	// logger.Printf("Error fetching url ")
}
