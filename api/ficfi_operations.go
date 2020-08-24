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

func (s *server) getAllFICFIOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all FICFI operations")
	result, err := s.db.GetAll(wallet.FICFI{})
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve FICFI operations: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// getFICFIOperationByID godoc
// @Summary Get FICFI operation by ID
// @Description get FICFI operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.FICFI
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /ficfi/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getFICFIOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving FICFI operation with id: %s", id)
	result := &wallet.FICFI{}
	if err := s.db.Get(id, result); err != nil {
		errMsg := fmt.Sprintf("Error on retrieve '%s' operations: %v", id, err)
		return logAndReturnError(c, errMsg)
	}
	if result == nil {
		errMsg := fmt.Sprintf("FICFI operation '%s' not found", id)
		return c.JSON(http.StatusNotFound, errorMessage(errMsg))
	}
	return c.JSON(http.StatusOK, result)
}

// insertFICFIOperation godoc
// @Summary Insert some FICFI operation
// @Description insert new FICFI operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /ficfi/operations [post]
func (s *server) insertFICFIOperation(c echo.Context) error {
	log.Debugf("[API] Inserting FICFI operation")

	data := wallet.NewFICFI()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind FICFI: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate FICFI: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Create(data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on insert FICFI: %v", err)
		return logAndReturnError(c, errMsg)
	}

	return c.JSON(http.StatusOK, result)
}

// updateFICFIOperationByID godoc
// @Summary Update some FICFI operation
// @Description update new FICFI operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 404 {object} api.ErrorMessage
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /ficfi/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateFICFIOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating FICFI operation with id %s", id)

	data := wallet.NewFICFI()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind FICFI: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate FICFI: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Update(id, data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on update FICFI: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	errMsg := fmt.Sprintf("FICFI operation '%s' not found", id)
	return c.JSON(http.StatusNotFound, errorMessage(errMsg))
}
