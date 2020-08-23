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
