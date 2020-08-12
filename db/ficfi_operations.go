// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertFICFIOperation(d *wallet.FICFI) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertFICFIOperation")
	return m.insertOperation(d)
}

func (m *mongoSession) UpdateFICFIOperationByID(id string, d *wallet.FICFI) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateFICFIOperationByID")
	return m.updateFICFIByID(operationsCollection, id, d)
}

func (m *mongoSession) GetAllFICFIOperations() (wallet.FICFIList, error) {
	log.Debug("[DB] GetAllFICFIOperations")
	return m.getAllFICFI(operationsCollection)
}

func (m *mongoSession) GetFICFIOperationsByPortfolioID(id string) (wallet.FICFIList, error) {
	log.Debug("[DB] GetFICFIOperationsByPortfolioID")
	return m.getFICFIByPortfolioID(operationsCollection, id)
}

func (m *mongoSession) GetFICFIOperationByID(id string) (*wallet.FICFI, error) {
	log.Debug("[DB] GetFICFIOperationByID")
	return m.getFICFIByID(operationsCollection, id)
}