package static_test

import (
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
    "github.com/stretchr/testify/assert"
    "github.com/zzzz465/portal/sd/internal/datasource/static"
    "io/ioutil"
    "os"
    "strings"
    "testing"
    "time"
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
              metadata:
                  tags:
                    region: Seoul
            - name: site-b.example.com
              metadata:
                  tags:
                    region: California
            - name: portal.domain.com
              metadata:
                  tags:
                    region: Tokyo
                    service: A
            - name: surf.domain.com
              metadata:
                  tags:
                    region:  New York
                    service: B
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

    // index 0 test
    assert.Equal(t, "site-a.example.com", records[0].Name)
    assert.NotNilf(t, records[0].Metadata, "metadata should not be nil.")
    assert.Equal(t, "static", records[0].Metadata.DataSource)
    assert.Equal(t, "Seoul", records[0].Metadata.Tags["region"])

    // index 1 test
    assert.Equal(t, "site-b.example.com", records[1].Name)
    assert.NotNilf(t, records[1].Metadata, "metadata should not be nil.")
    assert.Equal(t, "static", records[1].Metadata.DataSource)
    assert.Equal(t, "California", records[1].Metadata.Tags["region"])

    // index 2 test
    assert.Equal(t, "portal.domain.com", records[2].Name)
    assert.NotNilf(t, records[2].Metadata, "metadata should not be nil.")
    assert.Equal(t, "static", records[2].Metadata.DataSource)
    assert.Equal(t, "Tokyo", records[2].Metadata.Tags["region"])
    assert.Equal(t, "A", records[2].Metadata.Tags["service"])

    // index 3 test
    assert.Equal(t, "surf.domain.com", records[3].Name)
    assert.NotNilf(t, records[3].Metadata, "metadata should not be nil.")
    assert.Equal(t, "static", records[3].Metadata.DataSource)
    assert.Equal(t, "New York", records[3].Metadata.Tags["region"])
    assert.Equal(t, "B", records[3].Metadata.Tags["service"])
}

// TestDataSource_FetchRecords_Watched tests records are changed after config file is changed.
func TestDataSource_FetchRecords_Watched(t *testing.T) {
    f, err := os.CreateTemp("", "test-")
    if err != nil {
        t.Error(err)
    }

    if err = f.Close(); err != nil {
        t.Error(err)
    }
    defer os.Remove(f.Name())

    cfg := `
datasource:
    static:
        enabled: true
        values:
            - name: site-a.example.com
              metadata:
                  tags:
                    region: Seoul
            - name: site-b.example.com
              metadata:
                  tags:
                    region: California
            - name: portal.domain.com
              metadata:
                  tags:
                    region: Tokyo
                    service: A
            - name: surf.domain.com
              metadata:
                  tags:
                    region:  New York
                    service: B
`

    if err = ioutil.WriteFile(f.Name(), []byte(cfg), 0644); err != nil {
        t.Error(err)
    }

    staticDataSource := viper.New()
    staticDataSource.SetConfigFile(f.Name())
    staticDataSource.SetConfigType("yaml")
    staticDataSource.WatchConfig()

    err = staticDataSource.ReadInConfig()
    if err != nil {
        t.Error(err)
    }

    ds := static.NewDataSource(staticDataSource)

    // 1. check initial config.
    records, err := ds.FetchRecords()
    if err != nil {
        t.Error(err)
    }

    // index 0 test
    assert.Equal(t, "site-a.example.com", records[0].Name)
    assert.NotNilf(t, records[0].Metadata, "metadata should not be nil.")
    assert.Equal(t, "static", records[0].Metadata.DataSource)
    assert.Equal(t, "Seoul", records[0].Metadata.Tags["region"])

    // index 1 test
    assert.Equal(t, "site-b.example.com", records[1].Name)
    assert.NotNilf(t, records[1].Metadata, "metadata should not be nil.")
    assert.Equal(t, "static", records[1].Metadata.DataSource)
    assert.Equal(t, "California", records[1].Metadata.Tags["region"])

    // index 2 test
    assert.Equal(t, "portal.domain.com", records[2].Name)
    assert.NotNilf(t, records[2].Metadata, "metadata should not be nil.")
    assert.Equal(t, "static", records[2].Metadata.DataSource)
    assert.Equal(t, "Tokyo", records[2].Metadata.Tags["region"])
    assert.Equal(t, "A", records[2].Metadata.Tags["service"])

    // index 3 test
    assert.Equal(t, "surf.domain.com", records[3].Name)
    assert.NotNilf(t, records[3].Metadata, "metadata should not be nil.")
    assert.Equal(t, "static", records[3].Metadata.DataSource)
    assert.Equal(t, "New York", records[3].Metadata.Tags["region"])
    assert.Equal(t, "B", records[3].Metadata.Tags["service"])

    cfg = `
datasource:
    static:
        enabled: true
        values:
            - name: site-b.example.com
              metadata:
                  tags:
                    region: California
`

    if err := ioutil.WriteFile(f.Name(), []byte(cfg), 0644); err != nil {
        t.Error(err)
    }

    waitChan := make(chan interface{})
    timeoutChan := make(chan interface{})
    go func() {
        time.Sleep(10 * time.Second)
        timeoutChan <- struct{}{}
    }()

    staticDataSource.OnConfigChange(func(_ fsnotify.Event) {
        waitChan <- struct{}{}
    })

    select {
    case <-waitChan:
        break
    case <-timeoutChan:
        t.Errorf("timeout while waiting configuration change event.")
    }

    // 2. check config after file is updated.
    records, err = ds.FetchRecords()
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, "site-b.example.com", records[0].Name)
    assert.NotNilf(t, records[0].Metadata, "metadata should not be nil.")
    assert.Equal(t, "static", records[0].Metadata.DataSource)
    assert.Equal(t, "California", records[0].Metadata.Tags["region"])
}
