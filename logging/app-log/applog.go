package applog

import (
	"encoding/json"
	"log"
	"time"
)

type AppLogger struct {
	appName string
}

func NewAppLogger(appName string) *AppLogger {
	return &AppLogger{appName: appName}
}

func (al *AppLogger) Info(payload string) {
	logLevel := "INFO"
	logType := "APP"

	entry := newAppLog(logLevel, logType, payload, al.appName)
	jsonEntry, err := json.Marshal(entry)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(jsonEntry))
}

type appLog struct {
	Timestamp string `json:"timestamp"`
	LogLevel  string `json:"log_level"`
	LogType   string `json:"log_type"`
	Payload   string `json:"payload"`
	AppName   string `json:"app_name"`
}

func newAppLog(logLevel string, logType string, payload string, appName string) *appLog {
	currentTime := time.Now()
	return &appLog{
		Timestamp: currentTime.Format("2006-01-02T15:04:05.000-07:00"),
		LogLevel:  logLevel,
		LogType:   logType,
		Payload:   payload,
		AppName:   appName,
	}
}
