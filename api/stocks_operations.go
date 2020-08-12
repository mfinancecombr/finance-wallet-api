// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllStockOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks operations")
	result, err := s.db.GetAllStocksOperations()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getStockOperationByID godoc
// @Summary Get stocks operation by ID
// @Description get stocks operation data
// @Accept json
// @Produce json
// @Router /stocks/operations/{id} [get]
// @Success 200 {object} wallet.Stock
// @Param id path string true "Operation id"
func (s *server) getStockOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock operation with id: %s", id)
	result, err := s.db.GetStockOperationByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Stock operation data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertStockOperation godoc
// @Summary Insert some stocks operation
// @Description insert new stocks operation
// @Accept json
// @Produce json
// @Router /stocks/operations [post]
func (s *server) insertStockOperation(c echo.Context) error {
	log.Debugf("[API] Inserting stock operation")

	data := wallet.NewStock()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertStockOperation(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateStockOperationByID godoc
// @Summary Update some stocks operation
// @Description update new stocks operation
// @Accept json
// @Produce json
// @Router /stocks/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateStockOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock operation with id %s", id)

	data := wallet.NewStock()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateStockOperationByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Stock operation not found")
}
