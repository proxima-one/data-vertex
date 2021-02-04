package vertexTesting

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/errors"
)

func NewEntityTestCase(entityConfig map[string]interface{}) (*EntityTestCase, error) {
	return nil, nil
}

func GenerateRandomOfType(varType string) (interface{}, error) {
	return nil, nil
}

func GetEntities(fileName string) (map[string]EntityTestCase, error){
// 	//read the json file
// 	//unmarshal file into map[string]
// 	//for each key, value in entityMap
// 		//cast map
// 		//parseEntity
// 	//return entities, nil
	return nil, nil
}

func parseEntity(entityConfig map[string]interface{}) (EntityTestCase, error) {
	return nil, nil
}

type EntityTestCase {
	name string
	//construction of the entity test case
		//vars


}

func (entityTest *EntityTestCase) runTests() {
	//for all tests run
}

func (entityTest *EntityTestCase) generateTests() {
	//generateTests
	//generate the entities
	//get Test
	//put Test
	//getAll Test
	//search Test
}

func (entityTest *EntityTestCase) generateGetTest() {

}

func (entityTest *EntityTestCase) generatePutTest() {

}

func (entityTest *EntityTestCase) generateGetAllTest() {}

func (entityTest *EntityTestCase) generateSearchTest() {}


func (entityTest *EntityTestCase) generateRandomEntity() {}

func (entityTest *EntityTestCase) generateRandomFunctionVars() {}


func NewGQLTestCase(schema *graphql.Schema, queryString, operationName, expectedResult string, vars map[string]interface{}) *GQLTest {
	return &GQLTest{Context: context.TODO(), Schema: schema, Query: queryString,
		OperationName: operationName, Variables: vars, ExpectedResult: expectedResult, ExpectedErrors: nil}
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

// RunTests runs the given GraphQL test cases as subtests.
func RunTests(t *testing.T, tests []*Test) {
	if len(tests) == 1 {
		RunTest(t, tests[0])
		return
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			RunTest(t, test)
		})
	}
}

// RunTest runs a single GraphQL test case.
func RunTest(t *testing.T, test *Test) {
	if test.Context == nil {
		test.Context = context.Background()
	}
	result := test.Schema.Exec(test.Context, test.Query, test.OperationName, test.Variables)

	checkErrors(t, test.ExpectedErrors, result.Errors)

	if test.ExpectedResult == "" {
		if result.Data != nil {
			t.Fatalf("got: %s", result.Data)
			t.Fatalf("want: null")
		}
		return
	}

	// Verify JSON to avoid red herring errors.
	got, err := formatJSON(result.Data)
	if err != nil {
		t.Fatalf("got: invalid JSON: %s", err)
	}
	want, err := formatJSON([]byte(test.ExpectedResult))
	if err != nil {
		t.Fatalf("want: invalid JSON: %s", err)
	}

	if !bytes.Equal(got, want) {
		t.Logf("got:  %s", got)
		t.Logf("want: %s", want)
		t.Fail()
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
