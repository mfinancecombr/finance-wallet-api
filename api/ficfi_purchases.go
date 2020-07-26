// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllFICFIPurchases(c echo.Context) error {
	log.Debug("[API] Retrieving all FICFI purchases")
	result, err := s.db.GetAllFICFIPurchases()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getFICFIPurchaseByID godoc
// @Summary Get FICFI purchase by ID
// @Description get FIFCI purchase data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.FICFI
// @Router /ficfi/purchases/{id} [get]
// @Param id path string true "Purchase id"
func (s *server) getFICFIPurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving FICFI purchase with id: %s", id)
	result, err := s.db.GetFICFIPurchaseByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "FICFI purchase data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertFICFIPurchase godoc
// @Summary Insert some FICFI purchase
// @Description insert new FICFI purchase
// @Accept json
// @Produce json
// @Router /ficfi/purchases [post]
func (s *server) insertFICFIPurchase(c echo.Context) error {
	log.Debugf("[API] Inserting FICFI purchase")

	data := wallet.NewFICFI()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertFICFIPurchase(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateFICFIPurchaseByID godoc
// @Summary Update some FICFI purchase
// @Description update new FICFI purchase
// @Accept json
// @Produce json
// @Router /ficfi/purchases/{id} [put]
// @Param id path string true "Purchase id"
func (s *server) updateFICFIPurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating FICFI purchase with id %s", id)

	data := wallet.NewFICFI()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateFICFIPurchaseByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "FICFI purchase not found")
}
