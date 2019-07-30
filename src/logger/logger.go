package logger

import (
	"config"
	"fmt"
	"github.com/kmiku7/golog"
	"github.com/kmiku7/logrus-bridge/formatter"
	"github.com/kmiku7/logrus-bridge/hooks"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

//import "github.com/kmiku7/golog"
//
//func main() {
//  fileBackend, err := golog.NewFileBackend("./log")
//  if err != nil {
//    panic(err)
//  }
//  defer fileBackend.Close()
//  fileBackend.SetRotateFile(true, 24)
//
//  fileBackend.Log(golog.Info, "Hello World!")
//}

type fileBackendWrapper struct {
	fileBackend *golog.FileBackend
}

func (f fileBackendWrapper) Log(level logrus.Level, message []byte) error {
	levelMap := map[logrus.Level]golog.Level{
		logrus.DebugLevel: golog.Debug,
		logrus.InfoLevel:  golog.Info,
		logrus.WarnLevel:  golog.Warning,
		logrus.ErrorLevel: golog.Error,
		logrus.FatalLevel: golog.Fatal,
	}

	gologLevel, has := levelMap[level]
	if !has {
		return fmt.Errorf("unknown level: %v", level)
	}
	f.fileBackend.Log(gologLevel, message)
	return nil
}

func OpenLogger(logConfig *config.LogConfig) (*logrus.Logger, func(), error) {
	fileBackend, err := golog.NewFileBackend(logConfig.LogDir)
	if err != nil {
		return nil, nil, err
	}
	wrapper := fileBackendWrapper{
		fileBackend: fileBackend,
	}

	logClient := logrus.New()
	logClient.Out = ioutil.Discard
	logClient.Formatter = formatter.EmptyFormatter(0)
	logClient.SetLevel(logrus.DebugLevel)

	hook := hooks.NewBackendHook(wrapper, &logrus.TextFormatter{}, logrus.AllLevels)
	logClient.AddHook(hook)

	return logClient, func() {
		fileBackend.Close()
	}, nil
}
