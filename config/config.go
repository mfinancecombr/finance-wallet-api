// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	envReplacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(envReplacer)
	viper.SetEnvPrefix("finance.walletapi")
	viper.SetDefault("mongodb.endpoint", "mongodb://localhost:27017")
	viper.SetDefault("mongodb.name", "finance-wallet")
	viper.SetDefault("port", 8889)
	viper.SetDefault("debug", false)
	logLevel := log.InfoLevel
	if viper.GetBool("debug") {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	viper.SetDefault("db.operation.timeout", 3)
	viper.SetDefault("collection.operation.timeout", 3)
	viper.SetDefault("financeapi.operation.timeout", 3)
	viper.SetDefault("financeapi.url", "https://mfinance.com.br/api/v1")
}
