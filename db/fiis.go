// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongoSession) getAllFIIs(c string) (wallet.FIIList, error) {
	log.Debug("[DB] GetAllFIIsOperations")
	results, err := m.collection.FindAll(c, bson.M{})
	if err != nil {
		return nil, err
	}
	operationsList := wallet.FIIList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		buy := wallet.FII{}
		bson.Unmarshal(bsonBytes, &buy)
		operationsList = append(operationsList, buy)
	}
	return operationsList, nil
}

func (m *mongoSession) getFIIByPortfolioID(c, id string) (wallet.FIIList, error) {
	log.Debug("[DB] getFIIByPortfolioID")
	query := bson.M{"portfolioId": id, "itemType": "fiis"}
	results, err := m.collection.FindAll(c, query)
	if err != nil {
		return nil, err
	}
	operationsList := wallet.FIIList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		operation := wallet.FII{}
		bson.Unmarshal(bsonBytes, &operation)
		operationsList = append(operationsList, operation)
	}
	return operationsList, nil
}
