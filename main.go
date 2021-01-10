package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/99designs/gqlgen/handler"
	yaml "gopkg.in/yaml.v2"
)

//get the config
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

//start
func main() {
	gin.SetMode(gin.ReleaseMode)
	configPath := "../app-config.yml"
	dbConfigPath := "../database/db-config.yaml"
	config :=  getConfig(configFilePath)
	dbConfig := getDBConfig(dbConfigFilePath)

	applicationVertex, _ := CreateDataVertex(config, dbConfig)

	r := gin.Default()
	go r.POST("/query", applicationVertex.query())
	go r.GET("/", playgroundHandler())
	r.Run(":4000")
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
