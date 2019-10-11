# AWS Monitoring / Logging

CloudWatch vs CloudTrail

### DynamoDB

DynamoDB publishes consumed **capacity metrics** to Amazon **CloudWatch**.

##### Auto-Scaling

The **CloudWatch** alarm invokes **Application Auto Scaling** to evaluate your scaling policy using **SNS** which issues an UpdateTable request to adjust your table's provisioned throughput.

### Kinesis

- Metrics configured for streams are automatically collected and published to **CloudWatch** every _minute_. These metrics are published by **CloudWatch**

- The **kinesis agent** publishes _custom_ **CloudWatch** metrics with a namespace "AWSKinesisAgent"

- The **kinesis agent** publishes: 
    - Bytes sent
    - Number of records that returned failure
    - Number of calls to PutRecords that resulted in service error

### Elastic Transcoder

Monitor health of transcoding workflows using **CloudWatch**