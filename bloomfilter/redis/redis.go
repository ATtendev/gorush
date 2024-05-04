package redis

import (
	"context"
	"strings"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/logx"
	"github.com/redis/go-redis/v9"
)

type BloomFilter struct {
	ctx    context.Context
	config *config.ConfYaml
	rds    redis.UniversalClient
}

func New(config *config.ConfYaml) *BloomFilter {
	return &BloomFilter{
		ctx:    context.Background(),
		config: config,
	}
}

func (b *BloomFilter) Init() error {
	if b.config.BloomFilter.Redis.Cluster {
		b.rds = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    strings.Split(b.config.BloomFilter.Redis.Addr, ","),
			Password: b.config.BloomFilter.Redis.Password,
		})
	} else {
		b.rds = redis.NewClient(&redis.Options{
			Addr:     b.config.BloomFilter.Redis.Addr,
			Password: b.config.BloomFilter.Redis.Password,
			DB:       b.config.BloomFilter.Redis.DB,
		})
	}

	if err := b.rds.Ping(b.ctx).Err(); err != nil {
		return err
	}
	return nil
}

func (b *BloomFilter) Add(key []byte) error {
	_, err := b.rds.Do(b.ctx, "BF.ADD", "bf_key", key).Bool()
	if err != nil {
		logx.LogAccess.Debug(err)
		return err
	}
	return nil
}

func (b *BloomFilter) Exist(key []byte) bool {
	exist, err := b.rds.Do(b.ctx, "BF.EXISTS", "bf_key", key).Bool()
	if err != nil {
		logx.LogAccess.Debug(err)
		return exist
	}
	return exist
}
