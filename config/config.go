package config

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ConnectNeo4j(ctx context.Context) neo4j.DriverWithContext {
	uri := "bolt://localhost:7687"
	username := "neo4j"
	password := "password"

	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Fatal("Failed to create Neo4j driver: ", err)
	}

	err = pingDatabase(ctx, driver)
	if err != nil {
		log.Fatal("Failed to connect DB!", err)
	} else {
		log.Println("Connected to the DB!")
	}

	return driver
}

func pingDatabase(ctx context.Context, driver neo4j.DriverWithContext) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	_, err := session.Run(ctx, "RETURN 1", nil)
	return err
}
