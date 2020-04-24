// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mfinancecombr/finance-wallet-api/db"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

const Version = "0.1.0"

type Server interface {
	http.Handler
	Start()
}

type server struct {
	*echo.Echo
	db db.DB
}

func (s *server) Start() {
	addr := fmt.Sprintf(":%d", viper.GetInt("port"))
	s.Echo.Logger.Fatal(s.Echo.Start(addr))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewServerFromDB() (Server, error) {
	echoInstance := echo.New()
	echoInstance.HideBanner = true
	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return nil, err
	}

	server := &server{
		Echo: echoInstance,
		db:   dbInstance,
	}

	echoInstance.Use(
		middleware.LoggerWithConfig(
			middleware.LoggerConfig{
				Format: "timestamp=${time_rfc3339} " +
					"method=${method} " +
					"request_uri=${uri} " +
					"status=${status} " +
					"request_id=${id} " +
					"latency=${latency_human}\n",
			},
		),
	)
	echoInstance.Use(middleware.Recover())
	// FIXME
	echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			echo.GET, echo.OPTIONS, echo.POST, echo.DELETE, echo.PUT,
		},
	}))
	echoInstance.Pre(middleware.RemoveTrailingSlash())

	echoInstance.Validator = &CustomValidator{validator: validator.New()}

	echoInstance.File("/favicon.ico", "images/favicon.ico")
	echoInstance.GET("/", server.index)
	echoInstance.GET("/healthcheck", server.healthcheck)
	echoInstance.Static("/static/icons", "images/icons")

	echoInstance.DELETE("/api/v1/purchases/:id", server.deletePurchaseByID)
	echoInstance.GET("/api/v1/purchases", server.getAllPurchases)

	echoInstance.GET("/api/v1/sales", server.getAllSales)
	echoInstance.DELETE("/api/v1/sales/:id", server.deleteSaleByID)

	echoInstance.DELETE("/api/v1/brokers/:id", server.brokersDelete)
	echoInstance.GET("/api/v1/brokers", server.brokers)
	echoInstance.GET("/api/v1/brokers/:id", server.broker)
	echoInstance.POST("/api/v1/brokers", server.brokersAdd)
	echoInstance.PUT("/api/v1/brokers/:id", server.brokersUpdate)

	echoInstance.DELETE("/api/v1/portfolios/:id", server.portfoliosDelete)
	echoInstance.GET("/api/v1/portfolios", server.portfolios)
	echoInstance.GET("/api/v1/portfolios/:id", server.portfolio)
	echoInstance.POST("/api/v1/portfolios", server.portfoliosAdd)
	echoInstance.PUT("/api/v1/portfolios/:id", server.portfoliosUpdate)

	echoInstance.GET("/api/v1/stocks/purchases/:id", server.getStockPurchaseByID)
	echoInstance.GET("/api/v1/stocks/sales/:id", server.getStockSaleByID)
	echoInstance.POST("/api/v1/stocks/purchases", server.insertStockPurchase)
	echoInstance.POST("/api/v1/stocks/sales", server.insertStockSale)
	echoInstance.PUT("/api/v1/stocks/purchases/:id", server.updateStockPurchaseByID)
	echoInstance.PUT("/api/v1/stocks/sales/:id", server.updateStockSaleByID)

	echoInstance.GET("/api/v1/fiis/purchases/:id", server.getFIIPurchaseByID)
	echoInstance.GET("/api/v1/fiis/sales/:id", server.getFIISaleByID)
	echoInstance.POST("/api/v1/fiis/purchases", server.insertFIIPurchase)
	echoInstance.POST("/api/v1/fiis/sales", server.insertFIISale)
	echoInstance.PUT("/api/v1/fiis/purchases/:id", server.updateFIIPurchaseByID)
	echoInstance.PUT("/api/v1/fiis/sales/:id", server.updateFIISaleByID)

	echoInstance.GET("/api/v1/treasuries-direct/purchases/:id", server.getTreasuryDirectPurchaseByID)
	echoInstance.GET("/api/v1/treasuries-direct/sales/:id", server.getTreasuryDirectSaleByID)
	echoInstance.POST("/api/v1/treasuries-direct/purchases", server.insertTreasuryDirectPurchase)
	echoInstance.POST("/api/v1/treasuries-direct/sales", server.insertTreasuryDirectSale)
	echoInstance.PUT("/api/v1/treasuries-direct/purchases/:id", server.updateTreasuryDirectPurchaseByID)
	echoInstance.PUT("/api/v1/treasuries-direct/sales/:id", server.updateTreasuryDirectSaleByID)

	echoInstance.GET("/api/v1/certificates-of-deposit/purchases/:id", server.getCertificateOfDepositPurchaseByID)
	echoInstance.GET("/api/v1/certificates-of-deposit/sales/:id", server.getCertificateOfDepositSaleByID)
	echoInstance.POST("/api/v1/certificates-of-deposit/purchases", server.insertCertificateOfDepositPurchase)
	echoInstance.POST("/api/v1/certificates-of-deposit/sales", server.insertCertificateOfDepositSale)
	echoInstance.PUT("/api/v1/certificates-of-deposit/purchases/:id", server.updateCertificateOfDepositPurchaseByID)
	echoInstance.PUT("/api/v1/certificates-of-deposit/sales/:id", server.updateCertificateOfDepositSaleByID)

	echoInstance.GET("/api/v1/stocks-funds/purchases/:id", server.getStockFundPurchaseByID)
	echoInstance.GET("/api/v1/stocks-funds/sales/:id", server.getStockFundSaleByID)
	echoInstance.POST("/api/v1/stocks-funds/purchases", server.insertStockFundPurchase)
	echoInstance.POST("/api/v1/stocks-funds/sales", server.insertStockFundSale)
	echoInstance.PUT("/api/v1/stocks-funds/purchases/:id", server.updateStockFundPurchaseByID)
	echoInstance.PUT("/api/v1/stocks-funds/sales/:id", server.updateStockFundSaleByID)

	echoInstance.GET("/api/v1/ficfi/purchases/:id", server.getFICFIPurchaseByID)
	echoInstance.GET("/api/v1/ficfi/sales/:id", server.getFICFISaleByID)
	echoInstance.POST("/api/v1/ficfi/purchases", server.insertFICFIPurchase)
	echoInstance.POST("/api/v1/ficfi/sales", server.insertFICFISale)
	echoInstance.PUT("/api/v1/ficfi/purchases/:id", server.updateFICFIPurchaseByID)
	echoInstance.PUT("/api/v1/ficfi/sales/:id", server.updateFICFISaleByID)

	return server, nil
}
