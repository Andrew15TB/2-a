package main

import (
	"log"

	"go.uber.org/zap"
)

func getLoggerHandler(pack string) *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("[%s] Failed to initialize logger: %v", pack, err)
	}
	rLogger := logger.Named(pack)
	rLogger.Info("Get Logger Handler!")
	return rLogger
}

func init() {
	LH := getLoggerHandler("Logger")
	defer LH.Sync()
}