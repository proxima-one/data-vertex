package gql
import (
  "github.com/99designs/gqlgen/graphql"
)

type Config struct {
  Resolvers interface{}
}

type QueryResolver struct {
  val bool
}

func NewExecutableSchema(config Config) (graphql.ExecutableSchema) {
  config.Resolvers = nil
  return nil
}
