# CloudWatch

DynamoDB publishes consumed capacity metrics to Amazon CloudWatch A

DynamoDB auto-scaling. The CloudWatch alarm invokes Application Auto Scaling to evaluate your scaling policy using SNS which issues an UpdateTable request to adjust your table's provisioned throughput.