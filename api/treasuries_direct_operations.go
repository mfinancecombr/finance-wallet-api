// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllTreasuriesDirectOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all treasuries direct operations")
	result, err := s.db.GetAllTreasuriesDirectsOperations()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getTreasuryDirectOperationByID godoc
// @Summary Get treasury direct operation by ID
// @Description get treasury direct  operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.TreasuryDirect
// @Router /treasuries-direct/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getTreasuryDirectOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving treasury direct operation with id: %s", id)
	result, err := s.db.GetTreasuryDirectOperationByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Treasury direct operation data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertTreasuryDirectOperation godoc
// @Summary Insert some treasury direct operation
// @Description insert new treasury direct operation
// @Accept json
// @Produce json
// @Router /treasuries-direct/operations [post]
func (s *server) insertTreasuryDirectOperation(c echo.Context) error {
	log.Debugf("[API] Inserting treasury direct operation")

	data := wallet.NewTreasuryDirect()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertTreasuryDirectOperation(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateTreasuryDirectOperationByID godoc
// @Summary Update some treasury direct operation
// @Description update new treasury direct operation
// @Accept json
// @Produce json
// @Router /treasuries-direct/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateTreasuryDirectOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating treasury direct operation with id %s", id)

	data := wallet.NewTreasuryDirect()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateTreasuryDirectOperationByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Treasury direct operation not found")
}
