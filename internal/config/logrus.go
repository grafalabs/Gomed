package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(config *viper.Viper) *logrus.Logger {

    logLevel := logrus.Level(config.GetInt32("log.level"))
    log := logrus.New() 

    log.SetLevel(logLevel)
    log.SetFormatter(&logrus.JSONFormatter{})

    return log 
}

