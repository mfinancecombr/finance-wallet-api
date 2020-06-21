// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/financeapi"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

func getYearQueryParam(c echo.Context) int {
	yearString := c.QueryParam("year")
	year, err := strconv.Atoi(yearString)
	if err != nil {
		log.Debugf("Error on convert year: %v", err)
		year, _, _ = time.Now().Date()
	}
	return year
}

func getMonthQueryParam(c echo.Context) int {
	monthString := c.QueryParam("month")
	month, err := strconv.Atoi(monthString)
	if err != nil {
		log.Debugf("Error on convert month: %v", err)
		_, m, _ := time.Now().Date()
		month = int(m)
	}
	return month
}

// portfolio godoc
// @Summary Get a portfolio
// @Description get all portfolio data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.Portfolio
// @Router /portfolios/{id} [get]
// @Param id path string true "Broker id"
// @Param month query string false "filter by month"
// @Param year query string false "filter by year"
func (s *server) portfolio(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Retrieving %s data...", id)

	portfolio, err := s.db.GetPortfolioByID(id)
	if err != nil {
		errMsg := fmt.Sprintf("Error on get portfolio: %v", err)
		log.Error(errMsg)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	if portfolio == nil {
		return c.JSON(http.StatusNotFound, "Portfolio data not found")
	}

	year := getYearQueryParam(c)
	month := getMonthQueryParam(c)

	err = s.db.GetPortfolioItems(portfolio, year, month)
	if err != nil {
		log.Errorf("Error on get portfolio items: %v", err)
		return err
	}

	// FIXME
	for symbol, item := range portfolio.Items {
		// FIXME
		params := url.Values{}
		if c.QueryParam("year") != "" {
			params.Add("year", c.QueryParam("year"))
		}
		if c.QueryParam("month") != "" {
			params.Add("month", c.QueryParam("month"))
		}
		baseUrl, _ := url.Parse(fmt.Sprintf("/%s/%s", item.ItemType, symbol))
		baseUrl.RawQuery = params.Encode()
		if err := financeapi.GetJSON(baseUrl.String(), &item); err != nil {
			log.Errorf("Error on get item '%s' in finance API", err)
			continue
		}
		item.Recalculate()
		portfolio.Items[symbol] = item
	}
	portfolio.Recalculate()

	return c.JSON(http.StatusOK, portfolio)
}

// portfolios godoc
// @Summary List all portfolios
// @Description get all portfolio data
// @Accept json
// @Produce json
// @Success 200 {array} wallet.Portfolio
// @Router /portfolios [get]
// @Param month query string false "filter by month"
// @Param year query string false "filter by year"
func (s *server) portfolios(c echo.Context) error {
	log.Debug("Retrieving all portfolios")

	allPortfolios, err := s.db.GetAllPortfolios()
	if err != nil {
		errMsg := fmt.Sprintf("Error on get portfolios: %v", err)
		log.Error(errMsg)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	year := getYearQueryParam(c)
	month := getMonthQueryParam(c)

	portfolios := make([]wallet.Portfolio, len(allPortfolios))
	for idx, portfolio := range allPortfolios {
		err := s.db.GetPortfolioItems(&portfolio, year, month)
		if err != nil {
			errMsg := fmt.Sprintf("Error on get portfolio items: %v", err)
			log.Errorf(errMsg)
			return c.JSON(http.StatusInternalServerError, errMsg)
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

func (s *server) portfolioYield(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Retrieving %s data...", id)

	result, err := s.db.GetPortfolioByID(id)
	if err != nil {
		errMsg := fmt.Sprintf("Error on get portfolio: %v", err)
		log.Error(errMsg)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	if result == nil {
		return c.JSON(http.StatusNotFound, "Portfolio data not found")
	}

	portfolioYield := wallet.PortfolioYield{}

	initialYear, initialMonth, _ := time.Now().Date()

	// FIXME Get via historical
	err = s.db.GetPortfolioItems(result, initialYear, int(initialMonth))
	if err != nil {
		log.Errorf("Error on get portfolio items: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	yield := wallet.PortolioYiedValue{
		Year:           initialYear,
		Month:          int(initialMonth),
		Percentage:     result.OverallReturn,
		Value:          result.Gain,
		IsConsolidated: true,
	}
	portfolioYield = append(portfolioYield, yield)

	for month := int(initialMonth) - 1; month >= 1; month-- {
		// FIXME Get via historical
		err = s.db.GetPortfolioItems(result, initialYear, month)
		if err != nil {
			log.Errorf("Error on get portfolio items: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		yield := wallet.PortolioYiedValue{
			Year:           initialYear,
			Month:          month,
			Percentage:     result.OverallReturn,
			Value:          result.Gain,
			IsConsolidated: false,
		}
		portfolioYield = append(portfolioYield, yield)
	}

	// FIXME Get via historical
	err = s.db.GetPortfolioItems(result, initialYear-1, 12)
	if err != nil {
		log.Errorf("Error on get portfolio items: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	yield = wallet.PortolioYiedValue{
		Year:           initialYear - 1,
		Month:          12,
		Percentage:     result.OverallReturn,
		Value:          result.Gain,
		IsConsolidated: true,
	}
	portfolioYield = append(portfolioYield, yield)

	return c.JSON(http.StatusOK, portfolioYield)
}
