package awsroute53

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/service/route53"
    "github.com/aws/aws-sdk-go-v2/service/route53/types"
    "github.com/cockroachdb/errors"
)

func GetAllRecordSets(client *route53.Client) ([]types.ResourceRecordSet, error) {
    hostedZones, err := ListAllHostedZones(client)
    if err != nil {
        return nil, err
    }

    return ListAllRecords(client, hostedZones)
}

func ListAllHostedZones(client *route53.Client) ([]types.HostedZone, error) {
    var marker *string
    hostedZones := make([]types.HostedZone, 0)

    for {
        output, err := client.ListHostedZones(context.TODO(), &route53.ListHostedZonesInput{Marker: marker})
        if err != nil {
            return hostedZones, errors.Wrap(err, "failed ListHostedZones")
        }

        hostedZones = append(hostedZones, output.HostedZones...)
        marker = output.NextMarker

        if marker == nil {
            break
        }
    }

    return hostedZones, nil
}

func ListAllRecords(client *route53.Client, hostedZones []types.HostedZone) ([]types.ResourceRecordSet, error) {
    var startRecordName *string
    records := make([]types.ResourceRecordSet, 0)

    for _, zone := range hostedZones {
        output, err := client.ListResourceRecordSets(context.TODO(), &route53.ListResourceRecordSetsInput{
            HostedZoneId:    zone.Id,
            StartRecordName: startRecordName,
        })
        if err != nil {
            return nil, errors.Wrap(err, "failed ListResourceRecordSets")
        }

        records = append(records, output.ResourceRecordSets...)

        startRecordName = output.NextRecordName
        if startRecordName == nil {
            break
        }
    }

    return records, nil
}
