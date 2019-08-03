package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect()(error,*mongo.Client){
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:<password>@golang-jfmqy.mongodb.net/test?retryWrites=true&w=majority")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err,nil
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err,nil
	}
	return nil,client
}

func GetDatabaseConnection()(error,*mongo.Database){
	err, cliente := Connect()
	if err != nil {
		return err,nil
	}
	DB := cliente.Database("golang")
	return nil, DB
}

func SingleDBMongo()(error,*mongo.Database){
	if DB != nil {
		err,_ := GetDatabaseConnection()
		if err != nil {
			return err, nil
		}
	}
	return nil,DB
}

