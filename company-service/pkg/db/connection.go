package db

import (
	"company-service/pkg/config"
	"context"
	"fmt"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase(cfg config.Config) (*mongo.Database, error) {

	// to user password with special character, update the password to escape query
	updatedPass := url.QueryEscape(cfg.DBPassword)

	// add db protocol
	dbURI := fmt.Sprintf("%s://%s:%s@%s/?%s", cfg.DBProtocol, cfg.DBUser, updatedPass, cfg.DBHost, cfg.DBOptions)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(dbURI).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo atlas: %v", err)
	}

	return client.Database(cfg.DBName), nil

}
