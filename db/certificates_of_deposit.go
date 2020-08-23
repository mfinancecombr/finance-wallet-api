// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongoSession) getAllCertificatesOfDeposits(c string) (wallet.CertificateOfDepositList, error) {
	log.Debug("[DB] getAllCertificatesOfDeposits")
	results, err := m.collection.FindAll(c, bson.M{})
	if err != nil {
		return nil, err
	}
	operationsList := wallet.CertificateOfDepositList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		buy := wallet.CertificateOfDeposit{}
		bson.Unmarshal(bsonBytes, &buy)
		operationsList = append(operationsList, buy)
	}
	return operationsList, nil
}
