package providers

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"os"
	"sync"
)

const (
	ExpirationTime = 60 // 1 minute in seconds
)

var (
	queryCacheOnce sync.Once
	queryCacheProp *QueryCacheProvider
)

type QueryCacheProvider struct {
	client *redis.Client
}

func GetQueryCacheProvider() *QueryCacheProvider {
	queryCacheOnce.Do(func() {
		queryCacheProp = &QueryCacheProvider{
			client: redis.NewClient(&redis.Options{
				Addr:     os.Getenv("REDIS_ADDR"),
				Password: os.Getenv("REDIS_PASSWORD"),
				DB:       0,
			}),
		}
	})

	return queryCacheProp
}

func (p *QueryCacheProvider) Wrap(db *gorm.DB) *gorm.DB {
	// Structure the query as a String
	query := db.Statement.SQL.String()

	// Check if the query is already cached
	resultCached, err := p.client.Get(context.Background(), query).Result()

	// If the query is not cached, register the callback
	if err != nil {
		db.Callback().Query().After("gorm:query").Register("query_cache", func(db *gorm.DB) {
			// Get the result of the query
			value, err := db.Rows()
			if err != nil {
				return
			}

			// Convert the result into a JSON string
			result, err := json.Marshal(value)
			if err != nil {
				return
			}

			// Execute the query and cache the result
			p.client.Set(context.Background(), db.Statement.SQL.String(), result, ExpirationTime)
		})

		return db
	}

	// Parse the cached result into the appropriate format for GORM
	var result interface{}

	// Assuming the cached result is a string, you might need to adjust this based on the actual type
	if err := json.Unmarshal([]byte(resultCached), &result); err != nil {
		// Handle unmarshaling error if necessary
		panic("BUG: Failed to unmarshal cached result")
	}

	// Build the query from the cache
	db.Statement.Build(query)

	// Set the parsed result to the rows of the GORM DB
	db.Statement.Dest = result

	return db
}
