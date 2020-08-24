// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"context"
	"time"

	"github.com/mfinancecombr/finance-wallet-api/wallet"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// FIXME
const (
	brokersCollection    = "brokers"
	portfoliosCollection = "portfolios"
	operationsCollection = "operations"
)

type mongoSession struct {
	collection Collection
}

type DB interface {
	Create(d wallet.Queryable) (*mongo.InsertOneResult, error)
	Delete(collectionName, id string) (*mongo.DeleteResult, error)
	Get(id string, d wallet.Queryable) error
	GetAll(q wallet.Queryable) ([]wallet.Queryable, error)
	GetBySlug(slug string, d wallet.Queryable) error
	Update(id string, d wallet.Queryable) (*mongo.UpdateResult, error)

	GetPortfolioItems(p *wallet.Portfolio, year int) error
	GetAllOperations() (interface{}, error)
	GetAllPurchases() (interface{}, error)
	GetAllSales() (interface{}, error)

	Ping() error
}

func (m *mongoSession) GetBySlug(slug string, d wallet.Queryable) error {
	log.Debug("[DB] GetBySlug")
	query := bson.M{"slug": slug}
	return m.collection.FindOne(d.GetCollectionName(), query, d)
}

func (m *mongoSession) Get(id string, d wallet.Queryable) error {
	log.Debug("[DB] Get")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	query := bson.M{"_id": objectId}
	return m.collection.FindOne(d.GetCollectionName(), query, d)
}

func (m *mongoSession) GetAll(d wallet.Queryable) ([]wallet.Queryable, error) {
	log.Debug("[DB] GetAll")
	query := bson.M{}
	if d.GetItemType() != "" {
		query = bson.M{"itemType": d.GetItemType()}
	}
	results, err := m.collection.FindAll(d.GetCollectionName(), query)
	if err != nil {
		return nil, err
	}
	operationsList := []wallet.Queryable{}
	for _, result := range results {
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, d)
		operationsList = append(operationsList, d)
	}
	return operationsList, nil
}

func (m *mongoSession) Create(d wallet.Queryable) (*mongo.InsertOneResult, error) {
	log.Debug("[DB] Create")
	return m.collection.InsertOne(d.GetCollectionName(), d)
}

func (m *mongoSession) Update(id string, d wallet.Queryable) (*mongo.UpdateResult, error) {
	log.Debug("[DB] Update")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	dMarshal, _ := bson.Marshal(d)
	doc, err := bsonx.ReadDoc(dMarshal)
	doc = doc.Delete("_id")
	f := bson.D{{"_id", objectId}}
	u := bson.D{{"$set", doc}}
	return m.collection.UpdateOne(d.GetCollectionName(), f, u)
}

func (m *mongoSession) Delete(collectionName, id string) (*mongo.DeleteResult, error) {
	log.Debug("[DB] Delete")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	q := bson.M{"_id": objectId}
	return m.collection.DeleteOne(collectionName, q)
}

func newDBContext() (context.Context, context.CancelFunc) {
	log.Debug("[DB] New DB context")
	timeout := viper.GetDuration("db.operation.timeout")
	return context.WithTimeout(context.Background(), timeout*time.Second)
}

func NewMongoSession() (DB, error) {
	log.Debug("[DB] New mongo session")
	dbURI := viper.GetString("mongodb.endpoint")
	dbName := viper.GetString("mongodb.name")
	ctx, _ := newDBContext()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Errorf("[DB] Error on create mongo session: %s", err)
	}
	mongo := &mongoSession{
		collection: &mongoCollection{
			session: client,
			dbName:  dbName,
		},
	}
	return mongo, err
}

func (m *mongoSession) Ping() error {
	log.Debug("[DB] Ping")
	return m.collection.Ping()
}
