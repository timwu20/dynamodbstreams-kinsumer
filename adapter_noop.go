package dynamodbkinsumer

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func (ddbska DynamoDBStreamsKinesisAdapter) AddTagsToStream(*kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")

}
func (ddbska DynamoDBStreamsKinesisAdapter) AddTagsToStreamWithContext(aws.Context, *kinesis.AddTagsToStreamInput, ...request.Option) (*kinesis.AddTagsToStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")

}
func (ddbska DynamoDBStreamsKinesisAdapter) AddTagsToStreamRequest(*kinesis.AddTagsToStreamInput) (*request.Request, *kinesis.AddTagsToStreamOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) CreateStream(*kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) CreateStreamWithContext(aws.Context, *kinesis.CreateStreamInput, ...request.Option) (*kinesis.CreateStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) CreateStreamRequest(*kinesis.CreateStreamInput) (*request.Request, *kinesis.CreateStreamOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) DecreaseStreamRetentionPeriod(*kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DecreaseStreamRetentionPeriodWithContext(aws.Context, *kinesis.DecreaseStreamRetentionPeriodInput, ...request.Option) (*kinesis.DecreaseStreamRetentionPeriodOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DecreaseStreamRetentionPeriodRequest(*kinesis.DecreaseStreamRetentionPeriodInput) (*request.Request, *kinesis.DecreaseStreamRetentionPeriodOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) DeleteStream(*kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DeleteStreamWithContext(aws.Context, *kinesis.DeleteStreamInput, ...request.Option) (*kinesis.DeleteStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DeleteStreamRequest(*kinesis.DeleteStreamInput) (*request.Request, *kinesis.DeleteStreamOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) DeregisterStreamConsumer(*kinesis.DeregisterStreamConsumerInput) (*kinesis.DeregisterStreamConsumerOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DeregisterStreamConsumerWithContext(aws.Context, *kinesis.DeregisterStreamConsumerInput, ...request.Option) (*kinesis.DeregisterStreamConsumerOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DeregisterStreamConsumerRequest(*kinesis.DeregisterStreamConsumerInput) (*request.Request, *kinesis.DeregisterStreamConsumerOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) DescribeLimits(*kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DescribeLimitsWithContext(aws.Context, *kinesis.DescribeLimitsInput, ...request.Option) (*kinesis.DescribeLimitsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DescribeLimitsRequest(*kinesis.DescribeLimitsInput) (*request.Request, *kinesis.DescribeLimitsOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStreamWithContext(aws.Context, *kinesis.DescribeStreamInput, ...request.Option) (*kinesis.DescribeStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStreamRequest(*kinesis.DescribeStreamInput) (*request.Request, *kinesis.DescribeStreamOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStreamPages(*kinesis.DescribeStreamInput, func(*kinesis.DescribeStreamOutput, bool) bool) error {
	panic(fmt.Errorf("not implemented"))
}
func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStreamPagesWithContext(aws.Context, *kinesis.DescribeStreamInput, func(*kinesis.DescribeStreamOutput, bool) bool, ...request.Option) error {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStreamConsumer(*kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStreamConsumerWithContext(aws.Context, *kinesis.DescribeStreamConsumerInput, ...request.Option) (*kinesis.DescribeStreamConsumerOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStreamConsumerRequest(*kinesis.DescribeStreamConsumerInput) (*request.Request, *kinesis.DescribeStreamConsumerOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStreamSummary(*kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStreamSummaryWithContext(aws.Context, *kinesis.DescribeStreamSummaryInput, ...request.Option) (*kinesis.DescribeStreamSummaryOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DescribeStreamSummaryRequest(*kinesis.DescribeStreamSummaryInput) (*request.Request, *kinesis.DescribeStreamSummaryOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) DisableEnhancedMonitoring(*kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DisableEnhancedMonitoringWithContext(aws.Context, *kinesis.DisableEnhancedMonitoringInput, ...request.Option) (*kinesis.EnhancedMonitoringOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) DisableEnhancedMonitoringRequest(*kinesis.DisableEnhancedMonitoringInput) (*request.Request, *kinesis.EnhancedMonitoringOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) EnableEnhancedMonitoring(*kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) EnableEnhancedMonitoringWithContext(aws.Context, *kinesis.EnableEnhancedMonitoringInput, ...request.Option) (*kinesis.EnhancedMonitoringOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) EnableEnhancedMonitoringRequest(*kinesis.EnableEnhancedMonitoringInput) (*request.Request, *kinesis.EnhancedMonitoringOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) GetRecordsWithContext(aws.Context, *kinesis.GetRecordsInput, ...request.Option) (*kinesis.GetRecordsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) GetRecordsRequest(*kinesis.GetRecordsInput) (*request.Request, *kinesis.GetRecordsOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) GetShardIteratorWithContext(aws.Context, *kinesis.GetShardIteratorInput, ...request.Option) (*kinesis.GetShardIteratorOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) GetShardIteratorRequest(*kinesis.GetShardIteratorInput) (*request.Request, *kinesis.GetShardIteratorOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) IncreaseStreamRetentionPeriod(*kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) IncreaseStreamRetentionPeriodWithContext(aws.Context, *kinesis.IncreaseStreamRetentionPeriodInput, ...request.Option) (*kinesis.IncreaseStreamRetentionPeriodOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) IncreaseStreamRetentionPeriodRequest(*kinesis.IncreaseStreamRetentionPeriodInput) (*request.Request, *kinesis.IncreaseStreamRetentionPeriodOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) ListShardsWithContext(aws.Context, *kinesis.ListShardsInput, ...request.Option) (*kinesis.ListShardsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) ListShardsRequest(*kinesis.ListShardsInput) (*request.Request, *kinesis.ListShardsOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) ListStreamConsumers(*kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) ListStreamConsumersWithContext(aws.Context, *kinesis.ListStreamConsumersInput, ...request.Option) (*kinesis.ListStreamConsumersOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) ListStreamConsumersRequest(*kinesis.ListStreamConsumersInput) (*request.Request, *kinesis.ListStreamConsumersOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) ListStreamConsumersPages(*kinesis.ListStreamConsumersInput, func(*kinesis.ListStreamConsumersOutput, bool) bool) error {
	return fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) ListStreamConsumersPagesWithContext(aws.Context, *kinesis.ListStreamConsumersInput, func(*kinesis.ListStreamConsumersOutput, bool) bool, ...request.Option) error {
	return fmt.Errorf("not implemented")
}

func (ddbska DynamoDBStreamsKinesisAdapter) ListStreams(*kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) ListStreamsWithContext(aws.Context, *kinesis.ListStreamsInput, ...request.Option) (*kinesis.ListStreamsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) ListStreamsRequest(*kinesis.ListStreamsInput) (*request.Request, *kinesis.ListStreamsOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) ListStreamsPages(*kinesis.ListStreamsInput, func(*kinesis.ListStreamsOutput, bool) bool) error {
	return fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) ListStreamsPagesWithContext(aws.Context, *kinesis.ListStreamsInput, func(*kinesis.ListStreamsOutput, bool) bool, ...request.Option) error {
	return fmt.Errorf("not implemented")
}

func (ddbska DynamoDBStreamsKinesisAdapter) ListTagsForStream(*kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) ListTagsForStreamWithContext(aws.Context, *kinesis.ListTagsForStreamInput, ...request.Option) (*kinesis.ListTagsForStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) ListTagsForStreamRequest(*kinesis.ListTagsForStreamInput) (*request.Request, *kinesis.ListTagsForStreamOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) MergeShards(*kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) MergeShardsWithContext(aws.Context, *kinesis.MergeShardsInput, ...request.Option) (*kinesis.MergeShardsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) MergeShardsRequest(*kinesis.MergeShardsInput) (*request.Request, *kinesis.MergeShardsOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) PutRecord(*kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) PutRecordWithContext(aws.Context, *kinesis.PutRecordInput, ...request.Option) (*kinesis.PutRecordOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) PutRecordRequest(*kinesis.PutRecordInput) (*request.Request, *kinesis.PutRecordOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) PutRecords(*kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) PutRecordsWithContext(aws.Context, *kinesis.PutRecordsInput, ...request.Option) (*kinesis.PutRecordsOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) PutRecordsRequest(*kinesis.PutRecordsInput) (*request.Request, *kinesis.PutRecordsOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) RegisterStreamConsumer(*kinesis.RegisterStreamConsumerInput) (*kinesis.RegisterStreamConsumerOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) RegisterStreamConsumerWithContext(aws.Context, *kinesis.RegisterStreamConsumerInput, ...request.Option) (*kinesis.RegisterStreamConsumerOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) RegisterStreamConsumerRequest(*kinesis.RegisterStreamConsumerInput) (*request.Request, *kinesis.RegisterStreamConsumerOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) RemoveTagsFromStream(*kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) RemoveTagsFromStreamWithContext(aws.Context, *kinesis.RemoveTagsFromStreamInput, ...request.Option) (*kinesis.RemoveTagsFromStreamOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) RemoveTagsFromStreamRequest(*kinesis.RemoveTagsFromStreamInput) (*request.Request, *kinesis.RemoveTagsFromStreamOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) SplitShard(*kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) SplitShardWithContext(aws.Context, *kinesis.SplitShardInput, ...request.Option) (*kinesis.SplitShardOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) SplitShardRequest(*kinesis.SplitShardInput) (*request.Request, *kinesis.SplitShardOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) StartStreamEncryption(*kinesis.StartStreamEncryptionInput) (*kinesis.StartStreamEncryptionOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) StartStreamEncryptionWithContext(aws.Context, *kinesis.StartStreamEncryptionInput, ...request.Option) (*kinesis.StartStreamEncryptionOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) StartStreamEncryptionRequest(*kinesis.StartStreamEncryptionInput) (*request.Request, *kinesis.StartStreamEncryptionOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) StopStreamEncryption(*kinesis.StopStreamEncryptionInput) (*kinesis.StopStreamEncryptionOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) StopStreamEncryptionWithContext(aws.Context, *kinesis.StopStreamEncryptionInput, ...request.Option) (*kinesis.StopStreamEncryptionOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) StopStreamEncryptionRequest(*kinesis.StopStreamEncryptionInput) (*request.Request, *kinesis.StopStreamEncryptionOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) SubscribeToShard(*kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) SubscribeToShardWithContext(aws.Context, *kinesis.SubscribeToShardInput, ...request.Option) (*kinesis.SubscribeToShardOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) SubscribeToShardRequest(*kinesis.SubscribeToShardInput) (*request.Request, *kinesis.SubscribeToShardOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) UpdateShardCount(*kinesis.UpdateShardCountInput) (*kinesis.UpdateShardCountOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) UpdateShardCountWithContext(aws.Context, *kinesis.UpdateShardCountInput, ...request.Option) (*kinesis.UpdateShardCountOutput, error) {
	return nil, fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) UpdateShardCountRequest(*kinesis.UpdateShardCountInput) (*request.Request, *kinesis.UpdateShardCountOutput) {
	panic(fmt.Errorf("not implemented"))
}

func (ddbska DynamoDBStreamsKinesisAdapter) WaitUntilStreamExists(*kinesis.DescribeStreamInput) error {
	return fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) WaitUntilStreamExistsWithContext(aws.Context, *kinesis.DescribeStreamInput, ...request.WaiterOption) error {
	return fmt.Errorf("not implemented")
}

func (ddbska DynamoDBStreamsKinesisAdapter) WaitUntilStreamNotExists(*kinesis.DescribeStreamInput) error {
	return fmt.Errorf("not implemented")
}
func (ddbska DynamoDBStreamsKinesisAdapter) WaitUntilStreamNotExistsWithContext(aws.Context, *kinesis.DescribeStreamInput, ...request.WaiterOption) error {
	return fmt.Errorf("not implemented")
}
