// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllStockSales(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks sales")
	result, err := s.db.GetAllStocksSales()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getStockSaleByID godoc
// @Summary Get stocks sale by ID
// @Description get stocks sale data
// @Accept json
// @Produce json
// @Router /stocks/sales/{id} [get]
// @Success 200 {object} wallet.Stock
// @Param id path string true "Sale id"
func (s *server) getStockSaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock sale with id: %s", id)
	result, err := s.db.GetStockSaleByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Stock sale data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertStockSale godoc
// @Summary Insert some stocks sale
// @Description insert new stocks sale
// @Accept json
// @Produce json
// @Router /stocks/sales [post]
func (s *server) insertStockSale(c echo.Context) error {
	log.Debugf("[API] Inserting stock sale")

	data := wallet.NewStock()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertStockSale(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateStockSaleByID godoc
// @Summary Update some stocks sale
// @Description update new stocks sale
// @Accept json
// @Produce json
// @Router /stocks/sales/{id} [put]
// @Param id path string true "Sale id"
func (s *server) updateStockSaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock sale with id %s", id)

	data := wallet.NewStock()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateStockSaleByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Stock sale not found")
}
