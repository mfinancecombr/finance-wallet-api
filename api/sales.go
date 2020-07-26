// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// getAllSales godoc
// @Summary List all sales
// @Description get all sales data
// @Accept json
// @Produce json
// @Router /sales [get]
func (s *server) getAllSales(c echo.Context) error {
	log.Debug("[API] Retrieving all sales")
	result, err := s.db.GetAllSales()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// deleteSaleByID godoc
// @Summary Delete sale by ID
// @Description delete some sale by id
// @Accept json
// @Produce json
// @Router /sales/{id} [delete]
// @Param id path string true "Sale id"
func (s *server) deleteSaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Deleting %s data", id)
	result, err := s.db.DeleteSaleByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
