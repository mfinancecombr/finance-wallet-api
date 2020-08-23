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
	d.ID = ""
	return m.updateOperation(operationsCollection, id, d)
}

func (m *mongoSession) GetAllFICFIOperations() (wallet.FICFIList, error) {
	log.Debug("[DB] GetAllFICFIOperations")
	return m.getAllFICFI(operationsCollection)
}

func (m *mongoSession) GetFICFIOperationByID(id string) (*wallet.FICFI, error) {
	log.Debug("[DB] GetFICFIOperationByID")
	ficfi := &wallet.FICFI{}
	if err := m.getOperationByID(operationsCollection, id, ficfi); err != nil {
		return nil, err
	}
	if ficfi.Symbol == "" {
		return nil, nil
	}
	return ficfi, nil
}
