package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const API_VERSION = "v1"

var log = logrus.New()

type Error interface {
	Error() string
}

type LogDesc struct {
	LogId     string    `json:"log_id"`
	EventDate time.Time `json:"event_date"`
}

func Log() *logrus.Logger {
	log.Out = os.Stdout

	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(dir+"/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}

func (e *LogDesc) Error() string {
	return fmt.Sprintf("Erro ocurrent: id %s", e.LogId)
}
