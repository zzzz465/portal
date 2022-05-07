package awsroute53

import (
	"github.com/aws/aws-sdk-go-v2/service/route53"
	types2 "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/zzzz465/portal/sd/internal/types"
	"go.uber.org/zap"
)

const (
	DataSourceId = "Route53"
)

type DataSource struct {
	route53Client *route53.Client
	log           *zap.SugaredLogger
}

func NewDataSource(client *route53.Client, log *zap.SugaredLogger) (*DataSource, error) {
	ds := &DataSource{
		route53Client: client,
		log:           log,
	}

	if log == nil {
		log, err := zap.NewDevelopment()
		if err != nil {
			return nil, err
		}

		ds.log = log.Sugar()
	}

	return ds, nil
}

// FetchRecords returns TODO: add comment
func (ds *DataSource) FetchRecords() ([]types.Record, error) {
	recordSets, err := GetAllRecordSets(ds.route53Client)
	if err != nil {
		return nil, err
	}

	ds.log.Debugf("fetched records: %v", recordSets)

	return toRecords(recordSets), nil
}

func toRecords(recordSets []types2.ResourceRecordSet) []types.Record {
	records := make([]types.Record, 0)
	for _, recordSet := range recordSets {
		for _, record := range recordSet.ResourceRecords {
			records = append(records, types.Record{
				IP:   "",
				Host: *record.Value,
				Metadata: types.RecordMetadata{
					DataSource: DataSourceId,
					Tags:       map[string]string{},
				},
			})
		}
	}

	return records
}
