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
	d.ID = ""
	return m.updateOperation(operationsCollection, id, d)
}

func (m *mongoSession) GetAllStocksOperations() (wallet.StockList, error) {
	log.Debug("[DB] GetAllStocksOperations")
	return m.getAllStocks(operationsCollection)
}

func (m *mongoSession) GetStockOperationByID(id string) (*wallet.Stock, error) {
	log.Debug("[DB] GetStockOperationByID")
	stock := &wallet.Stock{}
	if err := m.getOperationByID(operationsCollection, id, stock); err != nil {
		return nil, err
	}
	if stock.Symbol == "" {
		return nil, nil
	}
	return stock, nil
}
