// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllStockPurchases(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks purchases")
	result, err := s.db.GetAllStocksPurchases()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getStockPurchaseByID godoc
// @Summary Get stocks purchase by ID
// @Description get stocks purchase data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.Stock
// @Router /stocks/purchases/{id} [get]
// @Param id path string true "Purchase id"
func (s *server) getStockPurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock purchase with id: %s", id)
	result, err := s.db.GetStockPurchaseByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Stock purchase data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertStockPurchase godoc
// @Summary Insert some stocks purchase
// @Description insert new stocksFII purchase
// @Accept json
// @Produce json
// @Router /stocks/purchases [post]
func (s *server) insertStockPurchase(c echo.Context) error {
	log.Debugf("[API] Inserting stock purchase")

	data := wallet.NewStock()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertStockPurchase(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateStockPurchaseByID godoc
// @Summary Update some stocks purchase
// @Description update new stocksFII purchase
// @Accept json
// @Produce json
// @Router /stocks/purchases/{id} [put]
// @Param id path string true "Purchase id"
func (s *server) updateStockPurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock purchase with id %s", id)

	data := wallet.NewStock()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateStockPurchaseByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Stock purchase not found")
}
