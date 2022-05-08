package runner

import "github.com/zzzz465/portal/sd/internal/types"

type datasource interface {
	Identifier() string
	TTL() int // TODO: replace to time.duration
	FetchRecords() ([]types.Record, error)
}
