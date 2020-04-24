// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the BSD 3-Clause License

package db

import (
	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *mongoSession) InsertBroker(d interface{}) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] InsertBroker")
	return m.collection.InsertOne(brokersCollection, d)
}

func (m *mongoSession) DeleteBrokerByID(id string) (*mongo.DeleteResult, error) {
	log.Debug("[DB] DeletePortfolioByID")
	q := bson.M{"_id": id}
	return m.collection.DeleteOne(brokersCollection, q)
}

func (m *mongoSession) UpdateBroker(id string, d interface{}) (*mongo.UpdateResult, error) {
	log.Debug("[DB] UpdateBroker")
	f := bson.D{{"_id", id}}
	u := bson.D{{"$set", d}}
	return m.collection.UpdateOne(brokersCollection, f, u)
}

func (m *mongoSession) GetAllBrokers() (*wallet.BrokersList, error) {
	log.Debug("[DB] GetAllBrokers")
	results, err := m.collection.FindAll(brokersCollection, bson.M{})
	if err != nil {
		return nil, err
	}
	brokersTemp := &wallet.BrokersList{Brokers: []wallet.Broker{}}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		broker := wallet.Broker{}
		bson.Unmarshal(bsonBytes, &broker)
		brokersTemp.Brokers = append(brokersTemp.Brokers, broker)
	}
	return brokersTemp, nil
}

func (m *mongoSession) GetBrokerByID(id string) (*wallet.Broker, error) {
	log.Debug("[DB] GetBrokerByID")
	query := bson.M{"_id": id}
	h := &wallet.Broker{}
	err := m.collection.FindOne(brokersCollection, query, h)
	if err != nil {
		return nil, err
	}
	if h.Name == "" {
		return nil, nil
	}
	return h, nil
}
