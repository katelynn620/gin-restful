package util

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DEFAULT_LOG_LEVEL = "info"
)

var (
	logger *zap.Logger
)

func getLogEncoder() zapcore.Encoder {
	// Get a specified EncoderConfig for customization
	encodeConfig := zap.NewDevelopmentEncoderConfig()
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encodeConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func InitLogger(debug bool) (err error) {
	var zapcores []zapcore.Core

	level := new(zapcore.Level)
	logLevel := DEFAULT_LOG_LEVEL
	if debug {
		logLevel = "debug"
	}
	err = level.UnmarshalText([]byte(logLevel))
	if err != nil {
		panic(fmt.Sprintf("unmarshal log level error: %v", err))
	}

	encoder := getLogEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, level)

	zapcores = append(zapcores, core)
	logger = zap.New(zapcore.NewTee(zapcores...), zap.AddCaller())

	// replace zap's global logger instance, and then you can call zap.L() in other packages
	zap.ReplaceGlobals(logger)
	return
}
