/*
Data Vertex
- Database creation
  - Tables
- DApp aggregator
- Dataloaders and resolvers (config)
*/

package main

import (
	"os"
	proxima "github.com/proxima-one/proxima-db-client-go"
	resolver "github.com/proxima-one/binance-chain-subgraph/pkg/resolvers"
	dataloader "github.com/proxima-one/binance-chain-subgraph/pkg/dataloader"
  gql "github.com/proxima-one/binance-chain-subgraph/pkg/gql"
	"github.com/99designs/gqlgen/handler"
)

//Structure and schema of the vertex
type ProximaDataVertex struct {
  name *string //map of tables
  id *string
  version *string
  applicationDB *ProximaDB
	executableSchema *gql.Exec
}

func query() gin.HandlerFunc {
	h := handler.GraphQL(vertex.executableSchema)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func CreateDataVertex() {
	resolvers, _ := CreateResolvers(config)
	//name, id
	 c := Config{
		Resolvers: &resolvers
		}
	 return ProximaDataVertex{name: , id: , version: , applicationDB: , executableSchema: c}
}

//executableSchema, dataloader is needed as well
func CreateResolvers(db *proxima.ProximaDB) (gql.Config) {
	loader, _  := CreateDataloaders(config)
	return resolver.NewResolver(loader)
}

func CreateDataloaders(db *proxima.ProximaDB) (*dataloader.Dataloader, error) {
  loader , err:= dataloader.NewDataloader(db)
  if err != nil {
    return nil, err
  }
  return loader, nil
}

//Database of the vertex...
func CreateDatabase(config) {

}

//new executableSchema
func InitClient(ip, port string) (*proxima.ProximaDB, error) {

}

func StartDatabase(tableList []string) (*proxima.ProximaDB, error) {
	ip := getEnv("DB_ADDRESS" , "0.0.0.0")
	port :=  getEnv("DB_PORT", "50051")
	proximaDB := proxima.NewProximaDB(ip, port)
	_, err := proximaDB.OpenAll(tableList)
	if err != nil {
		return proximaDB, err
	}
	return proximaDB, nil
}
