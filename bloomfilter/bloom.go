package bloomfilter

import (
	"github.com/appleboy/gorush/bloomfilter/memory"
	"github.com/appleboy/gorush/bloomfilter/redis"
	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
)

type BloomFilter struct {
	core.BloomFilter
}

var Bloom *BloomFilter

func InitBloomFilter(conf *config.ConfYaml) {
	var bl core.BloomFilter
	switch conf.BloomFilter.Engine {
	case "memory":
		bl = memory.New(conf)
	case "redis":
		bl = redis.New(conf)
	default:
		logx.LogAccess.Fatal("Unsupported bloom filter engine: " + conf.BloomFilter.Engine)
	}
	if err := bl.Init(); err != nil {
		panic(err)
	}
	Bloom = &BloomFilter{
		bl,
	}
}
