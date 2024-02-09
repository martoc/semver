package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

var file *os.File

const readWrite = 0o666

func init() {
	file, _ = os.OpenFile("semver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, readWrite)

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	log.SetOutput(file)
}

func GetInstance() *logrus.Logger {
	return log
}

func Close() {
	file.Close()
}
