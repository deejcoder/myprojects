package storage

import (
	"context"
	"fmt"

	"github.com/Dilicor/myprojects/config"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect connects to a MongoDB database server
func Connect(ctx context.Context) {

	cfg := config.GetConfig()

	log.Infof("Connecting to the database using host=%s:%d", cfg.Database.Host, cfg.Database.Port)
	uri := fmt.Sprintf("mongodb://%s:%d", cfg.Database.Host, cfg.Database.Port)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Connection to the database has failed", err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Connection to the database has failed", err)
	} else {
		log.Info("Successfully connected to the database.")
	}

	client.Database(cfg.Database.Database)

	log.Info(client.ListDatabases(context.Background()))

	done := make(chan struct{})
	go func() {
		<-ctx.Done()
		if err := client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
		close(done)
	}()

	<-done
}
