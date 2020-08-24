// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// getAllOperations godoc
// @Summary List all operations
// @Description get all operations data
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 500 {object} api.ErrorMessage
// @Router /operations [get]
func (s *server) getAllOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all operations")
	result, err := s.db.GetAllOperations()
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve all operations: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// getAllPurchases godoc
// @Summary List all purchases operations
// @Description get all purchases operations data
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 500 {object} api.ErrorMessage
// @Router /purchases [get]
func (s *server) getAllPurchases(c echo.Context) error {
	log.Debug("[API] Retrieving all purchases operations")
	result, err := s.db.GetAllPurchases()
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve purchases operations: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// getAllSales godoc
// @Summary List all sales operations
// @Description get all sales operations data
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 500 {object} api.ErrorMessage
// @Router /sales [get]
func (s *server) getAllSales(c echo.Context) error {
	log.Debug("[API] Retrieving all sales operations")
	result, err := s.db.GetAllSales()
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve sales operations: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// deleteOperationByID godoc
// @Summary Delete operation by ID
// @Description delete some operation by id
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /operations/{id} [delete]
// @Param id path string true "Operation id"
func (s *server) deleteOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Deleting %s data", id)
	result, err := s.db.Delete("operations", id)
	if err != nil {
		errMsg := fmt.Sprintf("Error on delete operation '%s': %v", id, err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}
