package dynamodbkinsumer

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type DynamoDBStreamsKinesisAdapter struct {
	streamsAPI            dynamodbstreamsiface.DynamoDBStreamsAPI
	PartitionKeyAttribute string
}

func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStream(input *kinesis.DescribeStreamInput) (output *kinesis.DescribeStreamOutput, err error) {
	streamsInput := dynamodbstreams.DescribeStreamInput{
		ExclusiveStartShardId: input.ExclusiveStartShardId,
		Limit:                 input.Limit,
		// StreamName should be set as stream arn
		StreamArn: input.StreamName,
	}
	streamsOut, err := ddbska.streamsAPI.DescribeStream(&streamsInput)
	if err != nil {
		return
	}
	shards := []*kinesis.Shard{}
	for _, shard := range streamsOut.StreamDescription.Shards {
		shards = append(shards, &kinesis.Shard{
			ParentShardId: shard.ParentShardId,
			ShardId:       shard.ShardId,
			SequenceNumberRange: &kinesis.SequenceNumberRange{
				EndingSequenceNumber:   shard.SequenceNumberRange.EndingSequenceNumber,
				StartingSequenceNumber: shard.SequenceNumberRange.StartingSequenceNumber,
			},
		})
	}
	var streamStatus string
	switch *streamsOut.StreamDescription.StreamStatus {
	case dynamodbstreams.StreamStatusEnabled:
		streamStatus = kinesis.StreamStatusActive
	case dynamodbstreams.StreamStatusDisabled:
		streamStatus = kinesis.StreamStatusDeleting
	case dynamodbstreams.StreamStatusEnabling:
		streamStatus = kinesis.StreamStatusCreating
	case dynamodbstreams.StreamStatusDisabling:
		streamStatus = kinesis.StreamStatusUpdating
	default:
		err = fmt.Errorf("unsupported StreamStatus: %s", *streamsOut.StreamDescription.StreamStatus)
		return
	}
	output = &kinesis.DescribeStreamOutput{
		StreamDescription: &kinesis.StreamDescription{
			Shards:        shards,
			StreamARN:     streamsOut.StreamDescription.StreamArn,
			StreamStatus:  &streamStatus,
			HasMoreShards: aws.Bool(false),
		},
	}
	return
}

func (ddbska DynamoDBStreamsKinesisAdapter) ListShards(input *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error) {
	// must use DescribeStream() since dynamodbstreams doesn't have ListShards
	kinesisDescribeStreamInput := &kinesis.DescribeStreamInput{
		ExclusiveStartShardId: input.ExclusiveStartShardId,
		Limit:                 input.MaxResults,
		StreamName:            input.StreamName,
	}
	kinesisDescribeStreamOutput, err := ddbska.DescribeStream(kinesisDescribeStreamInput)
	if err != nil {
		return nil, err
	}
	return &kinesis.ListShardsOutput{
		Shards: kinesisDescribeStreamOutput.StreamDescription.Shards,
	}, nil
}

func (ddbska DynamoDBStreamsKinesisAdapter) GetShardIterator(input *kinesis.GetShardIteratorInput) (output *kinesis.GetShardIteratorOutput, err error) {
	streamsOut, err := ddbska.streamsAPI.GetShardIterator(&dynamodbstreams.GetShardIteratorInput{
		ShardId:           input.ShardId,
		ShardIteratorType: input.ShardIteratorType,
		StreamArn:         input.StreamName,
		SequenceNumber:    input.StartingSequenceNumber,
	})
	if err != nil {
		return
	}
	output = &kinesis.GetShardIteratorOutput{
		ShardIterator: streamsOut.ShardIterator,
	}
	return
}

type StreamRecord struct {
	ApproximateCreationDateTime *time.Time
	Keys                        map[string]interface{}
	NewImage                    map[string]interface{}
	OldImage                    map[string]interface{}
	SequenceNumber              *string
	SizeBytes                   *int64
	StreamViewType              *string
}

type streamRecord dynamodbstreams.StreamRecord

func (sr streamRecord) MarshalJSON() (b []byte, err error) {
	keys := make(map[string]interface{})
	err = dynamodbattribute.UnmarshalMap(sr.Keys, &keys)
	if err != nil {
		return
	}
	newImage := make(map[string]interface{})
	err = dynamodbattribute.UnmarshalMap(sr.NewImage, &newImage)
	if err != nil {
		return
	}
	oldImage := make(map[string]interface{})
	err = dynamodbattribute.UnmarshalMap(sr.OldImage, &oldImage)
	if err != nil {
		return
	}
	h := StreamRecord{
		ApproximateCreationDateTime: sr.ApproximateCreationDateTime,
		Keys:                        keys,
		NewImage:                    newImage,
		OldImage:                    oldImage,
		SequenceNumber:              sr.SequenceNumber,
		SizeBytes:                   sr.SizeBytes,
		StreamViewType:              sr.StreamViewType,
	}
	b, err = json.Marshal(h)
	return
}

func (ddbska DynamoDBStreamsKinesisAdapter) GetRecords(input *kinesis.GetRecordsInput) (output *kinesis.GetRecordsOutput, err error) {
	var limit *int64
	if input.Limit != nil {
		// DynamoDB Streams has max limit of 1000 records
		if *input.Limit > 1000 {
			*limit = 1000
		} else {
			limit = input.Limit
		}
	}
	streamsOut, err := ddbska.streamsAPI.GetRecords(&dynamodbstreams.GetRecordsInput{
		Limit:         limit,
		ShardIterator: input.ShardIterator,
	})
	if err != nil {
		return
	}
	records := make([]*kinesis.Record, len(streamsOut.Records))
	for i, record := range streamsOut.Records {
		streamRecord := streamRecord(*record.Dynamodb)
		var data []byte
		data, err = streamRecord.MarshalJSON()
		if err != nil {
			return
		}
		var keys map[string]interface{}
		dynamodbattribute.UnmarshalMap(record.Dynamodb.Keys, &keys)
		records[i] = &kinesis.Record{
			ApproximateArrivalTimestamp: record.Dynamodb.ApproximateCreationDateTime,
			SequenceNumber:              record.Dynamodb.SequenceNumber,
			Data:                        data,
			PartitionKey:                aws.String(fmt.Sprintf("%s", keys[ddbska.PartitionKeyAttribute])),
		}
	}
	output = &kinesis.GetRecordsOutput{
		NextShardIterator:  streamsOut.NextShardIterator,
		Records:            records,
		MillisBehindLatest: aws.Int64(0),
	}
	return
}
