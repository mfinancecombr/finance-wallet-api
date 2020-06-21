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

func (m *mongoSession) insertSale(d interface{}) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] insertSale")
	return m.collection.InsertOne(salesCollection, d)
}

func (m *mongoSession) DeleteSaleByID(id string) (*mongo.DeleteResult, error) {
	log.Debug("[DB] DeleteSaleByID")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": objectId}
	return m.collection.DeleteOne(salesCollection, q)
}

func (m *mongoSession) GetSalesItemType() ([]interface{}, error) {
	log.Debug("[DB] GetSalesItemType")
	return m.collection.Distinct(salesCollection, "itemType", bson.M{})
}

func (m *mongoSession) getSalesSymbols(filter bson.M) ([]interface{}, error) {
	log.Debug("[DB] getSalesSymbols")
	return m.collection.Distinct(salesCollection, "symbol", filter)
}

func (m *mongoSession) getAllSalesBySymbol(symbol, itemType string, year, month int) (wallet.SalesList, error) {
	log.Debug("[DB] getAllSalesBySymbol")
	date := time.Date(year, time.Month(month), 31, 23, 59, 59, 0, time.UTC)
	query := bson.M{"symbol": symbol, "date": bson.M{"$lte": date}}
	results, err := m.collection.FindAll(salesCollection, query)
	if err != nil {
		return nil, err
	}
	//FIXME: duplicated
	salesList := wallet.SalesList{}
	for _, result := range results {
		var sale interface{}
		switch itemType {
		case "stocks":
			sale = &wallet.Stock{}
		case "fiis":
			sale = &wallet.FII{}
		case "certificates-of-deposit":
			sale = &wallet.CertificateOfDeposit{}
		case "treasuries-direct":
			sale = &wallet.TreasuryDirect{}
		case "stocks-funds":
			sale = &wallet.StockFund{}
		case "ficfi":
			sale = &wallet.FICFI{}
		default:
			log.Errorf("Item type '%s' not found", itemType)
		}
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, sale)
		salesList = append(salesList, sale)
	}
	return salesList, nil
}

func (m *mongoSession) GetAllSales() (interface{}, error) {
	log.Debug("[DB] GetAllSales")
	return m.collection.FindAll(salesCollection, bson.M{})
}
