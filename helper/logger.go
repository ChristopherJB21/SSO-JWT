package helper

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile(viper.GetString("appName") + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	return logger
}
