package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

func New(logLevel string) (*zap.Logger, func()) {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	level, err := zapcore.ParseLevel(logLevel)

	if err != nil {
		log.Fatal("failed to parse log level", err)
	}

	loggerConfig.Level.SetLevel(level)

	logger, err := loggerConfig.Build()

	if err != nil {
		log.Fatal("failed to initialize logger", err)
	}

	return logger, func() {
		_ = logger.Sync()
	}
}
