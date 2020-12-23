package dynamodbkinsumer

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
)

// StreamRecord is unmarshalled dynamodbstreams.StreamRecord
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

func (sr streamRecord) StreamRecord() (b []byte, err error) {
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
