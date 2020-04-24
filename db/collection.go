// Copyright (c) 2020, Marcelo Jorge Vieira (https://github.com/mfinancecombr)
// Licensed under the BSD 3-Clause License

package db

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Collection interface {
	DeleteOne(c string, d interface{}) (*mongo.DeleteResult, error)
	Distinct(c string, q string, f interface{}) ([]interface{}, error)
	FindAll(c string, q bson.M) ([]bson.M, error)
	FindOne(c string, q bson.M, r interface{}) error
	InsertOne(c string, d interface{}) (*mongo.InsertOneResult, error)
	Ping() error
	UpdateOne(c string, f, d interface{}) (*mongo.UpdateResult, error)
}

type mongoCollection struct {
	dbName  string
	session *mongo.Client
}

func newCollectionContext() (context.Context, context.CancelFunc) {
	log.Debug("[Collection] New collection context")
	timeout := viper.GetDuration("collection.operation.timeout")
	return context.WithTimeout(context.Background(), timeout*time.Second)
}

func (m *mongoCollection) Ping() error {
	log.Debug("[Collection] Ping")
	ctx, _ := newCollectionContext()
	return m.session.Ping(ctx, readpref.Primary())
}

func (m *mongoCollection) InsertOne(c string, d interface{}) (*mongo.InsertOneResult, error) {
	log.Debug("[Collection] InsertOne")
	collection := m.session.Database(m.dbName).Collection(c)
	ctx, _ := newCollectionContext()
	return collection.InsertOne(ctx, d)
}

func (m *mongoCollection) FindAll(c string, q bson.M) ([]bson.M, error) {
	log.Debug("[Collection] FindAll")
	collection := m.session.Database(m.dbName).Collection(c)
	ctx, _ := newCollectionContext()
	cur, err := collection.Find(ctx, q)
	if err != nil {
		log.Errorf("[Collection] Find: %s", err)
		return nil, err
	}
	var results []bson.M
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		if err := cur.Decode(&result); err != nil {
			log.Errorf("[Collection] Decode: %s", err)
			return nil, err
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Errorf("[Collection] Cursor: %s", err)
		return nil, err
	}
	return results, nil
}

func (m *mongoCollection) FindOne(c string, q bson.M, r interface{}) error {
	log.Debug("[Collection] FindOne...")
	collection := m.session.Database(m.dbName).Collection(c)
	ctx, _ := newCollectionContext()
	err := collection.FindOne(ctx, q).Decode(r)
	if err == mongo.ErrNoDocuments {
		return nil
	}
	if err != nil {
		log.Errorf("[Collection] Error on retrieve data: %s", err)
		return err
	}
	return nil
}

func (m *mongoCollection) DeleteOne(c string, d interface{}) (*mongo.DeleteResult, error) {
	log.Debug("[Collection] DeleteOne")
	collection := m.session.Database(m.dbName).Collection(c)
	ctx, _ := newCollectionContext()
	return collection.DeleteOne(ctx, d)
}

func (m *mongoCollection) UpdateOne(c string, f, d interface{}) (*mongo.UpdateResult, error) {
	log.Debug("[Collection] UpdateOne")
	collection := m.session.Database(m.dbName).Collection(c)
	ctx, _ := newCollectionContext()
	return collection.UpdateOne(ctx, f, d)
}

func (m *mongoCollection) Distinct(c string, q string, f interface{}) ([]interface{}, error) {
	log.Debug("[Collection] Distinct")
	collection := m.session.Database(m.dbName).Collection(c)
	ctx, _ := newCollectionContext()
	return collection.Distinct(ctx, q, f)
}
