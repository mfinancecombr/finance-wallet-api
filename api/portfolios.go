// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"fmt"
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
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /portfolios/{slug} [get]
// @Param slug path string true "Broker slug"
// @Param year query string false "filter by year"
func (s *server) portfolio(c echo.Context) error {
	slug := c.Param("id")
	log.Debugf("[API] Retrieving %s data...", slug)

	// FIXME
	yearString := c.QueryParam("year")
	year, err := strconv.Atoi(yearString)
	if err != nil {
		log.Debugf("[API] Error on convert year: %v", err)
		// FIXME
		year = 2020
	}

	result := &wallet.Portfolio{}
	if err := s.db.GetBySlug(slug, result); err != nil {
		errMsg := fmt.Sprintf("Error on get portfolio '%s': %v", slug, err)
		return logAndReturnError(c, errMsg)
	}

	if result.Name == "" {
		errMsg := fmt.Sprintf("Portfolio '%s' not found", slug)
		return c.JSON(http.StatusNotFound, errorMessage(errMsg))
	}

	// FIXME
	err = s.db.GetPortfolioData(result, year)
	if err != nil {
		errMsg := fmt.Sprintf("Error on get portfolio '%s' items: %v", slug, err)
		return logAndReturnError(c, errMsg)
	}

	return c.JSON(http.StatusOK, result)
}

// portfolios godoc
// @Summary List all portfolios
// @Description get all portfolio data
// @Accept json
// @Produce json
// @Success 200 {array} wallet.Portfolio
// @Failure 500 {object} api.ErrorMessage
// @Router /portfolios [get]
// @Param year query string false "filter by year"
func (s *server) portfolios(c echo.Context) error {
	log.Debug("Retrieving all portfolios")

	// FIXME
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil {
		log.Debugf("Error on convert year: %v", err)
		// FIXME
		year = 2020
	}

	allPortfolios, err := s.db.GetAll(&wallet.Portfolio{})
	if err != nil {
		errMsg := fmt.Sprintf("Error on get all portfolios: %v", err)
		return logAndReturnError(c, errMsg)
	}

	portfolios := make([]wallet.Portfolio, len(allPortfolios))
	for idx, p := range allPortfolios {
		portfolio := p.(*wallet.Portfolio)
		err := s.db.GetPortfolioData(portfolio, year)
		if err != nil {
			errMsg := fmt.Sprintf("Error on get portfolio items: %v", err)
			return logAndReturnError(c, errMsg)
		}
		portfolios[idx] = *portfolio
	}
	return c.JSON(http.StatusOK, portfolios)
}

// portfoliosAdd godoc
// @Summary Insert some portfolio
// @Description insert new portfolio
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /portfolios [post]
func (s *server) portfoliosAdd(c echo.Context) error {
	log.Debug("Insert portfolio data")

	portfolio := &wallet.Portfolio{}
	if err := c.Bind(portfolio); err != nil {
		errMsg := fmt.Sprintf("Error on bind portfolio: %v", err)
		return logAndReturnError(c, errMsg)
	}

	portfolio.Slug = slug.Make(portfolio.Name)

	if err := c.Validate(portfolio); err != nil {
		errMsg := fmt.Sprintf("Error on validate portfolio: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Create(portfolio)
	if err != nil {
		errMsg := fmt.Sprintf("Error on insert portfolio: %v", err)
		return logAndReturnError(c, errMsg)
	}

	return c.JSON(http.StatusOK, result)
}

// portfoliosDelete godoc
// @Summary Delete portfolio by ID
// @Description delete some portfolio by id
// @Accept json
// @Produce json
// @Success 200 {object} interface{}
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /portfolios/{id} [delete]
// @Param id path string true "Portfolio id"
func (s *server) portfoliosDelete(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Deleting %s data", id)
	result, err := s.db.Delete("portfolios", id)
	if err != nil {
		errMsg := fmt.Sprintf("Error on delete portolio '%s': %v", id, err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// portfoliosUpdate godoc
// @Summary Update portfolio data by ID
// @Description Update some portfolio by id
// @Accept json
// @Produce json
// @Success 200 {object} wallet.Portfolio
// @Failure 404 {object} api.ErrorMessage
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /portfolios/{id} [put]
// @Param id path string true "Portfolio id"
func (s *server) portfoliosUpdate(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Updating %s data", id)

	portfolio := &wallet.Portfolio{}
	if err := c.Bind(portfolio); err != nil {
		errMsg := fmt.Sprintf("Error on bind portfolio: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(portfolio); err != nil {
		errMsg := fmt.Sprintf("Error on validate portfolio: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Update(id, portfolio)
	if err != nil {
		errMsg := fmt.Sprintf("Error on update portfolio: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	errMsg := fmt.Sprintf("Portfolio '%s' not found", id)
	return c.JSON(http.StatusNotFound, errorMessage(errMsg))
}
