// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package api

import (
	"fmt"
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
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /brokers/{id} [get]
// @Param id path string true "Broker id"
func (s *server) broker(c echo.Context) error {
	slug := c.Param("id")
	log.Debugf("[API] Retrieving broker slug: %s", slug)
	result := &wallet.Broker{}
	if err := s.db.GetBySlug(slug, result); err != nil {
		errMsg := fmt.Sprintf("Error on retrieve broker id '%s': %v", slug, err)
		return logAndReturnError(c, errMsg)
	}
	if result.Name == "" {
		errMsg := fmt.Sprintf("Broker '%s' not found", slug)
		return c.JSON(http.StatusNotFound, errorMessage(errMsg))
	}
	return c.JSON(http.StatusOK, result)
}

// brokers godoc
// @Summary List all brokers
// @Description get all brokers data
// @Accept json
// @Produce json
// @Success 200 {array} wallet.Broker
// @Failure 500 {object} api.ErrorMessage
// @Router /brokers [get]
func (s *server) brokers(c echo.Context) error {
	log.Debug("Retrieving all brokers")
	result, err := s.db.GetAll(&wallet.Broker{})
	if err != nil {
		errMsg := fmt.Sprintf("Error on retrieve brokers: %v", err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// brokersAdd godoc
// @Summary Insert some broker
// @Description insert new broker
// @Accept json
// @Produce json
// @Success 200 {array} interface{}
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /brokers [post]
func (s *server) brokersAdd(c echo.Context) error {
	log.Debug("Insert brokers data")

	broker := &wallet.Broker{}
	if err := c.Bind(broker); err != nil {
		errMsg := fmt.Sprintf("Error on bind broker: %v", err)
		return logAndReturnError(c, errMsg)
	}

	broker.Slug = slug.Make(broker.Name)

	if err := c.Validate(broker); err != nil {
		errMsg := fmt.Sprintf("Error on validate broker: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Create(broker)
	if err != nil {
		errMsg := fmt.Sprintf("Error on insert broker: %v", err)
		return logAndReturnError(c, errMsg)
	}

	return c.JSON(http.StatusOK, result)
}

// brokersDelete godoc
// @Summary Delete broker by ID
// @Description delete some broker by id
// @Accept json
// @Produce json
// @Success 200 {array} interface{}
// @Failure 404 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /brokers/{id} [delete]
// @Param id path string true "Broker id"
func (s *server) brokersDelete(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Deleting %s data", id)
	result, err := s.db.Delete("brokers", id)
	if err != nil {
		errMsg := fmt.Sprintf("Error on delete broker '%s': %v", id, err)
		return logAndReturnError(c, errMsg)
	}
	return c.JSON(http.StatusOK, result)
}

// brokersUpdate godoc
// @Summary Update broker data by ID
// @Description Update some broker by id
// @Accept json
// @Produce json
// @Success 200 {array} interface{}
// @Failure 404 {object} api.ErrorMessage
// @Failure 422 {object} api.ErrorMessage
// @Failure 500 {object} api.ErrorMessage
// @Router /brokers/{id} [put]
// @Param id path string true "Broker id"
func (s *server) brokersUpdate(c echo.Context) error {
	id := c.Param("id")
	log.Debugf("Updating %s data", id)

	broker := &wallet.Broker{}
	if err := c.Bind(broker); err != nil {
		errMsg := fmt.Sprintf("Error on bind broker: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if err := c.Validate(broker); err != nil {
		errMsg := fmt.Sprintf("Error on validate broker: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, errorMessage(errMsg))
	}

	result, err := s.db.Update(id, broker)
	if err != nil {
		errMsg := fmt.Sprintf("Error on update broker: %v", err)
		return logAndReturnError(c, errMsg)
	}

	if result.MatchedCount != 0 {
		return c.JSON(http.StatusOK, result)
	}

	errMsg := fmt.Sprintf("Broker '%s' not found", id)
	return c.JSON(http.StatusNotFound, errorMessage(errMsg))
}
