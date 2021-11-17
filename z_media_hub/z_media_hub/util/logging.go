package util

import (
    "log"
    "os"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func initLogger() {
		log.Println("Setting up utility loggers...")
    file, err := os.OpenFile("main.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(data string){
	InfoLogger.Println(data)
}

func LogWarning(data string){
	WarningLogger.Println(data)
}

func LogError(data string) {
	ErrorLogger.Println(data)
}