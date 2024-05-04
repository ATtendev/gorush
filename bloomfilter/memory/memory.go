package memory

import (
	"github.com/appleboy/gorush/config"
	"github.com/bits-and-blooms/bloom/v3"
)

type BloomFilter struct {
	config *config.ConfYaml
	bl     *bloom.BloomFilter
}

func New(config *config.ConfYaml) *BloomFilter {
	return &BloomFilter{
		config: config,
	}
}

func (b *BloomFilter) Init() error {
	b.bl = bloom.New(b.config.BloomFilter.Size, b.config.BloomFilter.HashNum)
	return nil
}

func (b *BloomFilter) Add(key []byte) error {
	b.bl.Add(key)
	return nil
}

func (b *BloomFilter) Exist(key []byte) bool {
	return b.bl.Test(key)
}
