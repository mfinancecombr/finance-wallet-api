// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/mfinancecombr/finance-wallet-api/api"
	_ "github.com/mfinancecombr/finance-wallet-api/config"
)

func main() {
	server, err := api.NewServerFromDB()
	if err != nil {
		log.Error(err)
	}
	server.Start()
}
