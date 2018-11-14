package logger

type Logger interface {
	InitLogger() TypeLogger
}

type TypeLogger struct {

}


func (l TypeLogger) InitLogger() TypeLogger {
	return TypeLogger{}
}

func NewLogger() TypeLogger {
	return TypeLogger{}
}