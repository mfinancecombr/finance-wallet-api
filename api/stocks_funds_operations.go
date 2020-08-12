// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllStockFundsOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks funds operations")
	result, err := s.db.GetAllStocksFundsOperations()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getStockFundOperationByID godoc
// @Summary Get stocks fund operation by ID
// @Description get stocks fund operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.StockFund
// @Router /stocks-funds/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getStockFundOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock fund operation with id: %s", id)
	result, err := s.db.GetStockFundOperationByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Stock fund operation data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertStockFundOperation godoc
// @Summary Insert some stocks fund operation
// @Description insert new stocks fund operation
// @Accept json
// @Produce json
// @Router /stocks-funds/operations [post]
func (s *server) insertStockFundOperation(c echo.Context) error {
	log.Debugf("[API] Inserting stock fund operation")

	data := wallet.NewStockFund()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertStockFundOperation(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateStockFundOperationByID godoc
// @Summary Update some stocks fund operation
// @Description update new stocks fund operation
// @Accept json
// @Produce json
// @Router /stocks-funds/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateStockFundOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock fund operation with id %s", id)

	data := wallet.NewStockFund()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateStockFundOperationByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Stock fund operation not found")
}
