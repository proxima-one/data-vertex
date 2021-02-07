package vertex

import (
	_ "fmt"
	proxima "github.com/proxima-one/proxima-db-client-go/pkg/database"
	_ "math/rand"
	"testing"
	_ "time"
)

func TestDataVertex(t *testing.T) {
	var configFilePath string = "../../app-config.yml"
	var dbConfigFilePath string = "../database/db-config.yaml"

	config, configErr := getConfig(configFilePath)

	if configErr != nil {
		t.Errorf("Application config reading error: %v", configErr)

	}
	dbConfig, dbConfigErr := getDBConfig(dbConfigFilePath)
	if dbConfigErr != nil {
		t.Errorf("Database config readig error: %v", dbConfigErr)
	}
	var db *proxima.ProximaDatabase

	db, dbErr := CreateApplicationDatabase(dbConfig);
	if (dbErr != nil) {
	  t.Error("Error creating the database", dbErr);
	}

	_, dataloaderErr := CreateDataloaders(db)
	if dataloaderErr != nil {
		t.Error("Error creating the resolvers", dataloaderErr)
	}

	_, resolverErr := CreateResolvers(db)
	if resolverErr != nil {
		t.Error("Error creating the resolvers", resolverErr)
	}

	applicationVertex, err := CreateDataVertex(config, dbConfig)
	if err != nil {
		t.Errorf("Data vertex creation error: %v", err)
	}

	query := applicationVertex.query()
	if query == nil {
		t.Error("Error creating the query")
	}
}

func TestLoadDataVertex(t *testing.T) {
	//homeDir
	var configFilePath string = "../../app-config.yml"
	var dbConfigFilePath string = "../database/db-config.yaml"

	applicationVertex, err := LoadDataVertex(configFilePath, dbConfigFilePath)
	if err != nil {
		t.Errorf("Data vertex creation error: %v", err)
	}

	query := applicationVertex.query()
	if query == nil {
		t.Error("Error creating the resolvers", query)
	}
}

func TestDataVertexResolvers(t *testing.T) {
	//homeDir
	var configFilePath string = "../../app-config.yml"
	var dbConfigFilePath string = "../database/db-config.yaml"
	var testConfigFilePath string = "../test/entity_test.json"

	applicationVertex, err := LoadDataVertex(configFilePath, dbConfigFilePath)
	if err != nil {
		t.Errorf("Data vertex creation error: %v", err)
	}
	entityTests, entityErr := GenerateTestEntities(applicationVertex, testConfigFilePath)
	if entityErr != nil || entityTests == nil {
		t.Error("Error creating the resolver tests", entityErr)
	}

  // for name, entityTest := range entityTests {
  //   go entityTest.generate(100)
  // }

	for name, entityTest := range entityTests {
		go entityTest.runTests(t, 100)
	}
}
