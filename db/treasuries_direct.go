// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

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
