package pkg

import (
	"errors"
	"io"
	"log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	logOnce sync.Once
	logger  *Log
)

type Log struct {
	*logrus.Logger
}

func NewLog() *Log {
	return &Log{}
}

func GetLog() *Log {
	logOnce.Do(func() {
		logger = NewLog()
	})
	return logger
}

func (l *Log) Initialize(level string) {
	var err error
	l.Logger, err = l.stdoutInit(level)
	if err != nil {
		log.Panic(err)
	}
}

func (Log) stdoutInit(lvl string) (*logrus.Logger, error) {
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
