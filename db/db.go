// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"context"
	"time"

	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	brokersCollection    = "brokers"
	portfoliosCollection = "portfolios"
	purchasesCollection  = "purchases"
	salesCollection      = "sales"
)

type mongoSession struct {
	collection Collection
}

type DB interface {
	GetAllStocksPurchases() (wallet.StockList, error)
	GetAllStocksSales() (wallet.StockList, error)
	GetStockPurchaseByID(id string) (*wallet.Stock, error)
	GetStockSaleByID(id string) (*wallet.Stock, error)
	InsertStockPurchase(d *wallet.Stock) (*mongo.InsertOneResult, error)
	InsertStockSale(d *wallet.Stock) (*mongo.InsertOneResult, error)
	UpdateStockPurchaseByID(id string, d *wallet.Stock) (*mongo.UpdateResult, error)
	UpdateStockSaleByID(id string, d *wallet.Stock) (*mongo.UpdateResult, error)

	GetAllFIIsPurchases() (wallet.FIIList, error)
	GetAllFIIsSales() (wallet.FIIList, error)
	GetFIIPurchaseByID(id string) (*wallet.FII, error)
	GetFIISaleByID(id string) (*wallet.FII, error)
	InsertFIIPurchase(d *wallet.FII) (*mongo.InsertOneResult, error)
	InsertFIISale(d *wallet.FII) (*mongo.InsertOneResult, error)
	UpdateFIIPurchaseByID(id string, d *wallet.FII) (*mongo.UpdateResult, error)
	UpdateFIISaleByID(id string, d *wallet.FII) (*mongo.UpdateResult, error)

	GetAllTreasuriesDirectsPurchases() (wallet.TreasuryDirectList, error)
	GetAllTreasuriesDirectsSales() (wallet.TreasuryDirectList, error)
	GetTreasuryDirectPurchaseByID(id string) (*wallet.TreasuryDirect, error)
	GetTreasuryDirectSaleByID(id string) (*wallet.TreasuryDirect, error)
	InsertTreasuryDirectPurchase(d *wallet.TreasuryDirect) (*mongo.InsertOneResult, error)
	InsertTreasuryDirectSale(d *wallet.TreasuryDirect) (*mongo.InsertOneResult, error)
	UpdateTreasuryDirectPurchaseByID(id string, d *wallet.TreasuryDirect) (*mongo.UpdateResult, error)
	UpdateTreasuryDirectSaleByID(id string, d *wallet.TreasuryDirect) (*mongo.UpdateResult, error)

	GetAllCertificatesOfDepositsPurchases() (wallet.CertificateOfDepositList, error)
	GetAllCertificatesOfDepositsSales() (wallet.CertificateOfDepositList, error)
	GetCertificateOfDepositPurchaseByID(id string) (*wallet.CertificateOfDeposit, error)
	GetCertificateOfDepositSaleByID(id string) (*wallet.CertificateOfDeposit, error)
	InsertCertificateOfDepositPurchase(d *wallet.CertificateOfDeposit) (*mongo.InsertOneResult, error)
	InsertCertificateOfDepositSale(d *wallet.CertificateOfDeposit) (*mongo.InsertOneResult, error)
	UpdateCertificateOfDepositPurchaseByID(id string, d *wallet.CertificateOfDeposit) (*mongo.UpdateResult, error)
	UpdateCertificateOfDepositSaleByID(id string, d *wallet.CertificateOfDeposit) (*mongo.UpdateResult, error)

	GetAllStocksFundsPurchases() (wallet.StockFundList, error)
	GetAllStocksFundsSales() (wallet.StockFundList, error)
	GetStockFundPurchaseByID(id string) (*wallet.StockFund, error)
	GetStockFundSaleByID(id string) (*wallet.StockFund, error)
	InsertStockFundPurchase(d *wallet.StockFund) (*mongo.InsertOneResult, error)
	InsertStockFundSale(d *wallet.StockFund) (*mongo.InsertOneResult, error)
	UpdateStockFundPurchaseByID(id string, d *wallet.StockFund) (*mongo.UpdateResult, error)
	UpdateStockFundSaleByID(id string, d *wallet.StockFund) (*mongo.UpdateResult, error)

	GetAllFICFIPurchases() (wallet.FICFIList, error)
	GetAllFICFISales() (wallet.FICFIList, error)
	GetFICFIPurchaseByID(id string) (*wallet.FICFI, error)
	GetFICFISaleByID(id string) (*wallet.FICFI, error)
	InsertFICFIPurchase(d *wallet.FICFI) (*mongo.InsertOneResult, error)
	InsertFICFISale(d *wallet.FICFI) (*mongo.InsertOneResult, error)
	UpdateFICFIPurchaseByID(id string, d *wallet.FICFI) (*mongo.UpdateResult, error)
	UpdateFICFISaleByID(id string, d *wallet.FICFI) (*mongo.UpdateResult, error)

	DeleteBrokerByID(id string) (*mongo.DeleteResult, error)
	GetAllBrokers() (*wallet.BrokersList, error)
	GetBrokerByID(id string) (*wallet.Broker, error)
	InsertBroker(d interface{}) (*mongo.InsertOneResult, error)
	UpdateBroker(id string, d interface{}) (*mongo.UpdateResult, error)

	DeletePortfolioByID(id string) (*mongo.DeleteResult, error)
	GetAllPortfolios() ([]wallet.Portfolio, error)
	GetPortfolioByID(id string) (*wallet.Portfolio, error)
	GetPortfolioItems(p *wallet.Portfolio, year, month int) error
	InsertPortfolio(d interface{}) (*mongo.InsertOneResult, error)
	UpdatePortfolio(id string, d interface{}) (*mongo.UpdateResult, error)

	DeletePurchaseByID(id string) (*mongo.DeleteResult, error)
	GetAllPurchases() (interface{}, error)
	GetPurchasesItemType() ([]interface{}, error)

	DeleteSaleByID(id string) (*mongo.DeleteResult, error)
	GetAllSales() (interface{}, error)
	GetSalesItemType() ([]interface{}, error)

	Ping() error
}

func newDBContext() (context.Context, context.CancelFunc) {
	log.Debug("[DB] New DB context")
	timeout := viper.GetDuration("db.operation.timeout")
	return context.WithTimeout(context.Background(), timeout*time.Second)
}

func NewMongoSession() (DB, error) {
	log.Debug("[DB] New mongo session")
	dbURI := viper.GetString("mongodb.endpoint")
	dbName := viper.GetString("mongodb.name")
	ctx, _ := newDBContext()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Errorf("[DB] Error on create mongo session: %s", err)
	}
	mongo := &mongoSession{
		collection: &mongoCollection{
			session: client,
			dbName:  dbName,
		},
	}
	return mongo, err
}

func (m *mongoSession) Ping() error {
	log.Debug("[DB] Ping")
	return m.collection.Ping()
}
