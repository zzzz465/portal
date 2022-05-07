package runner

import "github.com/zzzz465/portal/sd/internal/types"

type datasource interface {
	Identifier() string
	TTL() int
	FetchRecords() ([]types.Record, error)
}
