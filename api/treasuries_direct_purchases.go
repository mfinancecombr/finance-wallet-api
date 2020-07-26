// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllTreasuriesDirectPurchases(c echo.Context) error {
	log.Debug("[API] Retrieving all treasuries direct purchases")
	result, err := s.db.GetAllTreasuriesDirectsPurchases()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getTreasuryDirectPurchaseByID godoc
// @Summary Get treasury direct purchase by ID
// @Description get treasury direct purchase data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.TreasuryDirect
// @Router /treasuries-direct/purchases/{id} [get]
// @Param id path string true "Purchase id"
func (s *server) getTreasuryDirectPurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving treasury direct purchase with id: %s", id)
	result, err := s.db.GetTreasuryDirectPurchaseByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Treasury direct purchase data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertTreasuryDirectPurchase godoc
// @Summary Insert some treasury direct purchase
// @Description insert new treasury direct purchase
// @Accept json
// @Produce json
// @Router /treasuries-direct/purchases [post]
func (s *server) insertTreasuryDirectPurchase(c echo.Context) error {
	log.Debugf("[API] Inserting treasury direct purchase")

	data := wallet.NewTreasuryDirect()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertTreasuryDirectPurchase(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateTreasuryDirectPurchaseByID godoc
// @Summary Update some treasury direct purchase
// @Description update new treasury direct purchase
// @Accept json
// @Produce json
// @Router /treasuries-direct/purchases/{id} [put]
// @Param id path string true "Purchase id"
func (s *server) updateTreasuryDirectPurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating treasury direct purchase with id %s", id)

	data := wallet.NewTreasuryDirect()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateTreasuryDirectPurchaseByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Treasury direct purchase not found")
}
