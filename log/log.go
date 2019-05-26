package log

import "github.com/sirupsen/logrus"

func Error(err error) {
	if err != nil {
		logrus.Error(err)
	}
}
