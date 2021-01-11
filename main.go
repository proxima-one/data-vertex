package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/99designs/gqlgen/handler"
	yaml "gopkg.in/yaml.v2"
)

func getConfig(configPath string) (map[string]interface{}, error) {
	file, err := os.Open(configPath)
	var objmap map[string]interface{}
	err := yaml.Unmarshal(data, &objmap)
	if err != nil {
		return nil, nil
	}
	return objmap, nil
}

func getDBConfig(configPath string) (map[string]interface{}, error) {
	file, err := os.Open(configPath)
	var objmap map[string]interface{}
	err = yaml.Unmarshal(data, &objmap)
	if err != nil {
		return nil, nil
	}
	return objmap, nil
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	configPath := "./app-config.yml"
	dbConfigPath := "./database/db-config.yaml"

	config, configErr :=  getConfig(configFilePath)
	if configErr != nil {
		return nil, configErr
	}
	dbConfig, dbErr := getDBConfig(dbConfigFilePath)
	if dbErr != nil {
		return nil, dbErr
	}
	applicationVertex, err := CreateDataVertex(config, dbConfig)
	if err != nil {
		return nil, err
	}

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
