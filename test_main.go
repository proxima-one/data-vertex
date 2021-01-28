package main

import (
  "testing"
  "math/rand"
  "fmt"
  "time"
  vertexTesting "github.com/proxima-data-vertex/pkg/vertexTesting"
)

func TestLoadDataVertex(t *testing.T) {
  var configFilePath string = "./app-config.yml"
  var dbConfigFilePath string = "./database/db-config.yaml"

  applicationVertex, err := vertex.LoadDataVertex(configFilePath, dbConfigFilePath)
  if err != nil {
    t.Errorf("Data vertex creation error: %v", err)
  }

  _, queryErr := applicationVertex.query();
  if (queryErr != nil) {
    t.Errorf("Error creating the resolvers", resolverErr);
  }
}

// func TestApplicationResolvers(t *testing.T) {
//   //auto-generated, but using the resolvers config file
//   //vertexTesting
//   //generate the vertex
//   //testCases
//   schema := 0 //vertex.executableSchema()
//   queryString := ""
//   operationName := ""
//   variables := make(map[string]interface{})
//   expectedResult := ""
//   //generate test cases
//   //test batch for the entity...
//   //dummy data test cases
//
//   //
//
//   testCase := vertexTesting.NewDefaultTestCase(schema, queryString, operationName, expectedResult, variables)
//   vertexTesting.RunTest(t, testCase)
// }
