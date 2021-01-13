package main

import (
	"os"
	"github.com/gin-gonic/gin"
	vertex "github.com/proxima-one/proxima-data-vertex/pkg/vertex"
	yaml "gopkg.in/yaml.v2"
)

func getConfig(configPath string) (map[string]interface{}, error) {
	data, err := os.Open(configPath)
	var objmap map[string]interface{}
	err = yaml.Unmarshal(data, &objmap)
	if err != nil {
		return nil, nil
	}
	return objmap, nil
}

func getDBConfig(configPath string) (map[string]interface{}, error) {
	data, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	var objmap map[string]interface{};
	err = yaml.Unmarshal(data, &objmap)
	if err != nil {
		return nil, err
	}
	return objmap, nil
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	configFilePath := "./app-config.yml"
	dbConfigFilePath := "./database/db-config.yaml"

	config, configErr :=  getConfig(configFilePath)
	// if configErr != nil {
	// 	return nil, configErr
	// }
	dbConfig, dbErr := getDBConfig(dbConfigFilePath)
	// if dbErr != nil {
	// 	return nil, dbErr
	// }
	applicationVertex, err := vertex.CreateDataVertex(config, dbConfig)
	// if err != nil {
	// 	return nil, err
	// }

	applicationVertex.startVertexServer()
}

// r := gin.Default()
// go r.POST("/query", applicationVertex.query())
// go r.GET("/", playgroundHandler())
// r.Run(":4000")

// func playgroundHandler() gin.HandlerFunc {
// 	h := handler.Playground("GraphQL", "/query")
// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }
