package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB()(error,*mongo.Client){
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0-jfmqy.mongodb.net/test?retryWrites=true&w=majority")
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



