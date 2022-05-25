package static

import (
    "github.com/cockroachdb/errors"
    "github.com/fsnotify/fsnotify"
    "github.com/mitchellh/mapstructure"
    "github.com/spf13/viper"
    "github.com/zzzz465/portal/sd/internal/types"
    "time"
)

type DataSource struct {
    staticDataSource      *viper.Viper
    onDatasourceUpdatedCb func()
}

func NewDataSource(staticDataSource *viper.Viper) *DataSource {
    ds := &DataSource{
        staticDataSource: staticDataSource,
    }

    staticDataSource.OnConfigChange(ds.configChangeCb)

    return ds
}

func (ds *DataSource) Identifier() string {
    return "static"
}

func (ds *DataSource) TTL() time.Duration {
    return -1
}

func (ds *DataSource) FetchRecords() ([]types.Record, error) {
    recordMaps, ok := ds.staticDataSource.Get("records").([]interface{})
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

func (ds *DataSource) configChangeCb(in fsnotify.Event) {
    // TODO: how to handle this error?
    // TODO: this callback called twice. apply debounce required.
    _ = ds.staticDataSource.ReadInConfig()

    if ds.onDatasourceUpdatedCb != nil {
        ds.onDatasourceUpdatedCb()
    }
}

func (ds *DataSource) OnDatasourceUpdated(cb func()) {
    ds.onDatasourceUpdatedCb = cb
}
