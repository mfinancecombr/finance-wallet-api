// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) updateTreasuryDirectByID(c, id string, d *wallet.TreasuryDirect) (*mongo.UpdateResult, error) {
	log.Debug("[DB] updateTreasuryDirectByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	d.ID = ""
	f := bson.D{{"_id", objectId}}
	u := bson.D{{"$set", d}}
	return m.collection.UpdateOne(c, f, u)
}

func (m *mongoSession) getAllTreasuriesDirects(c string) (wallet.TreasuryDirectList, error) {
	log.Debug("[DB] getAllTreasuriesDirects")
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

func (m *mongoSession) getTreasuryDirectByPortfolioID(c, id string) (wallet.TreasuryDirectList, error) {
	log.Debug("[DB] getTreasuryDirectByPortfolioID")
	query := bson.M{"portfolioId": id, "itemType": "treasuries-direct"}
	results, err := m.collection.FindAll(c, query)
	if err != nil {
		return nil, err
	}
	operationsList := wallet.TreasuryDirectList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		operation := wallet.TreasuryDirect{}
		bson.Unmarshal(bsonBytes, &operation)
		operationsList = append(operationsList, operation)
	}
	return operationsList, nil
}

func (m *mongoSession) getTreasuryDirectByID(c, id string) (*wallet.TreasuryDirect, error) {
	log.Debug("[DB] getTreasuryDirectByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": objectId}
	h := &wallet.TreasuryDirect{}
	err = m.collection.FindOne(c, query, h)
	if err != nil {
		return nil, err
	}
	if h.Symbol == "" {
		return nil, nil
	}
	return h, nil
}
