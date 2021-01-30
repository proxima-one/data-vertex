package vertex

import (
  "testing"
  _ "math/rand"
  _ "fmt"
  _ "time"
  proxima "github.com/proxima-one/proxima-db-client-go/pkg/database"
)

func TestDataVertex(t *testing.T) {
  var configFilePath string = "../../app-config.yml"
  var dbConfigFilePath string = "../database/db-config.yaml"

  config, configErr :=  getConfig(configFilePath)

  if configErr != nil {
    t.Errorf("Application config reading error: %v", configErr)

  }
  dbConfig, dbConfigErr := getDBConfig(dbConfigFilePath)
  if dbConfigErr != nil {
    t.Errorf("Database config readig error: %v", dbConfigErr)
  }
  var db *proxima.ProximaDatabase;

  // db, dbErr := CreateApplicationDatabase(dbConfig);
  // if (dbErr != nil) {
  //   t.Error("Error creating the resolvers", dbErr);
  // }

  _, dataloaderErr := CreateDataloaders(db);
  if (dataloaderErr != nil) {
    t.Error("Error creating the resolvers", dataloaderErr);
  }

  _, resolverErr := CreateResolvers(db);
  if (resolverErr != nil) {
    t.Error("Error creating the resolvers", resolverErr);
  }

  applicationVertex, err := CreateDataVertex(config, dbConfig)
  if err != nil {
    t.Errorf("Data vertex creation error: %v", err)
  }

   query := applicationVertex.query();
  if query == nil {
    t.Error("Error creating the query");
  }
}

func TestLoadDataVertex(t *testing.T) {
  //homeDir
  var configFilePath string = "../app-config.yml"
  var dbConfigFilePath string = "../database/db-config.yaml"

  applicationVertex, err := LoadDataVertex(configFilePath, dbConfigFilePath)
  if err != nil {
    t.Errorf("Data vertex creation error: %v", err)
  }

  query := applicationVertex.query();
  if (query == nil) {
    t.Error("Error creating the resolvers", query);
  }
}
