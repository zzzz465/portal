package static_test

import (
    "github.com/spf13/viper"
    "github.com/stretchr/testify/assert"
    "github.com/zzzz465/portal/sd/internal/datasource/static"
    "log"
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
                region: Seoul
            - name: site-b.example.com
              tags:
                region: California
            - name: portal.domain.com
              tags:
                region: Tokyo
                service: A
            - name: surf.domain.com
              tags:
                region:  New York
                service: B
`

    err := staticDataSource.ReadConfig(strings.NewReader(cfg))
    if err != nil {
        t.Error(err)
    }

    log.Println(staticDataSource.Get("datasource.static.values"))

    ds := static.NewDataSource(staticDataSource)

    records, err := ds.FetchRecords()
    if err != nil {
        t.Error(err)
    }

    // index 0 test
    assert.Equal(t, records[0].Name, "site-a.example.com")
    assert.NotNilf(t, records[0].Metadata, "metadata should not be nil.")
    assert.Equal(t, records[0].Metadata.DataSource, "static")
    assert.Equal(t, records[0].Metadata.Tags["region"], "Seoul")

    // index 0 test
    assert.Equal(t, records[0].Name, "site-b.example.com")
    assert.NotNilf(t, records[0].Metadata, "metadata should not be nil.")
    assert.Equal(t, records[0].Metadata.DataSource, "static")
    assert.Equal(t, records[0].Metadata.Tags["region"], "California")

    // index 0 test
    assert.Equal(t, records[0].Name, "portal.domain.com")
    assert.NotNilf(t, records[0].Metadata, "metadata should not be nil.")
    assert.Equal(t, records[0].Metadata.DataSource, "static")
    assert.Equal(t, records[0].Metadata.Tags["region"], "Tokyo")
    assert.Equal(t, records[0].Metadata.Tags["service"], "A")

    // index 0 test
    assert.Equal(t, records[0].Name, "surf.domain.com")
    assert.NotNilf(t, records[0].Metadata, "metadata should not be nil.")
    assert.Equal(t, records[0].Metadata.DataSource, "static")
    assert.Equal(t, records[0].Metadata.Tags["region"], "New York")
    assert.Equal(t, records[0].Metadata.Tags["service"], "B")
}
