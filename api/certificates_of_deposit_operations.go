// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllCertificatesOfDepositOperations(c echo.Context) error {
	log.Debug("[API] Retrieving all certificates of deposit operations")
	result, err := s.db.GetAllCertificatesOfDepositsOperations()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getCertificateOfDepositOperationByID godoc
// @Summary Get get certificate of deposit operation by ID
// @Description get certificate of deposit operation data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.CertificateOfDeposit
// @Router /certificates-of-deposit/operations/{id} [get]
// @Param id path string true "Operation id"
func (s *server) getCertificateOfDepositOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving certificate of deposit operation with id: %s", id)
	result, err := s.db.GetCertificateOfDepositOperationByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Certificate of deposit operation data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertCertificateOfDepositOperation godoc
// @Summary Insert some certificate of deposit operation
// @Description insert new certificate of deposit operation
// @Accept json
// @Produce json
// @Router /certificates-of-deposit/operations [post]
func (s *server) insertCertificateOfDepositOperation(c echo.Context) error {
	log.Debugf("[API] Inserting certificate of deposit operation")

	data := wallet.NewCertificateOfDeposit()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertCertificateOfDepositOperation(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateCertificateOfDepositOperationByID godoc
// @Summary Update some certificate of deposit operation
// @Description update new certificate of deposit operation
// @Accept json
// @Produce json
// @Router /certificates-of-deposit/operations/{id} [put]
// @Param id path string true "Operation id"
func (s *server) updateCertificateOfDepositOperationByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating certificate of deposit operation with id %s", id)

	data := wallet.NewCertificateOfDeposit()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateCertificateOfDepositOperationByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Certificate of deposit operation not found")
}
