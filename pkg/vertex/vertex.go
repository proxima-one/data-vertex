package vertex

import (
	proxima "github.com/proxima-one/proxima-db-client-go"
	resolver "github.com/proxima-one/proxima-data-vertex/pkg/resolvers"
	dataloader "github.com/proxima-one/proxima-data-vertex/pkg/dataloaders"
	"github.com/99designs/gqlgen/handler"
	gql "github.com/proxima-one/proxima-data-vertex/pkg/gql"
)

type ProximaDataVertex struct {
  name *string
  id *string
  version *string
  applicationDB *ProximaDB
	executableSchema *gql.Exec
}

func CreateDataVertex(config, dbConfig map[string]interface{}) (*ProximaDataVertex, error) {
	database, dErr := CreateApplicationDatabase(dbConfig)
	if dErr != nil {
		return nil, dErr
	}
	resolvers, rErr := CreateResolvers(database)
	if rErr != nil {
		return nil, rErr
	}
	exec := gql.NewExecutableSchema(resolvers)
	newVertex := &ProximaDataVertex{name: config["name"].(string), id: config["id"].(string) , version: config["version"].(string), applicationDB: database, executableSchema: exec}
	return newVertex, nil
}

func CreateResolvers(db *proxima.ProximaDB) (gql.Config, error) {
	loader, err  := CreateDataloaders(db)
	if err != nil {
		return nil, err
	}
	return resolver.NewResolver(loader, db), nil
}

func CreateDataloaders(db *proxima.ProximaDB) (*dataloader.Dataloader, error) {
  loader , err := dataloader.NewDataloader(db)
  if err != nil {
    return nil, err
  }
  return loader, nil
}

func CreateApplicationDatabase(db_config map[string]interface{}) (*proxima.ProximaDB, error) {
	proximaDB, err := proxima.LoadApplicationDatabase(db_config)
	if err != nil {
		return nil, err
	}
	proximaDB.Sync();
	return proximaDB, nil
}

func (vertex *ProximaDataVertex) startVertexServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	go r.POST("/query", vertex.query())
	go r.GET("/", vertex.playgroundHandler())
	r.Run(":4000")
}

func (vertex *ProximaDataVertex) query() gin.HandlerFunc {
	h := handler.GraphQL(vertex.executableSchema)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (vertex *ProximaDataVertex) playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
