// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertStockFundOperation(d *wallet.StockFund) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertStockFundOperation")
	return m.insertOperation(d)
}

func (m *mongoSession) UpdateStockFundOperationByID(id string, d *wallet.StockFund) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateStockFundOperationByID")
	d.ID = ""
	return m.updateOperation(operationsCollection, id, d)
}

func (m *mongoSession) GetAllStocksFundsOperations() (wallet.StockFundList, error) {
	log.Debug("[DB] GetAllStocksFundsOperations")
	query := bson.M{"itemType": "stocks-funds"}
	results, err := m.collection.FindAll(operationsCollection, query)
	if err != nil {
		return nil, err
	}
	operationsList := wallet.StockFundList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		buy := wallet.StockFund{}
		bson.Unmarshal(bsonBytes, &buy)
		operationsList = append(operationsList, buy)
	}
	return operationsList, nil
}

func (m *mongoSession) GetStockFundOperationByID(id string) (*wallet.StockFund, error) {
	log.Debug("[DB] GetStockFundOperationByID")
	stockFund := &wallet.StockFund{}
	if err := m.getOperationByID(operationsCollection, id, stockFund); err != nil {
		return nil, err
	}
	if stockFund.Symbol == "" {
		return nil, nil
	}
	return stockFund, nil
}
