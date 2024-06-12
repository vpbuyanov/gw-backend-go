package logger

import (
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/gw-backend-go/internal/configs"
)

var Log = logrus.New()

func InitLogger(cfg configs.Logger) {
	Log.SetFormatter(&logrus.JSONFormatter{})

	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}

	err := Log.Level.UnmarshalText([]byte(cfg.LogLevel))
	if err != nil {
		Log.Panicf("failed to set log level: %v", err)
	}

	Log.Infof("logger level set to %v", cfg.LogLevel)
}
