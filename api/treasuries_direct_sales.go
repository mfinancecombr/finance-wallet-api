// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func (s *server) getAllTreasuriesDirectSales(c echo.Context) error {
	log.Debug("[API] Retrieving all treasuries direct sales")
	result, err := s.db.GetAllTreasuriesDirectsSales()
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

// getTreasuryDirectSaleByID godoc
// @Summary Get treasury direct sale by ID
// @Description get treasury direct  sale data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.TreasuryDirect
// @Router /treasuries-direct/sales/{id} [get]
// @Param id path string true "Sale id"
func (s *server) getTreasuryDirectSaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving treasury direct sale with id: %s", id)
	result, err := s.db.GetTreasuryDirectSaleByID(id)
	if err != nil {
		log.Errorf("[API] Error on retrieve data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Treasury direct sale data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// insertTreasuryDirectSale godoc
// @Summary Insert some treasury direct sale
// @Description insert new treasury direct sale
// @Accept json
// @Produce json
// @Router /treasuries-direct/sales [post]
func (s *server) insertTreasuryDirectSale(c echo.Context) error {
	log.Debugf("[API] Inserting treasury direct sale")

	data := wallet.NewTreasuryDirect()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertTreasuryDirectSale(data)
	if err != nil {
		log.Errorf("[API] Error on insert: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// updateTreasuryDirectSaleByID godoc
// @Summary Update some treasury direct sale
// @Description update new treasury direct sale
// @Accept json
// @Produce json
// @Router /treasuries-direct/sales/{id} [put]
// @Param id path string true "Sale id"
func (s *server) updateTreasuryDirectSaleByID(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Updating treasury direct sale with id %s", id)

	data := wallet.NewTreasuryDirect()

	if err := c.Bind(data); err != nil {
		log.Errorf("[API] Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(data); err != nil {
		log.Errorf("[API] Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateTreasuryDirectSaleByID(id, data)
	if err != nil {
		log.Errorf("[API] Error on update: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Treasury direct sale not found")
}
