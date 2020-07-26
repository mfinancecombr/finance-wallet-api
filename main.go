// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/mfinancecombr/finance-wallet-api/api"
	_ "github.com/mfinancecombr/finance-wallet-api/config"
)

// @title MFinance Wallet API
// @version 0.1.0
// @description mfinance Wallet API data.

// @contact.name API Support
// @contact.url https://github.com/mfinancecombr/finance-wallet-api

// @license.name BSD 3-Clause
// @license.url https://opensource.org/licenses/BSD-3-Clause

// @host localhost:8889
// @BasePath /api/v1
func main() {
	server, err := api.NewServerFromDB()
	if err != nil {
		log.Error(err)
	}
	server.Start()
}
