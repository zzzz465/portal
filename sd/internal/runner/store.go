package runner

import "github.com/zzzz465/portal/sd/internal/types"

// store
// example: InMemoryStore, RedisStore, FileStore, etc...
// TODO: add WriteMany
type store interface {
	Write(key string, record types.Record) error
	Delete(key string) error
}
