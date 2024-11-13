package data

import (
	"blug/internal/conf"
	"blug/internal/data/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewFriendRepo)

// Data .
type Data struct {
	DB    *ent.Client
	Cache *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	entClient := newDB(c, logger)
	cache := newCache(c, logger)
	return &Data{
		DB:    entClient,
		Cache: cache,
	}, cleanup, nil
}

func newDB(c *conf.Data, logger log.Logger) *ent.Client {
	cli, err := ent.Open("postgres", c.Database.Source)

	if err != nil {
		log.NewHelper(logger).Fatalf("failed opening connection to postgres: %v", err)
		panic(err)
	}
	return cli
}

func newCache(c *conf.Data, logger log.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           0,
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	if err := rdb.Close(); err != nil {
		log.NewHelper(logger).Fatalf("failed opening connection to redis: %v", err)
	}
	return rdb
}
