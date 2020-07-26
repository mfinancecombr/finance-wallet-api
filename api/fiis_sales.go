// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllFIISales(c echo.Context) error {
	log.Debug("[API] Retrieving all stocks sales")
	result, err := s.db.GetAllFIIsSales()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getFIISaleByID godoc
// @Summary Get FII sale by ID
// @Description get FII sale data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.FII
// @Router /fiis/sales/{id} [get]
// @Param id path string true "Sale id"
func (s *server) getFIISaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving stock sale with id: %s", id)
	result, err := s.db.GetFIISaleByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "FII sale data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertFIISale godoc
// @Summary Insert some FII sale
// @Description insert new FII sale
// @Accept json
// @Produce json
// @Router /fiis/sales [post]
func (s *server) insertFIISale(c echo.Context) error {
	log.Debugf("[API] Inserting stock sale")

	data := wallet.NewFII()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertFIISale(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateFIISaleByID godoc
// @Summary Update some FII sale
// @Description update new FII sale
// @Accept json
// @Produce json
// @Router /fiis/sales/{id} [put]
// @Param id path string true "Sale id"
func (s *server) updateFIISaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating stock sale with id %s", id)

	data := wallet.NewFII()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateFIISaleByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "FII sale not found")
}
