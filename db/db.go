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
	operationsCollection  = "operations"
)

type mongoSession struct {
	collection Collection
}

type DB interface {
	GetAllStocksOperations() (wallet.StockList, error)
	GetStockOperationByID(id string) (*wallet.Stock, error)
	InsertStockOperation(d *wallet.Stock) (*mongo.InsertOneResult, error)
	UpdateStockOperationByID(id string, d *wallet.Stock) (*mongo.UpdateResult, error)

	GetAllFIIsOperations() (wallet.FIIList, error)
	GetFIIOperationByID(id string) (*wallet.FII, error)
	InsertFIIOperation(d *wallet.FII) (*mongo.InsertOneResult, error)
	UpdateFIIOperationByID(id string, d *wallet.FII) (*mongo.UpdateResult, error)

	GetAllTreasuriesDirectsOperations() (wallet.TreasuryDirectList, error)
	GetTreasuryDirectOperationByID(id string) (*wallet.TreasuryDirect, error)
	InsertTreasuryDirectOperation(d *wallet.TreasuryDirect) (*mongo.InsertOneResult, error)
	UpdateTreasuryDirectOperationByID(id string, d *wallet.TreasuryDirect) (*mongo.UpdateResult, error)

	GetAllCertificatesOfDepositsOperations() (wallet.CertificateOfDepositList, error)
	GetCertificateOfDepositOperationByID(id string) (*wallet.CertificateOfDeposit, error)
	InsertCertificateOfDepositOperation(d *wallet.CertificateOfDeposit) (*mongo.InsertOneResult, error)
	UpdateCertificateOfDepositOperationByID(id string, d *wallet.CertificateOfDeposit) (*mongo.UpdateResult, error)

	GetAllStocksFundsOperations() (wallet.StockFundList, error)
	GetStockFundOperationByID(id string) (*wallet.StockFund, error)
	InsertStockFundOperation(d *wallet.StockFund) (*mongo.InsertOneResult, error)
	UpdateStockFundOperationByID(id string, d *wallet.StockFund) (*mongo.UpdateResult, error)

	GetAllFICFIOperations() (wallet.FICFIList, error)
	GetFICFIOperationByID(id string) (*wallet.FICFI, error)
	InsertFICFIOperation(d *wallet.FICFI) (*mongo.InsertOneResult, error)
	UpdateFICFIOperationByID(id string, d *wallet.FICFI) (*mongo.UpdateResult, error)

	DeleteBrokerByID(id string) (*mongo.DeleteResult, error)
	GetAllBrokers() (*wallet.BrokersList, error)
	GetBrokerByID(id string) (*wallet.Broker, error)
	InsertBroker(d interface{}) (*mongo.InsertOneResult, error)
	UpdateBroker(id string, d interface{}) (*mongo.UpdateResult, error)

	DeletePortfolioByID(id string) (*mongo.DeleteResult, error)
	GetAllPortfolios() ([]wallet.Portfolio, error)
	GetPortfolioByID(id string) (*wallet.Portfolio, error)
	GetPortfolioItems(p *wallet.Portfolio, year int) error
	InsertPortfolio(d interface{}) (*mongo.InsertOneResult, error)
	UpdatePortfolio(id string, d interface{}) (*mongo.UpdateResult, error)

	GetAllOperations() (interface{}, error)
	DeleteOperationByID(id string) (*mongo.DeleteResult, error)

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
