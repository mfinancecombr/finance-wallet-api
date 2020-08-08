// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllCertificatesOfDepositPurchases(c echo.Context) error {
	log.Debug("[API] Retrieving all certificates of deposit purchases")
	result, err := s.db.GetAllCertificatesOfDepositsPurchases()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getCertificateOfDepositPurchaseByID godoc
// @Summary Get get certificate of deposit purchase by ID
// @Description get certificate of deposi purchase data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.CertificateOfDeposit
// @Router /certificate-of-deposit/purchases/{id} [get]
// @Param id path string true "Purchase id"
func (s *server) getCertificateOfDepositPurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving certificate of deposit purchase with id: %s", id)
	result, err := s.db.GetCertificateOfDepositPurchaseByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Certificate of deposit purchase data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertCertificateOfDepositPurchase godoc
// @Summary Insert some certificate of deposit purchase
// @Description insert new certificate of deposit purchase
// @Accept json
// @Produce json
// @Router /certificate-of-deposit/purchases [post]
func (s *server) insertCertificateOfDepositPurchase(c echo.Context) error {
	log.Debugf("[API] Inserting certificate of deposit purchase")

	data := wallet.NewCertificateOfDeposit()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertCertificateOfDepositPurchase(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateCertificateOfDepositPurchaseByID godoc
// @Summary Update some certificate of deposit purchase
// @Description update new certificate of deposit purchase
// @Accept json
// @Produce json
// @Router /certificate-of-deposit/purchases/{id} [put]
// @Param id path string true "Purchase id"
func (s *server) updateCertificateOfDepositPurchaseByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating certificate of deposit purchase with id %s", id)

	data := wallet.NewCertificateOfDeposit()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateCertificateOfDepositPurchaseByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Certificate of deposit purchase not found")
}