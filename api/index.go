// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) index(c echo.Context) error {
	data := map[string]string{
		"version": Version,
	}
	return c.JSON(http.StatusOK, data)
}
