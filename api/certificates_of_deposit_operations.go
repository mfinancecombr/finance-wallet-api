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

func (s *server) getAllCertificatesOfDepositOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all certificates of deposit operations")
	result, err := s.db.GetAll(wallet.CertificateOfDeposit{})
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve certificates of deposit operations: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// getCertificateOfDepositOperationByID godoc
// @Summary Get get certificate of deposit operation by ID
// @Description get certificate of deposit operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.CertificateOfDeposit
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /certificates-of-deposit/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getCertificateOfDepositOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving certificate of deposit operation with id: %s", id)
	result := &wallet.CertificateOfDeposit{}
	if err := s.db.Get(id, result); err != nil {
		errMsg := fmt.Sprintf("Error on retrieve '%s' operations: %v", id, err)
		return logAndReturnError(c, errMsg)
	}
	if result == nil {
		errMsg := fmt.Sprintf("Certificate of deposit operation '%s' not found", id)
		return c.JSON(http.StatusNotFound, errorMessage(errMsg))
	}
	return c.JSON(http.StatusOK, result)
}

// insertCertificateOfDepositOperation godoc
// @Summary Insert some certificate of deposit operation
// @Description insert new certificate of deposit operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /certificates-of-deposit/operations [post]
func (s *server) insertCertificateOfDepositOperation(c echo.Context) error {
	log.Debugf("[API] Inserting certificate of deposit operation")

	data := wallet.NewCertificateOfDeposit()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind certificate of deposit: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate certificate of deposit: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Create(data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on insert certificate of deposit: %v", err)
		return logAndReturnError(c, errMsg)
	}

	return c.JSON(http.StatusOK, result)
}

// updateCertificateOfDepositOperationByID godoc
// @Summary Update some certificate of deposit operation
// @Description update new certificate of deposit operation
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 404 {object} api.ErrorMessage
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /certificates-of-deposit/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateCertificateOfDepositOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating certificate of deposit operation with id %s", id)

	data := wallet.NewCertificateOfDeposit()

	if err := c.Bind(data); err != nil {
		errMsg := fmt.Sprintf("Error on bind certificate of deposit: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(data); err != nil {
		errMsg := fmt.Sprintf("Error on validate certificate of deposit: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Update(id, data)
	if err != nil {
		errMsg := fmt.Sprintf("Error on update certificate of deposit: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	errMsg := fmt.Sprintf("Certificate of deposit operation '%s' not found", id)
	return c.JSON(http.StatusNotFound, errorMessage(errMsg))
}
