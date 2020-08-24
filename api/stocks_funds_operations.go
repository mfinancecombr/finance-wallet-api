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

func (s *server) getAllStockFundsOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks funds operations")
	result, err := s.db.GetAll(wallet.StockFund{})
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve stocks funds operations: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// getStockFundOperationByID godoc
// @Summary Get stocks fund operation by ID
// @Description get stocks fund operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.StockFund
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /stocks-funds/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getStockFundOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock fund operation with id: %s", id)
	result := &wallet.StockFund{}
	if err := s.db.Get(id, result); err != nil {
		errMsg := fmt.Sprintf("Error on retrieve '%s' operations: %v", id, err)
		return logAndReturnError(c, errMsg)
	}
	if result == nil {
		errMsg := fmt.Sprintf("stock fund operation '%s' not found", id)
		return c.JSON(http.StatusNotFound, errorMessage(errMsg))
	}
	return c.JSON(http.StatusOK, result)
}

// insertStockFundOperation godoc
// @Summary Insert some stocks fund operation
// @Description insert new stocks fund operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /stocks-funds/operations [post]
func (s *server) insertStockFundOperation(c echo.Context) error {
	log.Debugf("[API] Inserting stock fund operation")

	data := wallet.NewStockFund()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind stock fund: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate stock fund: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Create(data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on insert stock fund: %v", err)
		return logAndReturnError(c, errMsg)
	}

	return c.JSON(http.StatusOK, result)
}

// updateStockFundOperationByID godoc
// @Summary Update some stocks fund operation
// @Description update new stocks fund operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 404 {object} api.ErrorMessage
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /stocks-funds/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateStockFundOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock fund operation with id %s", id)

	data := wallet.NewStockFund()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind stock fund: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate stock fund: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Update(id, data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on update stock fund: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	errMsg := fmt.Sprintf("stock fund operation '%s' not found", id)
	return c.JSON(http.StatusNotFound, errorMessage(errMsg))
}
