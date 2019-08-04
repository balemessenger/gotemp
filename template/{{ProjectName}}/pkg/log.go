package pkg

import (
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

type Logger struct {
	*logrus.Logger
}

func NewLog(level string) *Logger {
	l, err := stdoutInit(level)
	if err != nil {
		log.Panic(err)
	}
	return &Logger{l}
}

func stdoutInit(lvl string) (*logrus.Logger, error) {
	var err error
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		err = errors.New("failed to parse level")
		return nil, err
	}
	logger.Level = level
	var logWriter io.Writer = os.Stdout
	logger.SetOutput(logWriter)

	return logger, err
}
