// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var operationTypes = map[string]func() wallet.Tradable{
	wallet.ItemTypeStocks:                 func() wallet.Tradable { return wallet.NewStock() },
	wallet.ItemTypeFIIS:                   func() wallet.Tradable { return wallet.NewFII() },
	wallet.ItemTypeCertificateOfDeposit:   func() wallet.Tradable { return wallet.NewCertificateOfDeposit() },
	wallet.ItemTypeStocksTreasuriesDirect: func() wallet.Tradable { return wallet.NewTreasuryDirect() },
	wallet.ItemTypeStocksStocksFunds:      func() wallet.Tradable { return wallet.NewStock() },
	wallet.ItemTypeStocksFICFI:            func() wallet.Tradable { return wallet.NewFICFI() },
}

func (m *mongoSession) getOperationsSymbols(filter bson.M) ([]interface{}, error) {
	log.Debug("[DB] getOperationSymbols")
	return m.collection.Distinct(operationsCollection, "symbol", filter)
}

func (m *mongoSession) getItemTypes() ([]interface{}, error) {
	log.Debug("[DB] getItemTypes")
	return m.collection.Distinct(operationsCollection, "itemType", bson.M{})
}

func (m *mongoSession) getAllOperationsBySymbol(symbol, itemType string, year int) (wallet.OperationsList, error) {
	log.Debug("[DB] getAllOperationsBySymbol", "year", year)
	date := time.Date(year+1, 1, 1, 0, 0, 0, 0, time.UTC)
	query := bson.M{"symbol": symbol, "date": bson.M{"$lt": date}}
	opts := options.Find().SetSort(bson.D{{Key: "date", Value: 1}})
	results, err := m.collection.FindAll(operationsCollection, query, opts)
	if err != nil {
		return nil, err
	}
	operationsList := wallet.OperationsList{}
	for _, result := range results {
		newOperationType, ok := operationTypes[itemType]
		if !ok {
			errMsg := fmt.Sprintf("operation type)Item type '%s' not found", itemType)
			return nil, errors.New(errMsg)
		}
		operation := newOperationType()
		bsonBytes, err := bson.Marshal(result)
		if err != nil {
			return nil, err
		}
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
