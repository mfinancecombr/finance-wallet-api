// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllTreasuriesDirectOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all treasuries direct operations")
	result, err := s.db.GetAll(wallet.TreasuryDirect{})
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve treasuries direct operations: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// getTreasuryDirectOperationByID godoc
// @Summary Get treasury direct operation by ID
// @Description get treasury direct  operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.TreasuryDirect
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /treasuries-direct/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getTreasuryDirectOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving treasury direct operation with id: %s", id)
	result := &wallet.TreasuryDirect{}
	if err := s.db.Get(id, result); err != nil {
		errMsg := fmt.Sprintf("Error on retrieve '%s' operations: %v", id, err)
		return logAndReturnError(c, errMsg)
	}
	if result == nil {
		errMsg := fmt.Sprintf("Treasury direct operation '%s' not found", id)
		return c.JSON(http.StatusNotFound, errorMessage(errMsg))
	}
	return c.JSON(http.StatusOK, result)
}

// insertTreasuryDirectOperation godoc
// @Summary Insert some treasury direct operation
// @Description insert new treasury direct operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /treasuries-direct/operations [post]
func (s *server) insertTreasuryDirectOperation(c echo.Context) error {
	log.Debugf("[API] Inserting treasury direct operation")

	data := wallet.NewTreasuryDirect()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind treasury direct: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate treasury direct: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Create(data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on insert treasury direct: %v", err)
		return logAndReturnError(c, errMsg)
	}

	return c.JSON(http.StatusOK, result)
}

// updateTreasuryDirectOperationByID godoc
// @Summary Update some treasury direct operation
// @Description update new treasury direct operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 404 {object} api.ErrorMessage
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /treasuries-direct/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateTreasuryDirectOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating treasury direct operation with id %s", id)

	data := wallet.NewTreasuryDirect()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind treasury direct: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate treasury direct: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Update(id, data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on update treasury direct: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	errMsg := fmt.Sprintf("Treasury direct operation '%s' not found", id)
	return c.JSON(http.StatusNotFound, errorMessage(errMsg))
}
