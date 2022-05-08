package static

import (
	"github.com/spf13/viper"
	"github.com/zzzz465/portal/sd/internal/types"
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

func (ds *DataSource) FetchRecords() ([]types.Record, error) {
	// TODO: impl this

	return nil, nil
}
