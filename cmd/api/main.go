package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/gw-backend-go/configs"
	"github.com/vpbuyanov/gw-backend-go/internal/app"
)

func main() {
	ctx := context.Background()
	config := configs.LoadConfig()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	err := logger.Level.UnmarshalText([]byte(logLevel))
	if err != nil {
		logger.Panicf("failed to set log level: %v", err)
	}

	logger.Infof("logger level set to %v", logLevel)

	application := app.New(logger, config)
	application.Run(ctx)
}
