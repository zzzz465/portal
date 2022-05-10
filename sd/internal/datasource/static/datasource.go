package static

import (
    "github.com/cockroachdb/errors"
    "github.com/mitchellh/mapstructure"
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
    recordMaps, ok := ds.staticDataSource.Get("datasource.static.values").([]interface{})
    if !ok {
        return nil, errors.New("cannot read static values from config.")
    }

    //records := lo.RepeatBy(len(recordMaps), func(_ int) types.Record { return types.NewRecord() })
    records := make([]types.Record, len(recordMaps))
    for i, record := range recordMaps {
        err := mapstructure.Decode(record, &records[i])
        if err != nil {
            return nil, err
        }

        records[i].Metadata.DataSource = ds.Identifier()
    }

    return records, nil
}
