// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertFIIOperation(d *wallet.FII) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertFIIOperation")
	return m.insertOperation(d)
}

func (m *mongoSession) UpdateFIIOperationByID(id string, d *wallet.FII) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateFIIOperationByID")
	d.ID = ""
	return m.updateOperation(operationsCollection, id, d)
}

func (m *mongoSession) GetAllFIIsOperations() (wallet.FIIList, error) {
	log.Debug("[DB] GetAllFIIsOperations")
	return m.getAllFIIs(operationsCollection)
}

func (m *mongoSession) GetFIIOperationsByPortfolioID(id string) (wallet.FIIList, error) {
	log.Debug("[DB] GetFIIOperationsByPortfolioID")
	return m.getFIIByPortfolioID(operationsCollection, id)
}

func (m *mongoSession) GetFIIOperationByID(id string) (*wallet.FII, error) {
	log.Debug("[DB] GetFIIOperationByID")
	fii := &wallet.FII{}
	if err := m.getOperationByID(operationsCollection, id, fii); err != nil {
		return nil, err
	}
	if fii.Symbol == "" {
		return nil, nil
	}
	return fii, nil
}
