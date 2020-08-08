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

func (m *mongoSession) updateCertificateOfDepositByID(c, id string, d *wallet.CertificateOfDeposit) (*mongo.UpdateResult, error) {
	log.Debug("[DB] updateCertificateOfDepositByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	d.ID = ""
	f := bson.D{{"_id", objectId}}
	u := bson.D{{"$set", d}}
	return m.collection.UpdateOne(c, f, u)
}

func (m *mongoSession) getAllCertificatesOfDeposits(c string) (wallet.CertificateOfDepositList, error) {
	log.Debug("[DB] getAllCertificatesOfDeposits")
	results, err := m.collection.FindAll(c, bson.M{})
	if err != nil {
		return nil, err
	}
	purchasesList := wallet.CertificateOfDepositList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		buy := wallet.CertificateOfDeposit{}
		bson.Unmarshal(bsonBytes, &buy)
		purchasesList = append(purchasesList, buy)
	}
	return purchasesList, nil
}

func (m *mongoSession) getCertificateOfDepositByPortfolioID(c, id string) (wallet.CertificateOfDepositList, error) {
	log.Debug("[DB] getCertificateOfDepositByPortfolioID")
	query := bson.M{"portfolioId": id, "itemType": "certificates-of-deposit"}
	results, err := m.collection.FindAll(c, query)
	if err != nil {
		return nil, err
	}
	purchasesList := wallet.CertificateOfDepositList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		purchase := wallet.CertificateOfDeposit{}
		bson.Unmarshal(bsonBytes, &purchase)
		purchasesList = append(purchasesList, purchase)
	}
	return purchasesList, nil
}

func (m *mongoSession) getCertificateOfDepositByID(c, id string) (*wallet.CertificateOfDeposit, error) {
	log.Debug("[DB] getCertificateOfDepositByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": objectId}
	h := &wallet.CertificateOfDeposit{}
	err = m.collection.FindOne(c, query, h)
	if err != nil {
		return nil, err
	}
	if h.Symbol == "" {
		return nil, nil
	}
	return h, nil
}