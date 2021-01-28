package vertex

import (
	proxima "github.com/proxima-one/proxima-db-client-go/pkg/database"
	resolver "github.com/proxima-one/proxima-data-vertex/pkg/resolvers"
	dataloader "github.com/proxima-one/proxima-data-vertex/pkg/dataloaders"
	"github.com/99designs/gqlgen/handler"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	gql "github.com/proxima-one/proxima-data-vertex/pkg/gql"
)

type ProximaDataVertex struct {
  name string
  id string
  version string
  applicationDB *proxima.ProximaDatabase
	executableSchema *graphql.ExecutableSchema
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

func CreateResolvers(db *proxima.ProximaDatabase) (gql.Config, error) {
	loader, err  := CreateDataloaders(db)
	if err != nil {
		return nil, err
	}
	return resolver.NewResolver(loader, db), nil
}

func CreateDataloaders(db *proxima.ProximaDatabase) (*dataloader.Dataloader, error) {
  loader , err := dataloader.NewDataloader(db)
  if err != nil {
    return nil, err
  }
  return loader, nil
}

func CreateApplicationDatabase(db_config map[string]interface{}) (*proxima.ProximaDatabase, error) {
	proximaDB, err := proxima.LoadProximaDatabase(db_config)
	if err != nil {
		return nil, err
	}
	proximaDB.Sync()
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

// // Group using gin.BasicAuth() middleware
// 	// gin.Accounts is a shortcut for map[string]string
// 	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
// 		"foo":    "bar",
// 		"austin": "1234",
// 		"lena":   "hello2",
// 		"manu":   "4321",
// 	}))
//
// 	// /admin/secrets endpoint
// 	// hit "localhost:8080/admin/secrets
// 	authorized.GET("/secrets", func(c *gin.Context) {
// 		// get user, it was set by the BasicAuth middleware
// 		user := c.MustGet(gin.AuthUserKey).(string)
// 		if secret, ok := secrets[user]; ok {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
// 		}
// 	})
