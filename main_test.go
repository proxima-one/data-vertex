package proxima_db_client_go

import (
  "testing"
  "math/rand"
  "fmt"
  "time"
)

func NewDatabaseClient() (*ProximaDB, error) {
  ip := "0.0.0.0"
  port := "50051"
  return DefaultProximaServiceClient(ip, port)
}


var appName string = "NewApplication";
var proximaClient *ProximaDatabase = NewDatabaseClient();
var tables map[string][]string = TableSetup();

func TableSetup() (map[string][]string) {
  return map[string]{"remove": remove, "full": full, "partial": partial}
}


  func TestDataVertex(t *testing.T) {
    resolvers, resolverErr := CreateResolvers();
    if (resolverErr != nil) {
      t.Error("Error creating the resolvers", resolverErr);
    }
    db, dbErr := CreateDatabase();
    if (dbErr != nil) {
      t.Error("Error creating the resolvers", resolverErr);
    }

    vertex, vertexErr := CreateDataVertex();
    if (vertexErr != nil) {
      t.Error("Error creating the resolvers", resolverErr);
    }
    _, queryErr := vertex.query();
    if (queryErr != nil) {
      t.Error("Error creating the resolvers", resolverErr);
    }
  }
