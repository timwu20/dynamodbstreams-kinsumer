//+build integration

package dynamodbkinsumer

import (
	"encoding/json"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/satori/uuid"
	"github.com/twitchscience/kinsumer"
)

func TestDynamoDBStreamsKinsumer(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	})
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	k, err := NewWithSession(
		sess,
		"someTable.dev",
		"PK",
		"someAppName",
		uuid.NewV4().String(),
		kinsumer.NewConfig(),
	)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	err = k.CreateRequiredTables()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	err = k.Run()
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			record, err := k.Next()
			if err != nil {
				log.Fatalf("k.Next returned error %v", err)
			}
			if record != nil {
				streamRecord := StreamRecord{}
				err := json.Unmarshal(record, &streamRecord)
				if err != nil {
					log.Fatalf("%v", err)
				}
				log.Printf("record: %+v\n", streamRecord)
			} else {
				return
			}
		}
	}()

	time.Sleep(5 * time.Minute)
	k.Stop()
	wg.Wait()
}
