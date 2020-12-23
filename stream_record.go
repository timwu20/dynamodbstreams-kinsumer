package dynamodbkinsumer

import (
	"time"
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
