package runner

import (
    "github.com/zzzz465/portal/sd/internal/types"
    "time"
)

type datasource interface {
    Identifier() string
    TTL() time.Duration
    FetchRecords() ([]types.Record, error)
}

type updatable interface {
    datasource
    OnDatasourceUpdated(cb func())
}
