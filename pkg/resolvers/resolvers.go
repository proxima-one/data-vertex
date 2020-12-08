package resolvers
//go:generate go run github.com/99designs/gqlgen
import (
	"context"
	//datasources "github.com/proxima-one/binance-chain-subgraph/pkg/datasources"
	models "github.com/proxima-one/binance-chain-subgraph/pkg/models"
	gql "github.com/proxima-one/binance-chain-subgraph/pkg/gql"
	dataloader "github.com/proxima-one/binance-chain-subgraph/pkg/dataloader"
	//json "github.com/json-iterator/go"
	_ "fmt"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
type Resolver struct{
	loader *dataloader.Dataloader
}

func (r *Resolver) Query() gql.QueryResolver {
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


func (r *queryResolver) BlockStats(ctx context.Context, prove *bool) (*models.ProximaBlockStats, error) {
	args := BlockStatsDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove }
	return r.loader.LoadProximaBlockStats(args)
}

func (r *queryResolver) Fees(rctx context.Context, prove *bool) (*models.ProximaFees, error) {
	args :=  FeesDefaultInputs;
	if (prove != nil ) { args["prove"] = *prove}
	return r.loader.LoadProximaFees(args)
}
