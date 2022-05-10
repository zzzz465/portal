package static

import (
    "github.com/cockroachdb/errors"
    "github.com/spf13/viper"
    "github.com/zzzz465/portal/sd/internal/types"
    "time"
)

type DataSource struct {
    staticDataSource *viper.Viper
}

func NewDataSource(staticDataSource *viper.Viper) *DataSource {
    ds := &DataSource{
        staticDataSource: staticDataSource,
    }

    return ds
}

func (ds *DataSource) Identifier() string {
    return "static"
}

func (ds *DataSource) TTL() time.Duration {
    return -1
}

func (ds *DataSource) FetchRecords() ([]types.Record, error) {
    records, ok := ds.staticDataSource.Get("datasource.static.values").([]types.Record)
    if !ok {
        return nil, errors.New("cannot read static values from config.")
    }

    return records, nil
}
