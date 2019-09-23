# DynamoDB

DynamoDB supports both key-value and document data models. This
enables DynamoDB to have a fexible schema, so each row can have any number of columns at any point in time.

DynamoDB provides two read/write capacity modes for each table: ondemand and provisioned. For workloads that are less predictable for which you are unsure that you will have high utilization, on-demandcapacity mode takes care of managing capacity for you, and you only pay for what you consume. Tables using provisioned capacity mode require you to set read and write capacity.

DynamoDB global tables replicate your data automatically across your choice of AWS Regions and automatically scale capacity to accommodate your workloads. With global tables, your globally distributed requests can access data locally in the selected regions to get single- digit millisecond read and write performance


### Auto-Scaling

Metrics are published to CloudWatch

If the table's consumed capacity exceeds your target utilization (or falls below the target) for a specifc length of time, Amazon CloudWatch triggers an alarm using SNS.

The CloudWatch alarm invokes Application Auto Scaling to evaluate your scaling policy using SNS which issues an UpdateTable request to adjust your table's provisioned throughput.

### Designing a Replica Table
- The table must have the **same partition key** as all of the other replicas.
- The table must have the **same write capacity** but **not** the same **read capacity**
- The table must have **DynamoDB Streams enabled**, with the stream containing both the new and the old images of the item.
- The table and global secondary indexes must have the same name.
- The global secondary indexes must have the same partition key and sort key (if present).
- None of the new or existing replica tables in the global table can contain any data.
("Local secondary" names is made up)


### rWCU (replicated write request units)
If you enable on-demand mode on a global table, and you perform **10 writes** to a **local table** that is replicated in **two additional Regions**, you will consume **60 write request units** (10 + 10 + 10 = 30; 30 x 2 = 60).
