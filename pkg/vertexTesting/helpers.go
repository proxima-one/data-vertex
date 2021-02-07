package vertex

import (
	"bytes"
	"context"
	"encoding/json"
	"math/rand"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/errors"
)

type EntityTestCase struct {
	name string
	applicationVertex *ProximaDataVertex
	operations map[string]interface{}
	entities []interface{}
	entity map[string]interface{}
	entityInput map[string]interface{}
	putTests []*GQLTest
	getTests []*GQLTest
	getAllTests  []*GQLTest
	searchTests []*GQLTest
}

type GQLTest struct {
	Context        context.Context
	Schema         *graphql.Schema
	Query          string
	OperationName  string
	Variables      map[string]interface{}
	ExpectedResult string
	ExpectedErrors []*errors.QueryError
}

func NewGQLTest(schema *graphql.Schema, queryString, operationName, expectedResult string, vars, expectedResult, expectedErrors map[string]interface{}) *GQLTest {
	return &GQLTest{Context: context.TODO(), Schema: schema, Query: queryString,
		OperationName: operationName, Variables: vars, ExpectedResult: expectedResult, ExpectedErrors: expectedErrors}
}

func LoadEntityTestCases(applicationVertex *ProximaDataVertex, fileName string) (map[string]*EntityTestCase, error){
	entityTestCaseMap := make(map[string]*EntityTestCase)
	entityConfigMap := make(map[string]interface{})
	data, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(data, &entityConfigMap)
	var entity map[string]interface{}
	for entityName, entityConfig := range entityConfigMap {
		entity = entityConfig.(map[string]interface{})
		entityTestCaseMap[entityName], _ = parseEntity(aentity)
	}
	return entityTestCaseMap, nil
}

func NewEntityTestCase(name string, applicationVertex *ProximaDataVertex, entity, entityInput, operations map[string]interface{}) (*EntityTestCase) {
	entities := make([]interface{}, 0)
	putTests := make([]*GQLTest, 0)
	getTests := make([]*GQLTest, 0)
	getAllTests := make([]*GQLTest, 0)
	searchTests := make([]*GQLTest, 0)
	entityTestCase := nil
	return entityTestCase
}

func parseEntity(applicationVertex *ProximaDataVertex, entityConfig map[string]interface{}) (*EntityTestCase, error) {
	name := entityConfig["name"].(string)
	applicationVertex := applicationVertex
	entityInput := entityConfig["entityInput"].(map[string]interface{})
	entity := entityConfig["entity"].(map[string]interface{})
	operations := entityConfig["operations"].(map[string]interface{})
	return NewEntityTestCase(name, applicationVertex, entity, entityInput, operations), nil
}

func RunEntityTestCases(t *testing.T, tests map[string]*EntityTestCase) {
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			RunEntityTestCase(t, testCase)
		})
	}
}

func RunEntityTestCase(t *testing.T, entityTestCase *EntityTestCase) {
	t.Run("Put Tests", func(t *testing.T) {
		RunGQLTests(t, entityTestCase.putTests)
	})

	t.Run("Get Tests", func(t *testing.T) {
		RunGQLTests(t, entityTestCase.getTests)
	})

	t.Run("Get All Test", func(t *testing.T) {
		RunGQLTests(t, entityTestCase.getAllTests)
	})

	t.Run("Search Test", func(t *testing.T) {
		RunGQLTests(t, entityTestCase.searchTests)
	})
}

func RunGQLTests(t *testing.T, tests []*GQLTest)  {
		for i, test := range tests {
			t.Run(strconv.Itoa(i+1), func(t *testing.T) {
				RunGQLTest(t, test)
			})
		}
}

func RunGQLTest(t *testing.T, test *GQLTest) {
	if test.Context == nil {
		test.Context = context.Background()
	}
	result := test.Schema.Exec(test.Context, test.Query, test.OperationName, test.Variables)
	checkErrors(t, test.ExpectedErrors, result.Errors)
	checkResult(t, test.ExpectedResult, result.Data)
	checkProof(t, test, result.Data)
	checkAudit(t, test, result.Data)
}

func checkProof(t *testing.T, test *GQLTest, data map[string]interface{}) {
	//get prove boolean
	//if prove
		//get proof from data
		//run the correct check proof for type of query (from type of operation)
	//not implemented
	return
}

func checkAudit(t *testing.T, test *GQLTest, data map[string]interface{}) {
	//if audit
		//run the correct auditting function (for get, put, etc):
		//get audit fn
		//check audit
	//not implemented
	return
}

func checkResult(t *testing.T, test *GQL, expectedResult, actualResult interface{}) {
	if expectedResult == "" && actualResult != nil {
			t.Fatalf("got: %s", actualResult)
			t.Fatalf("want: null")
	}
	expected, err := formatJSON([]byte(expectedResult))
	if err != nil {
		t.Fatalf("want: invalid JSON: %s", err)
	}
	actual, err := formatJSON(actualResult)
	if err != nil {
		t.Fatalf("got: invalid JSON: %s", err)
	}
	CompareEntities(expected, actual); //
}

func CompareEntities(expected, actual interface{}) {
	//check type
	//switch case


	//contains all of aspects of comparison
	//check for lists
		//process lists
	//check for maps
	//for elements of "lead" (don't compare proof)
		//compare the values (name), then convert to bytes

	//
	//name, expected, result, testing
	return
}


