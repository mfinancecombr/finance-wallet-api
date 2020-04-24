// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertStockFundSale(d *wallet.StockFund) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertStockFundSale")
	return m.insertSale(d)
}

func (m *mongoSession) UpdateStockFundSaleByID(id string, d *wallet.StockFund) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateStockFundSaleByID")
	return m.updateStockFundByID(salesCollection, id, d)
}

func (m *mongoSession) GetAllStocksFundsSales() (wallet.StockFundList, error) {
	log.Debug("[DB] GetAllStocksFundsSales")
	return m.getAllStocksFunds(salesCollection)
}

func (m *mongoSession) GetStockFundSalesByPortfolioID(id string) (wallet.StockFundList, error) {
	log.Debug("[DB] GetStockFundSalesByPortfolioID")
	return m.getStockFundByPortfolioID(salesCollection, id)
}

func (m *mongoSession) GetStockFundSaleByID(id string) (*wallet.StockFund, error) {
	log.Debug("[DB] GetStockFundSaleByID")
	return m.getStockFundByID(salesCollection, id)
}
