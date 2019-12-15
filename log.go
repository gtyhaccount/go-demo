package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	logInit()

	log.WithFields(log.Fields{
		"id":   "id",
		"age":  23,
		"name": "lee",
	}).Debug("123456")

	requestLogger := log.WithFields(log.Fields{"request_id": "request_id", "user_ip": "user_ip"})
	requestLogger.WithFields(log.Fields{
		"id":   "id",
		"age":  23,
		"name": "lee",
	}).Info("something happened on that request") // will log request_id and user_ip
	requestLogger.Error("something not great happened")

}

func logInit() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
