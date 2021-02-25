package resolvers
//go:generate go run github.com/99designs/gqlgen
import (
	models "github.com/proxima-one/proxima-data-vertex/pkg/models"
	gql "github.com/proxima-one/proxima-data-vertex/pkg/gql"
	dataloader "github.com/proxima-one/proxima-data-vertex/pkg/dataloaders"
	_ "github.com/json-iterator/go"
	_ "fmt"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
type Resolver struct{
	loader *dataloader.Dataloader
}

func (r *Resolver) Query() *queryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }


func NewResolver(loader *dataloader.Dataloader) (gql.Config) {
	r := Resolver{}
	r.loader = loader
	return gql.Config{
		Resolvers: &r,
	}
}
