package storage

import (
	"context"
	"fmt"

	"github.com/deejcoder/myprojects/config"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect connects to a MongoDB database server
func Connect() *mongo.Database {

	// get db info from config
	cfg := config.GetConfig()

	// connect to mongo using db info
	log.Infof("Connecting to the database using host=%s:%d", cfg.Database.Host, cfg.Database.Port)
	uri := fmt.Sprintf("mongodb://%s:%d", cfg.Database.Host, cfg.Database.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// try ping the db server
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connection to the database was successfully established")
	db := client.Database(cfg.Database.Database)
	return db
}
