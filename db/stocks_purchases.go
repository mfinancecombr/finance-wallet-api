// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertStockPurchase(d *wallet.Stock) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertStockPurchase")
	return m.insertPurchase(d)
}

func (m *mongoSession) UpdateStockPurchaseByID(id string, d *wallet.Stock) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateStockPurchaseByID")
	return m.updateStockByID(purchasesCollection, id, d)
}

func (m *mongoSession) GetAllStocksPurchases() (wallet.StockList, error) {
	log.Debug("[DB] GetAllStocksPurchases")
	return m.getAllStocks(purchasesCollection)
}

func (m *mongoSession) GetStockPurchasesByPortfolioID(id string) (wallet.StockList, error) {
	log.Debug("[DB] GetStockPurchasesByPortfolioID")
	return m.getStockByPortfolioID(purchasesCollection, id)
}

func (m *mongoSession) GetStockPurchaseByID(id string) (*wallet.Stock, error) {
	log.Debug("[DB] GetStockPurchaseByID")
	return m.getStockByID(purchasesCollection, id)
}
