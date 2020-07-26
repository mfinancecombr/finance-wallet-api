// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

// portfolio godoc
// @Summary Get a portfolio
// @Description get all portfolio data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.Portfolio
// @Router /portfolios/{id} [get]
// @Param id path string true "Broker id"
// @Param year query string false "filter by year"
func (s *server) portfolio(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Retrieving %s data...", id)

	// FIXME
	yearString := c.QueryParam("year")
	year, err := strconv.Atoi(yearString)
	if err != nil {
		log.Errorf("Error on convert year: %v", err)
		// FIXME
		year = 2020
	}

	result, err := s.db.GetPortfolioByID(id)
	if err != nil {
		log.Errorf("Error on get portfolio: %v", err)
		return err
	}

	if result == nil {
		return c.JSON(http.StatusNotFound, "Portfolio data not found")
	}

	// FIXME
	err = s.db.GetPortfolioItems(result, year)
	if err != nil {
		log.Errorf("Error on get portfolio items: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// portfolios godoc
// @Summary List all portfolios
// @Description get all portfolio data
// @Accept json
// @Produce json
// @Success 200 {array} wallet.Portfolio
// @Router /portfolios [get]
// @Param year query string false "filter by year"
func (s *server) portfolios(c echo.Context) error {
	log.Debug("Retrieving all portfolios")

	// FIXME
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil {
		log.Errorf("Error on convert year: %v", err)
		// FIXME
		year = 2020
	}

	allPortfolios, err := s.db.GetAllPortfolios()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	portfolios := make([]wallet.Portfolio, len(allPortfolios))
	for idx, portfolio := range allPortfolios {
		// FIXME
		err := s.db.GetPortfolioItems(&portfolio, year)
		if err != nil {
			log.Errorf("Error on get portfolio items: %v", err)
			return err
		}

		portfolios[idx] = portfolio
	}
	return c.JSON(http.StatusOK, portfolios)
}

// portfoliosAdd godoc
// @Summary Insert some portfolio
// @Description insert new portfolio
// @Accept json
// @Produce json
// @Router /portfolios [post]
func (s *server) portfoliosAdd(c echo.Context) error {
	log.Debug("Insert portfolio data")

	portfolio := &wallet.Portfolio{}
	if err := c.Bind(portfolio); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	portfolio.ID = slug.Make(portfolio.Name)

	if err := c.Validate(portfolio); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertPortfolio(portfolio)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// portfoliosDelete godoc
// @Summary Delete portfolio by ID
// @Description delete some portfolio by id
// @Accept json
// @Produce json
// @Router /portfolios/{id} [delete]
// @Param id path string true "Portfolio id"
func (s *server) portfoliosDelete(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Deleting %s data", id)
	result, err := s.db.DeletePortfolioByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

// portfoliosUpdate godoc
// @Summary Update portfolio data by ID
// @Description Update some portfolio by id
// @Accept json
// @Produce json
// @Router /portfolios/{id} [put]
// @Param id path string true "Portfolio id"
func (s *server) portfoliosUpdate(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Updating %s data", id)

	portfolio := &wallet.Portfolio{}
	if err := c.Bind(portfolio); err != nil {
		log.Errorf("Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(portfolio); err != nil {
		log.Errorf("Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdatePortfolio(id, portfolio)
	if err != nil {
		log.Errorf("Error on update portfolio: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Portfolio not found")
}
