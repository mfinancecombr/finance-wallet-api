// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *mongoSession) getAllStocksFunds(c string) (wallet.StockFundList, error) {
	log.Debug("[DB] getAllStocksFunds")
	results, err := m.collection.FindAll(c, bson.M{})
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

func (m *mongoSession) getStockFundByPortfolioID(c, id string) (wallet.StockFundList, error) {
	log.Debug("[DB] getStockFundByPortfolioID")
	query := bson.M{"portfolioId": id, "itemType": "stocks-funds"}
	results, err := m.collection.FindAll(c, query)
	if err != nil {
		return nil, err
	}
	operationsList := wallet.StockFundList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		operation := wallet.StockFund{}
		bson.Unmarshal(bsonBytes, &operation)
		operationsList = append(operationsList, operation)
	}
	return operationsList, nil
}

func (m *mongoSession) getStockFundByID(c, id string) (*wallet.StockFund, error) {
	log.Debug("[DB] getStockFundByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": objectId}
	h := &wallet.StockFund{}
	err = m.collection.FindOne(c, query, h)
	if err != nil {
		return nil, err
	}
	if h.Symbol == "" {
		return nil, nil
	}
	return h, nil
}
