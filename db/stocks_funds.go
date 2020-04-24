// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) updateStockFundByID(c, id string, d *wallet.StockFund) (*mongo.UpdateResult, error) {
	log.Debug("[DB] updateStockFundByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	d.ID = ""
	f := bson.D{{"_id", objectId}}
	u := bson.D{{"$set", d}}
	return m.collection.UpdateOne(c, f, u)
}

func (m *mongoSession) getAllStocksFunds(c string) (wallet.StockFundList, error) {
	log.Debug("[DB] getAllStocksFunds")
	results, err := m.collection.FindAll(c, bson.M{})
	if err != nil {
		return nil, err
	}
	salesList := wallet.StockFundList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		buy := wallet.StockFund{}
		bson.Unmarshal(bsonBytes, &buy)
		salesList = append(salesList, buy)
	}
	return salesList, nil
}

func (m *mongoSession) getStockFundByPortfolioID(c, id string) (wallet.StockFundList, error) {
	log.Debug("[DB] getStockFundByPortfolioID")
	query := bson.M{"portfolioId": id, "itemType": "stocks-funds"}
	results, err := m.collection.FindAll(c, query)
	if err != nil {
		return nil, err
	}
	salesList := wallet.StockFundList{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		sale := wallet.StockFund{}
		bson.Unmarshal(bsonBytes, &sale)
		salesList = append(salesList, sale)
	}
	return salesList, nil
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
