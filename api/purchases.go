// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// getAllPurchases godoc
// @Summary List all purchases
// @Description get all purchases data
// @Accept json
// @Produce json
// @Router /purchases [get]
func (s *server) getAllPurchases(c echo.Context) error {
	log.Debug("[API] Retrieving all purchases")
	result, err := s.db.GetAllPurchases()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// deletePurchaseByID godoc
// @Summary Delete purchase by ID
// @Description delete some purchase by id
// @Accept json
// @Produce json
// @Router /purchases/{id} [delete]
// @Param id path string true "Sale id"
func (s *server) deletePurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Deleting %s data", id)
	result, err := s.db.DeletePurchaseByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
