package dynamodbkinsumer

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
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
}

func (ns fakeStreamsAPI) DescribeStream(*dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error) {
	return &dynamodbstreams.DescribeStreamOutput{
		StreamDescription: &dynamodbstreams.StreamDescription{
			StreamStatus: &ns.StreamStatus,
			Shards:       ns.Shards,
		},
	}, ns.err
}

func TestDescribeStream(t *testing.T) {
	type fields struct {
		streamsAPI            dynamodbstreamsiface.DynamoDBStreamsAPI
		PartitionKeyAttribute string
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
				PartitionKeyAttribute: "PK",
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
				PartitionKeyAttribute: "PK",
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
				PartitionKeyAttribute: "PK",
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
				PartitionKeyAttribute: "PK",
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
			name: "one shards, active stream",
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
				PartitionKeyAttribute: "PK",
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
				PartitionKeyAttribute: "PK",
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
				PartitionKeyAttribute: tt.fields.PartitionKeyAttribute,
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
