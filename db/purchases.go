// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"time"

	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) insertPurchase(d interface{}) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] insertPurchase")
	return m.collection.InsertOne(purchasesCollection, d)
}

func (m *mongoSession) DeletePurchaseByID(id string) (*mongo.DeleteResult, error) {
	log.Debug("[DB] DeleteBuyByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": objectId}
	return m.collection.DeleteOne(purchasesCollection, q)
}

func (m *mongoSession) getPurchasesSymbols(filter bson.M) ([]interface{}, error) {
	log.Debug("[DB] getPurchaseSymbols")
	return m.collection.Distinct(purchasesCollection, "symbol", filter)
}

func (m *mongoSession) getAllPurchasesBySymbol(symbol, itemType string, year int) (wallet.PurchasesList, error) {
	log.Debug("[DB] getAllPurchasesBySymbol")
	date := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	query := bson.M{"symbol": symbol, "date": bson.M{"$lte": date}}
	results, err := m.collection.FindAll(purchasesCollection, query)
	if err != nil {
		return nil, err
	}
	// FIXME
	purchasesList := wallet.PurchasesList{}
	for _, result := range results {
		var purchase interface{}
		switch itemType {
		case "stocks":
			purchase = &wallet.Stock{}
		case "fiis":
			purchase = &wallet.FII{}
		case "certificates-of-deposit":
			purchase = &wallet.CertificateOfDeposit{}
		case "treasuries-direct":
			purchase = &wallet.TreasuryDirect{}
		case "stocks-funds":
			purchase = &wallet.StockFund{}
		case "ficfi":
			purchase = &wallet.FICFI{}
		default:
			log.Errorf("Item type '%s' not found", itemType)
		}
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, purchase)
		purchasesList = append(purchasesList, purchase)
	}
	return purchasesList, nil
}

func (m *mongoSession) GetAllPurchases() (interface{}, error) {
	log.Debug("[DB] GetAllPurchases")
	return m.collection.FindAll(purchasesCollection, bson.M{})
}
