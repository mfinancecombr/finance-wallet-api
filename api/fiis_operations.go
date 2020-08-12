// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllFIIOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks operations")
	result, err := s.db.GetAllFIIsOperations()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getFIIOperationByID godoc
// @Summary Get FII operation by ID
// @Description get FII operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.FII
// @Router /fiis/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getFIIOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock operation with id: %s", id)
	result, err := s.db.GetFIIOperationByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "FII operation data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertFIIOperation godoc
// @Summary Insert some FII operation
// @Description insert new FII operation
// @Accept json
// @Produce json
// @Router /fiis/operations [post]
func (s *server) insertFIIOperation(c echo.Context) error {
	log.Debugf("[API] Inserting stock operation")

	data := wallet.NewFII()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertFIIOperation(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateFIIOperationByID godoc
// @Summary Update some FII operation
// @Description update new FII operation
// @Accept json
// @Produce json
// @Router /fiis/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateFIIOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock operation with id %s", id)

	data := wallet.NewFII()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateFIIOperationByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "FII operation not found")
}
