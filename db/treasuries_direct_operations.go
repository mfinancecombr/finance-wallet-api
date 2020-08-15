// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertTreasuryDirectOperation(d *wallet.TreasuryDirect) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertTreasuryDirectOperation")
	return m.insertOperation(d)
}

func (m *mongoSession) UpdateTreasuryDirectOperationByID(id string, d *wallet.TreasuryDirect) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateTreasuryDirectOperationByID")
	return m.updateTreasuryDirectByID(operationsCollection, id, d)
}

func (m *mongoSession) GetAllTreasuriesDirectsOperations() (wallet.TreasuryDirectList, error) {
	log.Debug("[DB] GetAllTreasuriesDirectsOperations")
	return m.getAllTreasuriesDirects(operationsCollection)
}

func (m *mongoSession) GetTreasuryDirectOperationsByPortfolioID(id string) (wallet.TreasuryDirectList, error) {
	log.Debug("[DB] GetTreasuryDirectOperationsByPortfolioID")
	return m.getTreasuryDirectByPortfolioID(operationsCollection, id)
}

func (m *mongoSession) GetTreasuryDirectOperationByID(id string) (*wallet.TreasuryDirect, error) {
	log.Debug("[DB] GetTreasuryDirectOperationByID")
	return m.getTreasuryDirectByID(operationsCollection, id)
}
