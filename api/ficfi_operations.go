// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllFICFIOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all FICFI operations")
	result, err := s.db.GetAllFICFIOperations()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getFICFIOperationByID godoc
// @Summary Get FICFI operation by ID
// @Description get FICFI operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.FICFI
// @Router /ficfi/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getFICFIOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving FICFI operation with id: %s", id)
	result, err := s.db.GetFICFIOperationByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "FICFI operation data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertFICFIOperation godoc
// @Summary Insert some FICFI operation
// @Description insert new FICFI operation
// @Accept json
// @Produce json
// @Router /ficfi/operations [post]
func (s *server) insertFICFIOperation(c echo.Context) error {
	log.Debugf("[API] Inserting FICFI operation")

	data := wallet.NewFICFI()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertFICFIOperation(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateFICFIOperationByID godoc
// @Summary Update some FICFI operation
// @Description update new FICFI operation
// @Accept json
// @Produce json
// @Router /ficfi/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateFICFIOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating FICFI operation with id %s", id)

	data := wallet.NewFICFI()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateFICFIOperationByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "FICFI operation not found")
}
