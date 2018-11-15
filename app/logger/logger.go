package logger

import (
	"fmt"
	"time"
	"os"
	"go.uber.org/zap"
)

const (
	logPath   = "app_logs"
	logFile   = "logs"
	logPrefix = "postgres-service"
)

type Logger interface {
	InitLogger() (TypeLogger, error)
}

type TypeLogger struct {
	Log *zap.Logger
}

func (l TypeLogger) InitLogger() (TypeLogger, error) {
	filename := fmt.Sprintf("%s/%s_%s_%s.log", logPath, logPrefix, logFile, time.Now().Format("2006-01-02_15:04"))
	_, err := os.Create(filename)
	if err != nil {
		return TypeLogger{}, err
	}

	conf := zap.NewDevelopmentConfig()
	conf.OutputPaths = []string{
		filename,
	}

	zapLog, err := conf.Build()
	if err != nil {
		return TypeLogger{}, err
	}

	defer zapLog.Sync()

	return TypeLogger{Log: zapLog}, nil
}

func NewLogger() TypeLogger {
	return TypeLogger{}
}
