package log

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})
}

type myFormatter struct {
	module    string
	formatter logrus.TextFormatter
}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 31 // gray
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}

	msgText := fmt.Sprintf("[%s] - [%s] - %s - %s\n", entry.Time.Format(f.formatter.TimestampFormat), f.module, strings.ToUpper(entry.Level.String()), entry.Message)

	return []byte(fmt.Sprintf("\x1b[%dm%s\x1b[0m", levelColor, msgText)), nil
}

func GetLogger(module string) Logger {
	l := logrus.New()
	l.SetFormatter(&myFormatter{
		module: module,
		formatter: logrus.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05",
			ForceColors:            true,
			DisableLevelTruncation: true,
		},
	})
	l.SetReportCaller(false)

	return l
}
