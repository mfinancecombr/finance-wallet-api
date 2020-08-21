// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *mongoSession) getAllStocks(c string) (wallet.StockList, error) {
	log.Debug("[DB] getAllStocks")
	results, err := m.collection.FindAll(c, bson.M{})
	if err != nil {
		return nil, err
	}
	operationsList := wallet.StockList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		stock := wallet.Stock{}
		bson.Unmarshal(bsonBytes, &stock)
		operationsList = append(operationsList, stock)
	}
	return operationsList, nil
}

func (m *mongoSession) getStockByPortfolioID(c, id string) (wallet.StockList, error) {
	log.Debug("[DB] getStockByPortfolioID")
	query := bson.M{"portfolioId": id, "itemType": "stocks"}
	results, err := m.collection.FindAll(c, query)
	if err != nil {
		return nil, err
	}
	operationsList := wallet.StockList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		stock := wallet.Stock{}
		bson.Unmarshal(bsonBytes, &stock)
		operationsList = append(operationsList, stock)
	}
	return operationsList, nil
}

func (m *mongoSession) getStockByID(c, id string) (*wallet.Stock, error) {
	log.Debug("[DB] getStockByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": objectId}
	h := &wallet.Stock{}
	err = m.collection.FindOne(c, query, h)
	if err != nil {
		return nil, err
	}
	if h.Symbol == "" {
		return nil, nil
	}
	return h, nil
}
