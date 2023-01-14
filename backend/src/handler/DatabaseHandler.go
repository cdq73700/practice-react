package handler

import (
	"backend/src/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://admin:admin@fiber-mongo:27017"

type DbStruct struct {
	model *models.MongoStrict
}

type DbInterface models.MongoInterface

func NewDatabase(dbName string, collName string) (*models.MongoStrict, error) {
	var db DbStruct = DbStruct{model: &models.MongoStrict{DbName: dbName, CollName: collName}}
	var dbi DbInterface = &db

	if err := Process(dbi,&db); err != nil {
		return nil, err
	}

	return db.model, nil
}

func Process(dbi DbInterface,db *DbStruct) error {
	var err error
	if err = dbi.ConnectionDatabase(); err != nil {
		return err
	}
	dbi.SetDatabase(db.model.DbName)
	dbi.SetCollection(db.model.CollName)
	return err
}

func (db *DbStruct) ConnectionDatabase() error {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return  err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	
	db.model.Client = client

	return nil
}

func (db *DbStruct) SetDatabase(dbName string) {
	if db.model.Client == nil {
		fmt.Println("not client")
	}
	db.model.Database = db.model.Client.Database(dbName)
}

func (db *DbStruct) SetCollection(collName string) {
	db.model.Collection = db.model.Database.Collection(collName)
}