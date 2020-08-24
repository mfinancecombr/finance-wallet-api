// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"time"

	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongoSession) getOperationsSymbols(filter bson.M) ([]interface{}, error) {
	log.Debug("[DB] getOperationSymbols")
	return m.collection.Distinct(operationsCollection, "symbol", filter)
}

func (m *mongoSession) getAllOperationsBySymbol(symbol, itemType string, year int) (wallet.OperationsList, error) {
	log.Debug("[DB] getAllOperationsBySymbol")
	date := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	query := bson.M{"symbol": symbol, "date": bson.M{"$lte": date}}
	opts := options.Find().SetSort(bson.D{{"date", 1}})
	results, err := m.collection.FindAll(operationsCollection, query, opts)
	if err != nil {
		return nil, err
	}
	// FIXME
	operationsList := wallet.OperationsList{}
	for _, result := range results {
		var operation wallet.Tradable
		switch itemType {
		case "stocks":
			operation = &wallet.Stock{}
		case "fiis":
			operation = &wallet.FII{}
		case "certificates-of-deposit":
			operation = &wallet.CertificateOfDeposit{}
		case "treasuries-direct":
			operation = &wallet.TreasuryDirect{}
		case "stocks-funds":
			operation = &wallet.StockFund{}
		case "ficfi":
			operation = &wallet.FICFI{}
		default:
			log.Errorf("Item type '%s' not found", itemType)
		}
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, operation)
		operationsList = append(operationsList, operation)
	}
	return operationsList, nil
}

func (m *mongoSession) GetAllOperations() (interface{}, error) {
	log.Debug("[DB] GetAllOperations")
	return m.collection.FindAll(operationsCollection, bson.M{})
}

func (m *mongoSession) GetAllPurchases() (interface{}, error) {
	log.Debug("[DB] GetAllOperations")
	query := bson.M{"type": "purchase"}
	opts := options.Find().SetSort(bson.D{{"date", -1}})
	return m.collection.FindAll(operationsCollection, query, opts)
}

func (m *mongoSession) GetAllSales() (interface{}, error) {
	log.Debug("[DB] GetAllOperations")
	query := bson.M{"type": "sale"}
	opts := options.Find().SetSort(bson.D{{"date", -1}})
	return m.collection.FindAll(operationsCollection, query, opts)
}
