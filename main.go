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
	err = yaml.Unmarshal([]byte(data), &objmap)
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
	err = yaml.Unmarshal([]byte(data), &objmap)
	if err != nil {
		return nil, err
	}
	return objmap, nil
}

func LoadDataVertex(appConfigFile, dbConfigFile string) (*ProximaDataVertex, error) {
	config, configErr :=  getConfig(configFilePath)
	if configErr != nil {
		log.Fatalf("Application config reading error: %v", configErr)
	}
	dbConfig, dbErr := getDBConfig(dbConfigFilePath)
	if dbErr != nil {
		log.Fatalf("Database config readig error: %v", dbErr)
	}
	applicationVertex, err := vertex.CreateDataVertex(config, dbConfig)
	if err != nil {
		log.Fatalf("Data vertex creation error: %v", err)

	}
	return applicationVertex, err
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	configFilePath := "./app-config.yml"
	dbConfigFilePath := "./database/db-config.yaml"

	config, configErr :=  getConfig(configFilePath)
	if configErr != nil {
		log.Fatalf("Application config reading error: %v", configErr)
	}
	dbConfig, dbErr := getDBConfig(dbConfigFilePath)
	if dbErr != nil {
		log.Fatalf("Database config readig error: %v", dbErr)
	}
	applicationVertex, err := vertex.CreateDataVertex(config, dbConfig)
	if err != nil {
		log.Fatalf("Data vertex creation error: %v", err)
	}
	applicationVertex.startVertexServer()
}
