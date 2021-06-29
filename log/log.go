package log

import "github.com/sirupsen/logrus"

func Fatal(v ...interface{}) {
	logrus.Fatal(v...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func SetLogLevel(level int) {
	logrus.SetLevel(logrus.Level(level))
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}
