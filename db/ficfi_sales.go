// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertFICFISale(d *wallet.FICFI) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertFICFISale")
	return m.insertSale(d)
}

func (m *mongoSession) UpdateFICFISaleByID(id string, d *wallet.FICFI) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateFICFISaleByID")
	return m.updateFICFIByID(salesCollection, id, d)
}

func (m *mongoSession) GetAllFICFISales() (wallet.FICFIList, error) {
	log.Debug("[DB] GetAllFICFISales")
	return m.getAllFICFI(salesCollection)
}

func (m *mongoSession) GetFICFISalesByPortfolioID(id string) (wallet.FICFIList, error) {
	log.Debug("[DB] GetFICFISalesByPortfolioID")
	return m.getFICFIByPortfolioID(salesCollection, id)
}

func (m *mongoSession) GetFICFISaleByID(id string) (*wallet.FICFI, error) {
	log.Debug("[DB] GetFICFISaleByID")
	return m.getFICFIByID(salesCollection, id)
}
