// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func errorMessage(m string) ErrorMessage {
	return ErrorMessage{Message: m}
}

func logAndReturnError(c echo.Context, m string) error {
	log.Error(fmt.Sprintf("[API] %s", m))
	return c.JSON(http.StatusInternalServerError, errorMessage(m))
}
