package resolvers

//go:generate go run github.com/99designs/gqlgen
import (
	//dataloader "github.com/proxima-one/proxima-data-vertex/pkg/dataloaders"
	//dataloader "github.com/proxima-one/proxima-data-vertex/pkg/dataloaders"
	gql "github.com/proxima-one/proxima-data-vertex/pkg/gql"
	_ "github.com/proxima-one/proxima-data-vertex/pkg/models"
	proxima "github.com/proxima-one/proxima-db-client-go/pkg/database"
)

var DefaultInputs map[string]interface{}

type Resolver struct {
	//loader *dataloader.Dataloader
	db *proxima.ProximaDatabase
}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

//AUTOGEN RESOLVERS

/*
resolver - database and loader
*/
func NewResolver(db *proxima.ProximaDatabase) gql.Config {
	r := Resolver{}
	r.db = db
	return gql.Config{
		Resolvers: &r,
	}
}
