// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package financeapi

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var financeClient = &http.Client{
	Timeout: viper.GetDuration("financeapi.operation.timeout") * time.Second,
}

func GetJSON(path string, target interface{}) error {
	log.Debugf("[FinanceAPI] Retrieving %s", path)
	url := viper.GetString("financeapi.url") + path
	r, err := financeClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
