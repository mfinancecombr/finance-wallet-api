// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertFIISale(d *wallet.FII) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertFIISale")
	return m.insertSale(d)
}

func (m *mongoSession) UpdateFIISaleByID(id string, d *wallet.FII) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateFIISaleByID")
	return m.updateFIIByID(salesCollection, id, d)
}

func (m *mongoSession) GetAllFIIsSales() (wallet.FIIList, error) {
	log.Debug("[DB] GetAllFIIsSales")
	return m.getAllFIIs(salesCollection)
}

func (m *mongoSession) GetFIISalesByPortfolioID(id string) (wallet.FIIList, error) {
	log.Debug("[DB] GetFIISalesByPortfolioID")
	return m.getFIIByPortfolioID(salesCollection, id)
}

func (m *mongoSession) GetFIISaleByID(id string) (*wallet.FII, error) {
	log.Debug("[DB] GetFIISaleByID")
	return m.getFIIByID(salesCollection, id)
}
