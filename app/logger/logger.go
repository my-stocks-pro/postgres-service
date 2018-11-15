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
	InitLogger() TypeLogger
}

type TypeLogger struct {
	Log *zap.Logger
}

func (l TypeLogger) InitLogger() TypeLogger {
	filename := fmt.Sprintf("%s/%s_%s_%s.log", logPath, logPrefix, logFile, time.Now().Format("2006-01-02_15:04"))
	_, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Creation of log file errored!")
		os.Exit(1)
	}

	conf := zap.NewDevelopmentConfig()
	conf.OutputPaths = []string{
		filename,
	}

	zapLog, err := conf.Build()
	if err != nil {
		fmt.Println(err)
	}

	defer zapLog.Sync()

	zapLog.Info("Logger successful initialization")

	return TypeLogger{Log: zapLog}
}

func NewLogger() TypeLogger {
	return TypeLogger{}
}
