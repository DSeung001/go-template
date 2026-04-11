package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	// 후에 로그 수집기로 수집 후 로그 서버로 전달
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
}
