// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllStockFundsSales(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks funds sales")
	result, err := s.db.GetAllStocksFundsSales()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getStockFundSaleByID godoc
// @Summary Get stocks fund sale by ID
// @Description get stocks fund sale data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.StockFund
// @Router /stocks-funds/sales/{id} [get]
// @Param id path string true "Sale id"
func (s *server) getStockFundSaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock fund sale with id: %s", id)
	result, err := s.db.GetStockFundSaleByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Stock fund sale data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertStockFundSale godoc
// @Summary Insert some stocks fund sale
// @Description insert new stocks fund sale
// @Accept json
// @Produce json
// @Router /stocks-funds/sales [post]
func (s *server) insertStockFundSale(c echo.Context) error {
	log.Debugf("[API] Inserting stock fund sale")

	data := wallet.NewStockFund()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertStockFundSale(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateStockFundSaleByID godoc
// @Summary Update some stocks fund sale
// @Description update new stocks fund sale
// @Accept json
// @Produce json
// @Router /stocks-funds/sales/{id} [put]
// @Param id path string true "Sale id"
func (s *server) updateStockFundSaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock fund sale with id %s", id)

	data := wallet.NewStockFund()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateStockFundSaleByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Stock fund sale not found")
}
