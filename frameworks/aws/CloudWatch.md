# CloudWatch


## DynamoDB

DynamoDB publishes consumed capacity metrics to Amazon CloudWatch.

#### Auto-Scaling

The CloudWatch alarm invokes Application Auto Scaling to evaluate your scaling policy using SNS which issues an UpdateTable request to adjust your table's provisioned throughput.

##### Alarm
If the table's consumed capacity exceeds your target utilization (or falls below thetarget) for a specifc length of time, AWS CloudWatch triggers an alarm using SNS.

The alarm invokes Application Auto Scaling to evaluate your scaling policy and process
request using UpdateTable