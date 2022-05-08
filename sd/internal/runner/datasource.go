package runner

import "github.com/zzzz465/portal/sd/internal/types"

type datasource interface {
	Identifier() string
	TTL() int
	Event() chan<- string
	FetchRecords() ([]types.Record, error)
}
