// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertFIIPurchase(d *wallet.FII) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertFIIPurchase")
	return m.insertPurchase(d)
}

func (m *mongoSession) UpdateFIIPurchaseByID(id string, d *wallet.FII) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateFIIPurchaseByID")
	return m.updateFIIByID(purchasesCollection, id, d)
}

func (m *mongoSession) GetAllFIIsPurchases() (wallet.FIIList, error) {
	log.Debug("[DB] GetAllFIIsPurchases")
	return m.getAllFIIs(purchasesCollection)
}

func (m *mongoSession) GetFIIPurchasesByPortfolioID(id string) (wallet.FIIList, error) {
	log.Debug("[DB] GetFIIPurchasesByPortfolioID")
	return m.getFIIByPortfolioID(purchasesCollection, id)
}

func (m *mongoSession) GetFIIPurchaseByID(id string) (*wallet.FII, error) {
	log.Debug("[DB] GetFIIPurchaseByID")
	return m.getFIIByID(purchasesCollection, id)
}
