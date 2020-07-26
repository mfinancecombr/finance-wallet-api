// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllFIIsPurchases(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks purchases")
	result, err := s.db.GetAllFIIsPurchases()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getFIIPurchaseByID godoc
// @Summary Get FII purchase by ID
// @Description get FII purchase data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.FII
// @Router /fiis/purchases/{id} [get]
// @Param id path string true "Purchase id"
func (s *server) getFIIPurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock purchase with id: %s", id)
	result, err := s.db.GetFIIPurchaseByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "FII purchase data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertFIIPurchase godoc
// @Summary Insert some FII purchase
// @Description insert new FII purchase
// @Accept json
// @Produce json
// @Router /fiis/purchases [post]
func (s *server) insertFIIPurchase(c echo.Context) error {
	log.Debugf("[API] Inserting stock purchase")

	data := wallet.NewFII()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertFIIPurchase(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateFIIPurchaseByID godoc
// @Summary Update some FII purchase
// @Description update new FII purchase
// @Accept json
// @Produce json
// @Router /fiis/purchases/{id} [put]
// @Param id path string true "Purchase id"
func (s *server) updateFIIPurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock purchase with id %s", id)

	data := wallet.NewFII()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateFIIPurchaseByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "FII purchase not found")
}
