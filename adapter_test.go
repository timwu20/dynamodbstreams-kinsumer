package dynamodbkinsumer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type fakeStreamsAPI struct {
	// embed this struct for the interface methods
	dynamodbstreams.DynamoDBStreams
	StreamStatus string
	Shards       []*dynamodbstreams.Shard
	err          error
	Records      []*dynamodbstreams.Record
	Streams      []*dynamodbstreams.Stream
}

func (ns fakeStreamsAPI) DescribeStream(*dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error) {
	return &dynamodbstreams.DescribeStreamOutput{
		StreamDescription: &dynamodbstreams.StreamDescription{
			StreamStatus: &ns.StreamStatus,
			Shards:       ns.Shards,
		},
	}, ns.err
}
func (ns fakeStreamsAPI) GetShardIterator(*dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error) {
	return &dynamodbstreams.GetShardIteratorOutput{}, ns.err
}
func (ns fakeStreamsAPI) GetRecords(*dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error) {
	return &dynamodbstreams.GetRecordsOutput{
		Records: ns.Records,
	}, ns.err
}
func (ns fakeStreamsAPI) ListStreams(*dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error) {
	return &dynamodbstreams.ListStreamsOutput{
		Streams: ns.Streams,
	}, ns.err
}

func TestDescribeStream(t *testing.T) {
	type fields struct {
		streamsAPI            dynamodbstreamsiface.DynamoDBStreamsAPI
		partitionKeyAttribute string
	}
	type args struct {
		input *kinesis.DescribeStreamInput
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput *kinesis.DescribeStreamOutput
		wantErr    bool
	}{
		{
			name: "no shards, active stream",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					StreamStatus: dynamodbstreams.StreamStatusEnabled,
				},
				partitionKeyAttribute: "PK",
			},
			args: args{&kinesis.DescribeStreamInput{}},
			wantOutput: &kinesis.DescribeStreamOutput{
				StreamDescription: &kinesis.StreamDescription{
					HasMoreShards: aws.Bool(false),
					Shards:        []*kinesis.Shard{},
					StreamStatus:  aws.String(kinesis.StreamStatusActive),
				},
			},
			wantErr: false,
		},
		{
			name: "no shards, disabled stream",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					StreamStatus: dynamodbstreams.StreamStatusDisabled,
				},
				partitionKeyAttribute: "PK",
			},
			args: args{&kinesis.DescribeStreamInput{}},
			wantOutput: &kinesis.DescribeStreamOutput{
				StreamDescription: &kinesis.StreamDescription{
					HasMoreShards: aws.Bool(false),
					Shards:        []*kinesis.Shard{},
					StreamStatus:  aws.String(kinesis.StreamStatusDeleting),
				},
			},
			wantErr: false,
		},
		{
			name: "no shards, enabling stream",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					StreamStatus: dynamodbstreams.StreamStatusEnabling,
				},
				partitionKeyAttribute: "PK",
			},
			args: args{&kinesis.DescribeStreamInput{}},
			wantOutput: &kinesis.DescribeStreamOutput{
				StreamDescription: &kinesis.StreamDescription{
					HasMoreShards: aws.Bool(false),
					Shards:        []*kinesis.Shard{},
					StreamStatus:  aws.String(kinesis.StreamStatusCreating),
				},
			},
			wantErr: false,
		},
		{
			name: "no shards, disabling stream",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					StreamStatus: dynamodbstreams.StreamStatusDisabling,
				},
				partitionKeyAttribute: "PK",
			},
			args: args{&kinesis.DescribeStreamInput{}},
			wantOutput: &kinesis.DescribeStreamOutput{
				StreamDescription: &kinesis.StreamDescription{
					HasMoreShards: aws.Bool(false),
					Shards:        []*kinesis.Shard{},
					StreamStatus:  aws.String(kinesis.StreamStatusUpdating),
				},
			},
			wantErr: false,
		},
		{
			name: "one shard, active stream",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					StreamStatus: dynamodbstreams.StreamStatusEnabled,
					Shards: []*dynamodbstreams.Shard{
						{
							ShardId:             aws.String("someShardId"),
							SequenceNumberRange: &dynamodbstreams.SequenceNumberRange{},
						},
					},
				},
				partitionKeyAttribute: "PK",
			},
			args: args{&kinesis.DescribeStreamInput{}},
			wantOutput: &kinesis.DescribeStreamOutput{
				StreamDescription: &kinesis.StreamDescription{
					HasMoreShards: aws.Bool(false),
					Shards: []*kinesis.Shard{
						{
							ShardId:             aws.String("someShardId"),
							SequenceNumberRange: &kinesis.SequenceNumberRange{},
						},
					},
					StreamStatus: aws.String(kinesis.StreamStatusActive),
				},
			},
			wantErr: false,
		},
		{
			name: "DescribeStream error",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					StreamStatus: dynamodbstreams.StreamStatusEnabled,
					err:          fmt.Errorf("yo"),
				},
				partitionKeyAttribute: "PK",
			},
			args:       args{&kinesis.DescribeStreamInput{}},
			wantOutput: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ddbska := DynamoDBStreamsKinesisAdapter{
				streamsAPI:            tt.fields.streamsAPI,
				partitionKeyAttribute: tt.fields.partitionKeyAttribute,
			}
			gotOutput, err := ddbska.DescribeStream(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DynamoDBStreamsKinesisAdapter.DescribeStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("DynamoDBStreamsKinesisAdapter.DescribeStream() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestListShards(t *testing.T) {
	type fields struct {
		streamsAPI            dynamodbstreamsiface.DynamoDBStreamsAPI
		partitionKeyAttribute string
	}
	type args struct {
		input *kinesis.ListShardsInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *kinesis.ListShardsOutput
		wantErr bool
	}{
		{
			name: "one shard, active stream",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					StreamStatus: dynamodbstreams.StreamStatusEnabled,
					Shards: []*dynamodbstreams.Shard{
						{
							ShardId:             aws.String("someShardId"),
							SequenceNumberRange: &dynamodbstreams.SequenceNumberRange{},
						},
					},
				},
				partitionKeyAttribute: "PK",
			},
			args: args{
				input: &kinesis.ListShardsInput{},
			},
			want: &kinesis.ListShardsOutput{
				Shards: []*kinesis.Shard{
					{
						ShardId:             aws.String("someShardId"),
						SequenceNumberRange: &kinesis.SequenceNumberRange{},
					},
				},
			},
		},
		{
			name: "DescribeStream error",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					StreamStatus: dynamodbstreams.StreamStatusEnabled,
					err:          fmt.Errorf("yo"),
				},
				partitionKeyAttribute: "PK",
			},
			args:    args{&kinesis.ListShardsInput{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ddbska := DynamoDBStreamsKinesisAdapter{
				streamsAPI:            tt.fields.streamsAPI,
				partitionKeyAttribute: tt.fields.partitionKeyAttribute,
			}
			got, err := ddbska.ListShards(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DynamoDBStreamsKinesisAdapter.ListShards() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DynamoDBStreamsKinesisAdapter.ListShards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetShardIterator(t *testing.T) {
	type fields struct {
		streamsAPI            dynamodbstreamsiface.DynamoDBStreamsAPI
		partitionKeyAttribute string
	}
	type args struct {
		input *kinesis.GetShardIteratorInput
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput *kinesis.GetShardIteratorOutput
		wantErr    bool
	}{
		{
			name: "happy path",
			fields: fields{
				streamsAPI:            &fakeStreamsAPI{},
				partitionKeyAttribute: "PK",
			},
			args: args{
				input: &kinesis.GetShardIteratorInput{},
			},
			wantOutput: &kinesis.GetShardIteratorOutput{},
		},
		{
			name: "GetShardIterator error",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					err: fmt.Errorf("yo"),
				},
				partitionKeyAttribute: "PK",
			},
			args: args{
				input: &kinesis.GetShardIteratorInput{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ddbska := DynamoDBStreamsKinesisAdapter{
				streamsAPI:            tt.fields.streamsAPI,
				partitionKeyAttribute: tt.fields.partitionKeyAttribute,
			}
			gotOutput, err := ddbska.GetShardIterator(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DynamoDBStreamsKinesisAdapter.GetShardIterator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("DynamoDBStreamsKinesisAdapter.GetShardIterator() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestGetRecords(t *testing.T) {
	expectedStreamRecord := StreamRecord{
		Keys: map[string]interface{}{
			"PK": "somePK",
		},
		NewImage: map[string]interface{}{},
		OldImage: map[string]interface{}{},
	}
	streamRecordJSON, err := json.Marshal(expectedStreamRecord)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	type fields struct {
		streamsAPI            dynamodbstreamsiface.DynamoDBStreamsAPI
		partitionKeyAttribute string
	}
	type args struct {
		input *kinesis.GetRecordsInput
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput *kinesis.GetRecordsOutput
		wantErr    bool
	}{
		{
			name: "no records",
			fields: fields{
				streamsAPI:            &fakeStreamsAPI{},
				partitionKeyAttribute: "PK",
			},
			args: args{
				input: &kinesis.GetRecordsInput{},
			},
			wantOutput: &kinesis.GetRecordsOutput{
				MillisBehindLatest: aws.Int64(-1),
				Records:            []*kinesis.Record{},
			},
		},
		{
			name: "some records, limit 500",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					Records: []*dynamodbstreams.Record{
						{
							Dynamodb: &dynamodbstreams.StreamRecord{
								Keys: map[string]*dynamodb.AttributeValue{
									"PK": &dynamodb.AttributeValue{
										S: aws.String("somePK"),
									},
								},
							},
						},
					},
				},
				partitionKeyAttribute: "PK",
			},
			args: args{
				input: &kinesis.GetRecordsInput{
					Limit: aws.Int64(500),
				},
			},
			wantOutput: &kinesis.GetRecordsOutput{
				MillisBehindLatest: aws.Int64(-1),
				Records: []*kinesis.Record{
					{
						Data:         streamRecordJSON,
						PartitionKey: aws.String("somePK"),
					},
				},
			},
		},
		{
			name: "some records, limit 5000",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					Records: []*dynamodbstreams.Record{
						{
							Dynamodb: &dynamodbstreams.StreamRecord{
								Keys: map[string]*dynamodb.AttributeValue{
									"PK": &dynamodb.AttributeValue{
										S: aws.String("somePK"),
									},
								},
							},
						},
					},
				},
				partitionKeyAttribute: "PK",
			},
			args: args{
				input: &kinesis.GetRecordsInput{
					Limit: aws.Int64(5000),
				},
			},
			wantOutput: &kinesis.GetRecordsOutput{
				MillisBehindLatest: aws.Int64(-1),
				Records: []*kinesis.Record{
					{
						Data:         streamRecordJSON,
						PartitionKey: aws.String("somePK"),
					},
				},
			},
		},
		{
			name: "GetRecords error",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					err: fmt.Errorf("yo"),
				},
				partitionKeyAttribute: "PK",
			},
			args: args{
				input: &kinesis.GetRecordsInput{},
			},
			wantErr: true,
		},
		{
			name: "partition key error",
			fields: fields{
				streamsAPI: &fakeStreamsAPI{
					Records: []*dynamodbstreams.Record{
						{
							Dynamodb: &dynamodbstreams.StreamRecord{
								Keys: map[string]*dynamodb.AttributeValue{},
							},
						},
					},
				},
				partitionKeyAttribute: "PK",
			},
			args: args{
				input: &kinesis.GetRecordsInput{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ddbska := DynamoDBStreamsKinesisAdapter{
				streamsAPI:            tt.fields.streamsAPI,
				partitionKeyAttribute: tt.fields.partitionKeyAttribute,
			}
			gotOutput, err := ddbska.GetRecords(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DynamoDBStreamsKinesisAdapter.GetRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("DynamoDBStreamsKinesisAdapter.GetRecords() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
