# dynamodbstreams-kinsumer

[![timwu20](https://circleci.com/gh/timwu20/dynamodbstreams-kinsumer.svg?style=shield)](https://app.circleci.com/pipelines/github/timwu20/dynamodbstreams-kinsumer) [![codecov](https://codecov.io/gh/timwu20/dynamodbstreams-kinsumer/branch/main/graph/badge.svg?token=X8J1NRB0F5)](https://codecov.io/gh/timwu20/dynamodbstreams-kinsumer)

Provides the same functionality as DynamoDB Streams Kinesis Adapter but written in Go.  This library does not depend on the Java MultiLangDaemon.  Consuming and checkpointing is handled by [Kinsumer](https://github.com/twitchscience/kinsumer). 

## Example
See `kinsumer_integration_test.go` for an example
