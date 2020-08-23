// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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
