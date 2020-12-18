package dynamodbkinsumer

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/twitchscience/kinsumer"
)

type DynamoDBStreamsKinsumer struct {
	*kinsumer.Kinsumer
}

// New returns a Kinsumer Interface with default kinesis and dynamodb instances, to be used in ec2 instances to get default auth and config
func New(tableName, partitionKey, applicationName, clientName string, config kinsumer.Config) (*DynamoDBStreamsKinsumer, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	return NewWithSession(s, tableName, partitionKey, applicationName, clientName, config)
}

// NewWithSession should be used if you want to override the Kinesis and Dynamo instances with a non-default aws session
func NewWithSession(session *session.Session, tableName, partitionKey, applicationName, clientName string, config kinsumer.Config) (*DynamoDBStreamsKinsumer, error) {
	s := dynamodbstreams.New(session)
	k := &DynamoDBStreamsKinesisAdapter{s, partitionKey}
	d := dynamodb.New(session)
	return NewWithInterfaces(k, d, s, tableName, partitionKey, applicationName, clientName, config)
}

// NewWithInterfaces allows you to override the Kinesis and Dynamo instances for mocking or using a local set of servers
func NewWithInterfaces(kinesis kinesisiface.KinesisAPI, dynamodb dynamodbiface.DynamoDBAPI, streamsAPI dynamodbstreamsiface.DynamoDBStreamsAPI, tableName, partitionKey, applicationName, clientName string, config kinsumer.Config) (*DynamoDBStreamsKinsumer, error) {
	listStreamsOutput, err := streamsAPI.ListStreams(&dynamodbstreams.ListStreamsInput{
		TableName: &tableName,
	})
	if err != nil {
		return nil, err
	}
	if len(listStreamsOutput.Streams) == 0 {
		err = fmt.Errorf("no streams found for table: %s", tableName)
		return nil, err
	}
	// streamName is first stream for table
	streamName := *listStreamsOutput.Streams[0].StreamArn
	k, err := kinsumer.NewWithInterfaces(kinesis, dynamodb, streamName, applicationName, clientName, config)
	if err != nil {
		return nil, err
	}
	return &DynamoDBStreamsKinsumer{
		k,
	}, nil
}

func (ddbsk *DynamoDBStreamsKinsumer) Next() (streamRecord *StreamRecord, err error) {
	data, err := ddbsk.Kinsumer.Next()
	if err != nil {
		return
	}
	sr := StreamRecord{}
	err = json.Unmarshal(data, &sr)
	if err != nil {
		return
	}
	streamRecord = &sr
	return
}
