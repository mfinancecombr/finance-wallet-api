// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertTreasuryDirectOperation(d *wallet.TreasuryDirect) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertTreasuryDirectOperation")
	return m.insertOperation(d)
}

func (m *mongoSession) UpdateTreasuryDirectOperationByID(id string, d *wallet.TreasuryDirect) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateTreasuryDirectOperationByID")
	d.ID = ""
	return m.updateOperation(operationsCollection, id, d)
}

func (m *mongoSession) GetAllTreasuriesDirectsOperations() (wallet.TreasuryDirectList, error) {
	log.Debug("[DB] GetAllTreasuriesDirectsOperations")
	results, err := m.collection.FindAll(operationsCollection, bson.M{})
	if err != nil {
		return nil, err
	}
	treasuryDirectList := wallet.TreasuryDirectList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		treasuryDirect := wallet.TreasuryDirect{}
		bson.Unmarshal(bsonBytes, &treasuryDirect)
		treasuryDirectList = append(treasuryDirectList, treasuryDirect)
	}
	return treasuryDirectList, nil
}

func (m *mongoSession) GetTreasuryDirectOperationByID(id string) (*wallet.TreasuryDirect, error) {
	log.Debug("[DB] GetTreasuryDirectOperationByID")
	treasuryDirect := &wallet.TreasuryDirect{}
	if err := m.getOperationByID(operationsCollection, id, treasuryDirect); err != nil {
		return nil, err
	}
	if treasuryDirect.Symbol == "" {
		return nil, nil
	}
	return treasuryDirect, nil
}
