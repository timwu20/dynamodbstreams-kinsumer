package dynamodbkinsumer

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/twitchscience/kinsumer"
)

func TestNew(t *testing.T) {
	type args struct {
		tableName       string
		partitionKey    string
		applicationName string
		clientName      string
		config          kinsumer.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *DynamoDBStreamsKinsumer
		wantErr bool
	}{
		{
			name: "session error",
			args: args{
				tableName:       "someTableName",
				partitionKey:    "somePartitionKey",
				applicationName: "someApplicationName",
				clientName:      "someClientName",
				config:          kinsumer.NewConfig(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.tableName, tt.args.partitionKey, tt.args.applicationName, tt.args.clientName, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

type fakeKinesisAPI struct {
	*kinesis.Kinesis
}
type fakeDynamoDBAPI struct {
	*dynamodb.DynamoDB
}

func TestNewWithInterfaces(t *testing.T) {
	expected, err := kinsumer.NewWithInterfaces(&fakeKinesisAPI{}, &fakeDynamoDBAPI{}, "someStreamARN", "someApplicationName", "someClientName", kinsumer.NewConfig())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	type args struct {
		kinesis         kinesisiface.KinesisAPI
		dynamodb        dynamodbiface.DynamoDBAPI
		streamsAPI      dynamodbstreamsiface.DynamoDBStreamsAPI
		tableName       string
		applicationName string
		clientName      string
		config          kinsumer.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *DynamoDBStreamsKinsumer
		wantErr bool
	}{
		{
			name: "one stream",
			args: args{
				kinesis:  &fakeKinesisAPI{},
				dynamodb: &fakeDynamoDBAPI{},
				streamsAPI: &fakeStreamsAPI{
					Streams: []*dynamodbstreams.Stream{
						{
							StreamArn: aws.String("someStreamARN"),
						},
					},
				},
				tableName:       "someTableName",
				applicationName: "someApplicationName",
				clientName:      "someClientName",
				config:          kinsumer.NewConfig(),
			},
			want: &DynamoDBStreamsKinsumer{
				Kinsumer: expected,
			},
		},
		{
			name: "no streams",
			args: args{
				kinesis:         &fakeKinesisAPI{},
				dynamodb:        &fakeDynamoDBAPI{},
				streamsAPI:      &fakeStreamsAPI{},
				tableName:       "someTableName",
				applicationName: "someApplicationName",
				clientName:      "someClientName",
				config:          kinsumer.NewConfig(),
			},
			wantErr: true,
		},
		{
			name: "ListStreams error",
			args: args{
				kinesis:  &fakeKinesisAPI{},
				dynamodb: &fakeDynamoDBAPI{},
				streamsAPI: &fakeStreamsAPI{
					err: fmt.Errorf("yo"),
				},
				tableName:       "someTableName",
				applicationName: "someApplicationName",
				clientName:      "someClientName",
				config:          kinsumer.NewConfig(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWithInterfaces(tt.args.kinesis, tt.args.dynamodb, tt.args.streamsAPI, tt.args.tableName, tt.args.applicationName, tt.args.clientName, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWithInterfaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreUnexported(kinsumer.Kinsumer{})); diff != "" {
				t.Errorf("NewWithInterfaces() diff: %s", diff)
			}
		})
	}
}

type fakeKinsumer struct {
	data []byte
	err  error
}

func (fk *fakeKinsumer) Next() ([]byte, error) {
	return fk.data, fk.err
}
func TestDynamoDBStreamsKinsumerNext(t *testing.T) {
	jsonStreamRecord, err := json.Marshal(StreamRecord{})
	if err != nil {
		t.Errorf("%v", err)
	}
	type fields struct {
		Kinsumer *kinsumer.Kinsumer
	}
	type args struct {
		k kinsumerNext
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantStreamRecord *StreamRecord
		wantErr          bool
	}{
		{
			name: "happy path",
			args: args{
				&fakeKinsumer{
					data: jsonStreamRecord,
				},
			},
			wantStreamRecord: &StreamRecord{},
		},
		{
			name: "Next error",
			args: args{
				&fakeKinsumer{
					err: fmt.Errorf("yo"),
				},
			},
			wantErr: true,
		},
		{
			name: "json.Unmarshal error",
			args: args{
				&fakeKinsumer{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ddbsk := &DynamoDBStreamsKinsumer{
				Kinsumer: tt.fields.Kinsumer,
			}
			gotStreamRecord, err := ddbsk.next(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("DynamoDBStreamsKinsumer.next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStreamRecord, tt.wantStreamRecord) {
				t.Errorf("DynamoDBStreamsKinsumer.next() = %v, want %v", gotStreamRecord, tt.wantStreamRecord)
			}
		})
	}
}
