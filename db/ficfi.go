// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) updateFICFIByID(c, id string, d *wallet.FICFI) (*mongo.UpdateResult, error) {
	log.Debug("[DB] updateFICFIByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	d.ID = ""
	f := bson.D{{"_id", objectId}}
	u := bson.D{{"$set", d}}
	return m.collection.UpdateOne(c, f, u)
}

func (m *mongoSession) getAllFICFI(c string) (wallet.FICFIList, error) {
	log.Debug("[DB] getAllFICFI")
	results, err := m.collection.FindAll(c, bson.M{})
	if err != nil {
		return nil, err
	}
	operationsList := wallet.FICFIList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		buy := wallet.FICFI{}
		bson.Unmarshal(bsonBytes, &buy)
		operationsList = append(operationsList, buy)
	}
	return operationsList, nil
}

func (m *mongoSession) getFICFIByPortfolioID(c, id string) (wallet.FICFIList, error) {
	log.Debug("[DB] getFICFIByPortfolioID")
	query := bson.M{"portfolioId": id, "itemType": "ficfi"}
	results, err := m.collection.FindAll(c, query)
	if err != nil {
		return nil, err
	}
	operationsList := wallet.FICFIList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		operation := wallet.FICFI{}
		bson.Unmarshal(bsonBytes, &operation)
		operationsList = append(operationsList, operation)
	}
	return operationsList, nil
}

func (m *mongoSession) getFICFIByID(c, id string) (*wallet.FICFI, error) {
	log.Debug("[DB] getFICFIByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": objectId}
	h := &wallet.FICFI{}
	err = m.collection.FindOne(c, query, h)
	if err != nil {
		return nil, err
	}
	if h.Symbol == "" {
		return nil, nil
	}
	return h, nil
}
