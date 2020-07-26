// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllFICFISales(c echo.Context) error {
	log.Debug("[API] Retrieving all FICFI sales")
	result, err := s.db.GetAllFICFISales()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getFICFISaleByID godoc
// @Summary Get FICFI sale by ID
// @Description get FICFI sale data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.FICFI
// @Router /ficfi/sales/{id} [get]
// @Param id path string true "Sale id"
func (s *server) getFICFISaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving FICFI sale with id: %s", id)
	result, err := s.db.GetFICFISaleByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "FICFI sale data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertFICFISale godoc
// @Summary Insert some FICFI sale
// @Description insert new FICFI sale
// @Accept json
// @Produce json
// @Router /ficfi/sales [post]
func (s *server) insertFICFISale(c echo.Context) error {
	log.Debugf("[API] Inserting FICFI sale")

	data := wallet.NewFICFI()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertFICFISale(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateFICFISaleByID godoc
// @Summary Update some FICFI sale
// @Description update new FICFI sale
// @Accept json
// @Produce json
// @Router /ficfi/sales/{id} [put]
// @Param id path string true "Sale id"
func (s *server) updateFICFISaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating FICFI sale with id %s", id)

	data := wallet.NewFICFI()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateFICFISaleByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "FICFI sale not found")
}
