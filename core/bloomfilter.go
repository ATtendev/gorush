package core

type BloomFilter interface {
	Init() error
	Add(key []byte) error
	Exist(key []byte) bool
}
