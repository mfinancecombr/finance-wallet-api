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

func (s *server) getAllFIIOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks operations")
	result, err := s.db.GetAll(wallet.FII{})
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve operations: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// getFIIOperationByID godoc
// @Summary Get FII operation by ID
// @Description get FII operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.FII
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /fiis/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getFIIOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock operation with id: %s", id)
	result := &wallet.FII{}
	if err := s.db.Get(id, result); err != nil {
		errMsg := fmt.Sprintf("Error on retrieve '%s' operations: %v", id, err)
		return logAndReturnError(c, errMsg)
	}
	if result == nil {
		errMsg := fmt.Sprintf("FII operation '%s' not found", id)
		return c.JSON(http.StatusNotFound, errorMessage(errMsg))
	}
	return c.JSON(http.StatusOK, result)
}

// insertFIIOperation godoc
// @Summary Insert some FII operation
// @Description insert new FII operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /fiis/operations [post]
func (s *server) insertFIIOperation(c echo.Context) error {
	log.Debugf("[API] Inserting stock operation")

	data := wallet.NewFII()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind FII: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate FII: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Create(data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on insert FII: %v", err)
		return logAndReturnError(c, errMsg)
	}

	return c.JSON(http.StatusOK, result)
}

// updateFIIOperationByID godoc
// @Summary Update some FII operation
// @Description update new FII operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 404 {object} api.ErrorMessage
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /fiis/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateFIIOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock operation with id %s", id)

	data := wallet.NewFII()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind FII: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate FII: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Update(id, data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on update FII: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	errMsg := fmt.Sprintf("Stock operation '%s' not found", id)
	return c.JSON(http.StatusNotFound, errorMessage(errMsg))
}
