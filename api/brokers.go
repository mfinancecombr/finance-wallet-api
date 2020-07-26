// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"net/http"

	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
)

// broker godoc
// @Summary Get a broker
// @Description get all broker data
// @Accept json
// @Produce json
// @Success 200 {object} wallet.Broker
// @Router /brokers/{id} [get]
// @Param id path string true "Broker id"
func (s *server) broker(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("[API] Retrieving broker id: %s", id)
	result, err := s.db.GetBrokerByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return c.JSON(http.StatusNotFound, "Broker data not found")
	}
	return c.JSON(http.StatusOK, result)
}

// brokers godoc
// @Summary List all brokers
// @Description get all brokers data
// @Accept json
// @Produce json
// @Success 200 {array} wallet.Broker
// @Router /brokers [get]
func (s *server) brokers(c echo.Context) error {
	log.Debug("Retrieving all brokers")

	result, err := s.db.GetAllBrokers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// brokersAdd godoc
// @Summary Insert some broker
// @Description insert new broker
// @Accept json
// @Produce json
// @Router /brokers [post]
func (s *server) brokersAdd(c echo.Context) error {
	log.Debug("Insert brokers data")

	broker := &wallet.Broker{}
	if err := c.Bind(broker); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	broker.ID = slug.Make(broker.Name)

	if err := c.Validate(broker); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.InsertBroker(broker)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

// brokersDelete godoc
// @Summary Delete broker by ID
// @Description delete some broker by id
// @Accept json
// @Produce json
// @Router /brokers/{id} [delete]
// @Param id path string true "Broker id"
func (s *server) brokersDelete(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Deleting %s data", id)
	result, err := s.db.DeleteBrokerByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

// brokersUpdate godoc
// @Summary Update broker data by ID
// @Description Update some broker by id
// @Accept json
// @Produce json
// @Router /brokers/{id} [put]
// @Param id path string true "Broker id"
func (s *server) brokersUpdate(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Updating %s data", id)

	broker := &wallet.Broker{}
	if err := c.Bind(broker); err != nil {
		log.Errorf("Error on bind: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(broker); err != nil {
		log.Errorf("Error on validate: %v", err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := s.db.UpdateBroker(id, broker)
	if err != nil {
		log.Errorf("Error on update broker: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	return c.JSON(http.StatusNotFound, "Broker not found")
}
