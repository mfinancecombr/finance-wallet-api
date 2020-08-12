// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// getAllOperations godoc
// @Summary List all operations
// @Description get all operations data
// @Accept json
// @Produce json
// @Router /operations [get]
func (s *server) getAllOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all operations")
	result, err := s.db.GetAllOperations()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// deleteOperationByID godoc
// @Summary Delete operation by ID
// @Description delete some operation by id
// @Accept json
// @Produce json
// @Router /operations/{id} [delete]
// @Param id path string true "Operation id"
func (s *server) deleteOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Deleting %s data", id)
	result, err := s.db.DeleteOperationByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
