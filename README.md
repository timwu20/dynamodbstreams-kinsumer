# dynamodbstreams-kinsumer

Provides the same functionality as DynamoDB Streams Kinesis Adapter but written in Go.  This library does not depend on the Java MultiLangDaemon.  Consuming and checkpointing is handled by [Kinsumer](github.com/twitchscience/kinsumer). 

## Example
See `kinsumer_integration_test.go` for an example
