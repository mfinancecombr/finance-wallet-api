// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertStockFundPurchase(d *wallet.StockFund) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertStockFundPurchase")
	return m.insertPurchase(d)
}

func (m *mongoSession) UpdateStockFundPurchaseByID(id string, d *wallet.StockFund) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateStockFundPurchaseByID")
	return m.updateStockFundByID(purchasesCollection, id, d)
}

func (m *mongoSession) GetAllStocksFundsPurchases() (wallet.StockFundList, error) {
	log.Debug("[DB] GetAllStocksFundsPurchases")
	return m.getAllStocksFunds(purchasesCollection)
}

func (m *mongoSession) GetStockFundPurchasesByPortfolioID(id string) (wallet.StockFundList, error) {
	log.Debug("[DB] GetStockFundPurchasesByPortfolioID")
	return m.getStockFundByPortfolioID(purchasesCollection, id)
}

func (m *mongoSession) GetStockFundPurchaseByID(id string) (*wallet.StockFund, error) {
	log.Debug("[DB] GetStockFundPurchaseByID")
	return m.getStockFundByID(purchasesCollection, id)
}
