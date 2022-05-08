package store

import "github.com/zzzz465/portal/sd/internal/types"

// store
// example: InMemoryStore, RedisStore, FileStore, etc...
// TODO: add WriteMany
type Store interface {
	GetRecords() ([]types.Record, error)
	GetRecord(key string) (*types.Record, error)
	WriteRecord(record types.Record) error
	DeleteRecord(key string) error
}
