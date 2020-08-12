// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertStockOperation(d *wallet.Stock) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertStockOperation")
	return m.insertOperation(d)
}

func (m *mongoSession) UpdateStockOperationByID(id string, d *wallet.Stock) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateStockOperationByID")
	return m.updateStockByID(operationsCollection, id, d)
}

func (m *mongoSession) GetAllStocksOperations() (wallet.StockList, error) {
	log.Debug("[DB] GetAllStocksOperations")
	return m.getAllStocks(operationsCollection)
}

func (m *mongoSession) GetStockOperationsByPortfolioID(id string) (wallet.StockList, error) {
	log.Debug("[DB] GetStockOperationsByPortfolioID")
	return m.getStockByPortfolioID(operationsCollection, id)
}

func (m *mongoSession) GetStockOperationByID(id string) (*wallet.Stock, error) {
	log.Debug("[DB] GetStockOperationByID")
	return m.getStockByID(operationsCollection, id)
}
