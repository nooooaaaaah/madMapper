package config

import (
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("log/madMapper.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(logFile)
}

func LogInfo(message string, optional ...interface{}) {
	if len(optional) > 0 {
		log.Println("INFO:", message, optional)
	} else {
		log.Println("INFO:", message)
	}
}

func LogError(message string, err error) {
	log.Printf("ERROR: %s - %v\n", message, err)
}
