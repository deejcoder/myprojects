package storage

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect connects to a MongoDB database server
func Connect(ctx context.Context, host string, port int, database string, user string, password string) {

	uri := fmt.Sprintf("mongodb://%s:%d", host, port)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Connection to the database has failed", err)
	}

	client.Database(database)

}
