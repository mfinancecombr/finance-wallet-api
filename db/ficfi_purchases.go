// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertFICFIPurchase(d *wallet.FICFI) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertFICFIPurchase")
	return m.insertPurchase(d)
}

func (m *mongoSession) UpdateFICFIPurchaseByID(id string, d *wallet.FICFI) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateFICFIPurchaseByID")
	return m.updateFICFIByID(purchasesCollection, id, d)
}

func (m *mongoSession) GetAllFICFIPurchases() (wallet.FICFIList, error) {
	log.Debug("[DB] GetAllFICFIPurchases")
	return m.getAllFICFI(purchasesCollection)
}

func (m *mongoSession) GetFICFIPurchasesByPortfolioID(id string) (wallet.FICFIList, error) {
	log.Debug("[DB] GetFICFIPurchasesByPortfolioID")
	return m.getFICFIByPortfolioID(purchasesCollection, id)
}

func (m *mongoSession) GetFICFIPurchaseByID(id string) (*wallet.FICFI, error) {
	log.Debug("[DB] GetFICFIPurchaseByID")
	return m.getFICFIByID(purchasesCollection, id)
}
