package resolvers
//go:generate go run github.com/99designs/gqlgen
import (
	_ "github.com/proxima-one/proxima-data-vertex/pkg/models"
	gql "github.com/proxima-one/proxima-data-vertex/pkg/gql"
	dataloader "github.com/proxima-one/proxima-data-vertex/pkg/dataloaders"
	proxima "github.com/proxima-one/proxima-db-client-go/pkg/database"
	_ "github.com/json-iterator/go"
	_ "fmt"
)

var DefaultInputs map[string]interface{} = map[string]interface{}{"prove": false, "limit": 100}

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
type Resolver struct {
	db *proxima.ProximaDatabase
	loader *dataloader.Dataloader
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }


func NewResolver(loader *dataloader.Dataloader, db *proxima.ProximaDatabase) (gql.ResolverRoot) {
	return &Resolver{db: db, loader: loader}
}