func (entityTest *EntityTestCase)generateTestEntities(num int) ([]interface{}, error) {
	entities := make([]interface{}, 0)
	for i := 0; i < num; i++ {
		entityMap := make(map[string]interface{})
		entityMap["entityInput"] := GenerateRandomStruct(entityTest.entityInput.(map[string]interface{}))
		entityMap["entity"] := GenerateRandomStruct(entityTest.entity.(map[string]interface{}))
		entities = append(entities, entityMap)
	}
	entityTestCase.entities = entities
	return entities, nil
}

func (entityTest *EntityTestCase) generateTests(num int) {
	entityTest.putTests := make([]*GQLTest, 0)
	entityTest.getTests := make([]*GQLTest, 0)
	entityTest.getAllTests := make([]*GQLTest, 0)
	entityTest.searchTests := make([]*GQLTest, 0)

	entities, _ := entityTest.generateTestEntities(num)
	for i := 0; i < num; i++ {
		entity := entities[i].(map[string]interface{})
		entityTest.putTests = append(putTests, entityTest.generatePutTest(entity))
		entityTest.getTests = append(entityTest.getTests, entityTest.generateGetTest(entity))
		entityTest.getAllTests = append(entityTest.getAllTests, entityTest.generateGetAllTest(entities))
		entityTest.searchTests = append(entityTest.searchTests, entityTest.generateSearchTest(entities))
	}
}

func (entityTest *EntityTestCase) generateGetTest(entityMap map[string]interface{}) (*GQLTest) {
	schema := entityTest.applicationVertex.executableSchema
	operation := entityTest.operations["get"].(map[string]interface{})
	queryString := operation["queryString"]
	operationName := operation["name"]

	entity := entityMap["entity"].(map[string]interface{})

	vars := make(map[string]interface{})
	vars["id"] = entity["id"]
	vars["prove"] = false

	expectedResult := entity
	expectedErrors  := nil
	return NewGQLTest(schema, queryString, operationName, vars, expectedResult, expectedErrors)
}

func (entityTest *EntityTestCase) generatePutTest(entityMap map[string]interface{}) (*GQLTest){
	schema := entityTest.applicationVertex.executableSchema
	operation := entityTest.operations["put"].(map[string]interface{})
	queryString := operation["queryString"]
	operationName := operation["name"]

	entityInput := entityMap["entityInput"].(map[string]interface{})

	vars := make(map[string]interface{})
	vars["prove"] = false
	vars["input"] = entityInput
	expectedResult := true
	expectedErrors  := nil
	return NewGQLTest(schema, queryString, operationName, vars, expectedResult, expectedErrors)
}

func (entityTest *EntityTestCase) generateGetAllTest(entities []interface{}) (*GQLTest) {
	schema := entityTest.applicationVertex.executableSchema
	operation := entityTest.operations["getAll"].(map[string]interface{})
	queryString := operation["queryString"]
	operationName := operation["name"]

	rand.Seed(time.Now().UnixNano())

	vars := make(map[string]interface{})
	vars["proof"] = false
	isFirst := true
	numEntities := len(entities)
	vars["limit"] := rand.Intn(numEntities)
	num :=  rand.Intn(numEntities)
	if isFirst {
		vars["first"] = num
		vars["last"] = 0
	} else {
		vars["first"] = 0
		vars["last"] = num
	}

	expectedResult := make(interface{})
	expectedErrors  := nil
	return NewGQLTest(schema, queryString, operationName, vars, expectedResult, expectedErrors)
}

func (entityTest *EntityTestCase) generateSearchTest(entities []interface{}) (*GQLTest) {
	schema := entityTest.applicationVertex.executableSchema
	operation := entityTest.operations["search"].(map[string]interface{})
	queryString := operation["queryString"]
	operationName := operation["name"]
	vars := make(map[string]interface{})

	vars["queryText"] = "Not implemented"

	vars["proof"] = false
	expectedResult := entity.entity
	expectedErrors  := nil
	return NewGQLTest(schema, queryString, operationName, vars, expectedResult, expectedErrors)
}

func GenerateRandomStruct(entityStruct map[string]interface{}) (map[string]interface{}, error) {
	randomEntity := make(map[string]interface{})
	for varName, varType := entityStruct {
		randomEntity[varName], _ = GenerateRandomOfType(varType)
	}
	return randomEntity, nil
}

func RandomString(size int) (string) {
	rand.Seed(time.Now().UnixNano())
  bytes := make([]byte, size)
  rand.Read(bytes)
  return string(bytes)
}

func GenerateRandomOfType(varType string) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	switch (varType) {
		case "String":
			return RandomString(32), nil
		case "Float":
			return rand.NormFloat64(), nil
		case "ID":
			return RandomString(32)), nil
		case "Int":
			return rand.Int(), nil
		case "Bool":
			return (rand.Intn(2) != 0), nil
		default:
			return RandomString(32), nil
	}
}

func formatJSON(data []byte) ([]byte, error) {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	formatted, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return formatted, nil
}

func checkErrors(t *testing.T, want, got []*errors.QueryError) {
	sortErrors(want)
	sortErrors(got)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected error: got %+v, want %+v", got, want)
	}
}

func sortErrors(errors []*errors.QueryError) {
	if len(errors) <= 1 {
		return
	}
	sort.Slice(errors, func(i, j int) bool {
		return fmt.Sprintf("%s", errors[i].Path) < fmt.Sprintf("%s", errors[j].Path)
	})
}
