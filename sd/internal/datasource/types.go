package datasource

import (
    "github.com/zzzz465/portal/sd/internal/types"
)

type Datasource interface {
    Identifier() string
    FetchRecords() ([]types.Record, error)
}
