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

type HttpResponse struct {
	Url      string
	Response *http.Response
	Err      error
}

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

func GetAsync(urls []string) map[string]*HttpResponse {
	ch := make(chan *HttpResponse, len(urls))
	responses := map[string]*HttpResponse{}
	for _, path := range urls {
		go func(path string) {
			log.Debugf("[FinanceAPI] Fetching %s", path)
			url := viper.GetString("financeapi.url") + path
			resp, err := http.Get(url)
			ch <- &HttpResponse{Url: path, Response: resp, Err: err}
		}(path)
	}

	for {
		select {
		case r := <-ch:
			log.Debugf("[FinanceAPI] %s was fetched", r.Url)
			responses[r.Url] = r
			if len(responses) == len(urls) {
				return responses
			}
		default:
			time.Sleep(1 * time.Millisecond)
		}
	}

	return responses
}
