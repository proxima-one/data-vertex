package dataloader

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/json-iterator/go"
	"github.com/proxima-one/proxima-data-vertex/pkg/models"
	proxima "github.com/proxima-one/proxima-db-client-go/pkg/database"
)

const loadersKey = "dataloaders"

type Loaders struct {
	DPoolListById           DPoolListLoader
	DPoolById               DPoolLoader
	UserById                UserLoader
	UserTotalDepositById    UserTotalDepositLoader
	DepositById             DepositLoader
	FunderById              FunderLoader
	FunderTotalInterestById FunderTotalInterestLoader
	FundingById             FundingLoader
	MPHHolderById           MPHHolderLoader
	MPHById                 MPHLoader
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func Middleware(db *proxima.ProximaDatabase) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), loadersKey, &Loaders{
			DPoolListById: DPoolListLoader{
				maxBatch: 100,
				wait:     5 * time.Millisecond,
				fetch: func(ids []string) ([]*models.DPoolList, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					table, _ := db.GetTable("DPoolLists")
					fmt.Println("Table")
					fmt.Print(table)
					valueById := map[string]*models.DPoolList{}

					for _, key := range ids {
						result, err := table.Get(key, false)
						data := result.GetValue()

						var value models.DPoolList
						json.Unmarshal(data, &value)
						if err != nil {
							panic(err)
						}
						valueById[*value.ID] = &value
					}
					values := make([]*models.DPoolList, len(ids))
					for i, id := range ids {
						values[i] = valueById[id]
					}
					return values, nil
				},
			},
			DPoolById: DPoolLoader{
				maxBatch: 100,
				wait:     50 * time.Millisecond,
				fetch: func(ids []string) ([]*models.DPool, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					table, _ := db.GetTable("DPools")
					valueById := map[string]*models.DPool{}

					for _, key := range ids {
						result, err := table.Get(key, false)
						data := result.GetValue()

						var value models.DPool
						json.Unmarshal(data, &value)
						if err != nil {
							panic(err)
						}
						valueById[*value.ID] = &value
					}
					values := make([]*models.DPool, len(ids))
					for i, id := range ids {
						values[i] = valueById[id]
					}
					return values, nil
				},
			},
			UserById: UserLoader{
				maxBatch: 100,
				wait:     5 * time.Millisecond,
				fetch: func(ids []string) ([]*models.User, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					table, _ := db.GetTable("Users")
					valueById := map[string]*models.User{}

					for _, key := range ids {
						result, err := table.Get(key, false)
						data := result.GetValue()

						var value models.User
						json.Unmarshal(data, &value)
						if err != nil {
							panic(err)
						}
						valueById[*value.ID] = &value
					}
					values := make([]*models.User, len(ids))
					for i, id := range ids {
						values[i] = valueById[id]
					}
					return values, nil
				},
			},
			UserTotalDepositById: UserTotalDepositLoader{
				maxBatch: 100,
				wait:     5 * time.Millisecond,
				fetch: func(ids []string) ([]*models.UserTotalDeposit, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					table, _ := db.GetTable("UserTotalDeposits")
					valueById := map[string]*models.UserTotalDeposit{}

					for _, key := range ids {
						result, err := table.Get(key, false)
						data := result.GetValue()

						var value models.UserTotalDeposit
						json.Unmarshal(data, &value)
						if err != nil {
							panic(err)
						}
						valueById[*value.ID] = &value
					}
					values := make([]*models.UserTotalDeposit, len(ids))
					for i, id := range ids {
						values[i] = valueById[id]
					}
					return values, nil
				},
			},
			DepositById: DepositLoader{
				maxBatch: 100,
				wait:     5 * time.Millisecond,
				fetch: func(ids []string) ([]*models.Deposit, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					table, _ := db.GetTable("Deposits")
					valueById := map[string]*models.Deposit{}

					for _, key := range ids {
						result, err := table.Get(key, false)
						data := result.GetValue()

						var value models.Deposit
						json.Unmarshal(data, &value)
						if err != nil {
							panic(err)
						}
						valueById[*value.ID] = &value
					}
					values := make([]*models.Deposit, len(ids))
					for i, id := range ids {
						values[i] = valueById[id]
					}
					return values, nil
				},
			},
			FunderById: FunderLoader{
				maxBatch: 100,
				wait:     5 * time.Millisecond,
				fetch: func(ids []string) ([]*models.Funder, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					table, _ := db.GetTable("Funders")
					valueById := map[string]*models.Funder{}

					for _, key := range ids {
						result, err := table.Get(key, false)
						data := result.GetValue()

						var value models.Funder
						json.Unmarshal(data, &value)
						if err != nil {
							panic(err)
						}
						valueById[*value.ID] = &value
					}
					values := make([]*models.Funder, len(ids))
					for i, id := range ids {
						values[i] = valueById[id]
					}
					return values, nil
				},
			},
			FunderTotalInterestById: FunderTotalInterestLoader{
				maxBatch: 100,
				wait:     5 * time.Millisecond,
				fetch: func(ids []string) ([]*models.FunderTotalInterest, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					table, _ := db.GetTable("FunderTotalInterests")
					valueById := map[string]*models.FunderTotalInterest{}

					for _, key := range ids {
						result, err := table.Get(key, false)
						data := result.GetValue()

						var value models.FunderTotalInterest
						json.Unmarshal(data, &value)
						if err != nil {
							panic(err)
						}
						valueById[*value.ID] = &value
					}
					values := make([]*models.FunderTotalInterest, len(ids))
					for i, id := range ids {
						values[i] = valueById[id]
					}
					return values, nil
				},
			},
			FundingById: FundingLoader{
				maxBatch: 100,
				wait:     5 * time.Millisecond,
				fetch: func(ids []string) ([]*models.Funding, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					table, _ := db.GetTable("Fundings")
					valueById := map[string]*models.Funding{}

					for _, key := range ids {
						result, err := table.Get(key, false)
						data := result.GetValue()

						var value models.Funding
						json.Unmarshal(data, &value)
						if err != nil {
							panic(err)
						}
						valueById[*value.ID] = &value
					}
					values := make([]*models.Funding, len(ids))
					for i, id := range ids {
						values[i] = valueById[id]
					}
					return values, nil
				},
			},
			MPHHolderById: MPHHolderLoader{
				maxBatch: 100,
				wait:     5 * time.Millisecond,
				fetch: func(ids []string) ([]*models.MPHHolder, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					table, _ := db.GetTable("MPHHolders")
					valueById := map[string]*models.MPHHolder{}

					for _, key := range ids {
						result, err := table.Get(key, false)
						data := result.GetValue()

						var value models.MPHHolder
						json.Unmarshal(data, &value)
						if err != nil {
							panic(err)
						}
						valueById[*value.ID] = &value
					}
					values := make([]*models.MPHHolder, len(ids))
					for i, id := range ids {
						values[i] = valueById[id]
					}
					return values, nil
				},
			},
			MPHById: MPHLoader{
				maxBatch: 100,
				wait:     5 * time.Millisecond,
				fetch: func(ids []string) ([]*models.Mph, []error) {
					placeholders := make([]string, len(ids))
					args := make([]interface{}, len(ids))
					for i := 0; i < len(ids); i++ {
						placeholders[i] = "?"
						args[i] = i
					}

					table, _ := db.GetTable("MPHs")
					valueById := map[string]*models.Mph{}

					for _, key := range ids {
						result, err := table.Get(key, false)
						data := result.GetValue()

						var value models.Mph
						json.Unmarshal(data, &value)
						if err != nil {
							panic(err)
						}
						valueById[*value.ID] = &value
					}
					values := make([]*models.Mph, len(ids))
					for i, id := range ids {
						values[i] = valueById[id]
					}
					return values, nil
				},
			},
		})
		c.Request.WithContext(ctx)
		c.Next()
	})
}

//get dataloader
//fetch function

//entity  basics

//by id

//batch
