// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertStockSale(d *wallet.Stock) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertStockSale")
	return m.insertSale(d)
}

func (m *mongoSession) UpdateStockSaleByID(id string, d *wallet.Stock) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateStockSaleByID")
	return m.updateStockByID(salesCollection, id, d)
}

func (m *mongoSession) GetAllStocksSales() (wallet.StockList, error) {
	log.Debug("[DB] GetAllStocksSales")
	return m.getAllStocks(salesCollection)
}

func (m *mongoSession) GetStockSalesByPortfolioID(id string) (wallet.StockList, error) {
	log.Debug("[DB] GetStockSalesByPortfolioID")
	return m.getStockByPortfolioID(salesCollection, id)
}

func (m *mongoSession) GetStockSaleByID(id string) (*wallet.Stock, error) {
	log.Debug("[DB] GetStockSaleByID")
	return m.getStockByID(salesCollection, id)
}
