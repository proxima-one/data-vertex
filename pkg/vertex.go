package main

import (
	_ "os"
	proxima "github.com/proxima-one/proxima-db-client-go"
	_ "github.com/proxima-one/proxima-data-vertex/pkg/resolvers"
	dataloader "github.com/proxima-one/proxima-data-vertex/pkg/dataloaders"
	"github.com/99designs/gqlgen/handler"
	gql "github.com/proxima-one/proxima-data-vertex/pkg/gql" //gql
)

//Structure and schema of the vertex
type ProximaDataVertex struct {
  name *string //map of tables
  id *string
  version *string
  applicationDB *ProximaDB
	executableSchema *gql.Exec
}

func (vertex *ProximaDataVertex) query() gin.HandlerFunc {
	h := handler.GraphQL(vertex.executableSchema)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (vertex *ProximaDataVertex) startVertexServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	go r.POST("/query", vertex.query())
	go r.GET("/", vertex.playgroundHandler())
	r.Run(":4000")
}

func (vertex *ProximaDataVertex) playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func CreateDataVertex(config map[string]interface{}) (ProximaDataVertex, error) {
	database, _ := CreateDatabase(config)
	resolvers, _ := CreateResolvers(database)
	//executableSchema
	 c := Config{
		Resolvers: &resolvers
		}

	 return ProximaDataVertex{name: config['name'], id: config.id , version: config.version, applicationDB: database, executableSchema: c}, nil
}


//Database of the vertex...
func CreateDatabase(db_config) {
	//load, then init
	//database creation
	//database start()
	//return database
	ip := getEnv("DB_ADDRESS" , "0.0.0.0")
	port :=  getEnv("DB_PORT", "50051")
	proximaDB := proxima.NewProximaDB(ip, port)
	_, err := proximaDB.OpenAll(tableList)
	if err != nil {
		return proximaDB, err
	}
	return proximaDB, nil
}

//executableSchema, dataloader is needed as well
func CreateResolvers(db *proxima.ProximaDB) (gql.Config) {
	loader, _  := CreateDataloaders(db)
	return resolver.NewResolver(loader, db)
}

func CreateDataloaders(db *proxima.ProximaDB) (*dataloader.Dataloader, error) {
  loader , err:= dataloader.NewDataloader(db)
  if err != nil {
    return nil, err
  }
  return loader, nil
}
