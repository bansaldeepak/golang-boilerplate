// package logger

// import (
// 	"go.uber.org/zap"
// )

// func NewLogger() (*zap.Logger, *zap.SugaredLogger) {
// 	logger, _ := zap.NewProduction()
// 	sugar := logger.Sugar()
// 	return logger, sugar
// }

package logger

import (
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func init() {
	Logger, _ = zap.NewProduction()
	Sugar = Logger.Sugar()
}

func NewLogger() (*zap.Logger, *zap.SugaredLogger) {
	return Logger, Sugar
}
