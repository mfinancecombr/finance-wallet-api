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

func (s *server) getAllStockOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks operations")
	result, err := s.db.GetAll(wallet.Stock{})
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve all stocks operations: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// getStockOperationByID godoc
// @Summary Get stocks operation by ID
// @Description get stocks operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.Stock
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /stocks/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getStockOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock operation with id: %s", id)
	result := &wallet.Stock{}
	if err := s.db.Get(id, result); err != nil {
		errMsg := fmt.Sprintf("Error on retrieve '%s' operations: %v", id, err)
		return logAndReturnError(c, errMsg)
	}
	if result == nil {
		errMsg := fmt.Sprintf("Stock operation '%s' not found", id)
		return c.JSON(http.StatusNotFound, errorMessage(errMsg))
	}
	return c.JSON(http.StatusOK, result)
}

// insertStockOperation godoc
// @Summary Insert some stocks operation
// @Description insert new stocks operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /stocks/operations [post]
func (s *server) insertStockOperation(c echo.Context) error {
	log.Debugf("[API] Inserting stock operation")

	data := wallet.NewStock()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind stock: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate stock: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Create(data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on insert stock: %v", err)
		return logAndReturnError(c, errMsg)
	}

	return c.JSON(http.StatusOK, result)
}

// updateStockOperationByID godoc
// @Summary Update some stocks operation
// @Description update new stocks operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /stocks/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateStockOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock operation with id %s", id)

	data := wallet.NewStock()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind stock: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate stock: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Update(id, data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on update stock: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	errMsg := fmt.Sprintf("Stock operation '%s' not found", id)
	return c.JSON(http.StatusNotFound, errorMessage(errMsg))
}
