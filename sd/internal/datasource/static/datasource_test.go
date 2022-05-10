package static_test

import (
    "github.com/spf13/viper"
    "github.com/zzzz465/portal/sd/internal/datasource/static"
    "strings"
    "testing"
)

func TestDataSource_FetchRecords(t *testing.T) {
    staticDataSource := viper.New()
    staticDataSource.SetConfigType("yaml")

    cfg := `
datasource:
    static:
        enabled: true
        values:
            - name: site-a.example.com
              tags:
                region: seoul
                
`

    err := staticDataSource.ReadConfig(strings.NewReader(cfg))
    if err != nil {
        t.Error(err)
    }

    ds := static.NewDataSource(staticDataSource)

    records, err := ds.FetchRecords()
    if err != nil {
        t.Error(err)
    }
}
