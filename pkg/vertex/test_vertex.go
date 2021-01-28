package vertex

import (
  "testing"
  "math/rand"
  "fmt"
  "time"
)

func TestDataVertex(t *testing.T) {
  var configFilePath string = "./app-config.yml"
  var dbConfigFilePath string = "./database/db-config.yaml"

  config, configErr :=  getConfig(configFilePath)
  if configErr != nil {
    t.Errorf("Application config reading error: %v", configErr)

  }
  dbConfig, dbErr := getDBConfig(dbConfigFilePath)
  if dbErr != nil {
    t.Errorf("Database config readig error: %v", dbErr)
  }

  db, dbErr := CreateDatabase(dbConfig);
  if (dbErr != nil) {
    t.Errorf("Error creating the resolvers", resolverErr);
  }

  dataloader, dataloaderErr := CreateDataloaders(db);
  if (resolverErr != nil) {
    t.Errorf("Error creating the resolvers", dataloaderErr);
  }

  resolvers, resolverErr := CreateResolvers(db);
  if (resolverErr != nil) {
    t.Errorf("Error creating the resolvers", resolverErr);
  }

  applicationVertex, err := CreateDataVertex(config, dbConfig)
  if err != nil {
    t.Errorf("Data vertex creation error: %v", err)
  }

  _, queryErr := applicationVertex.query();
  if (queryErr != nil) {
    t.Errorf("Error creating the resolvers", resolverErr);
  }
}
