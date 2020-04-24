// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllCertificatesOfDepositSales(c echo.Context) error {
	log.Debug("[API] Retrieving all certificates of deposit sales")
	result, err := s.db.GetAllCertificatesOfDepositsSales()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (s *server) getCertificateOfDepositSaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving certificate of deposit sale with id: %s", id)
	result, err := s.db.GetCertificateOfDepositSaleByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Certificate of deposit sale data not found")
	}
	return c.JSON(http.StatusOK, result)
}

func (s *server) insertCertificateOfDepositSale(c echo.Context) error {
	log.Debugf("[API] Inserting certificate of deposit sale")

	data := wallet.NewCertificateOfDeposit()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertCertificateOfDepositSale(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (s *server) updateCertificateOfDepositSaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating certificate of deposit sale with id %s", id)

	data := wallet.NewCertificateOfDeposit()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateCertificateOfDepositSaleByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Certificate of deposit sale not found")
}
