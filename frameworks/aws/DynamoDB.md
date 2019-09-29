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


### Indexes

Allow for quicker lookups, way to query it by primary key (hash value).

**Sparse** - For any item in a table, DynamoDB writes a corresponding index entry only if the index sort key value is present in the item. If the sort key doesn't appear in every table item, the index is said to be sparse. Sparse indexes are useful for queries over a small subsection of a table. For example, suppose that you have a table where you store all your customer orders, with the following key attributes:

```
Partition key: CustomerId
Sort key: OrderId
```

**Global secondary index** — An index with a partition key and a sort key that can be different from those on the base table. A global secondary index is considered "global" because queries on the index can span all of the data in the base table, across all partitions. A global secondary index has no size limitations and has its own provisioned throughput settings for read and write activity that are separate from those of the table.

**Local secondary index** — An index that has the same partition key as the base table, but a different sort key. A local secondary index is "local" in the sense that every partition of a local secondary index is scoped to a base table partition that has the same partition key value. As a result, the total size of indexed items for any one partition key value can't exceed 10 GB. Also, a local secondary index shares provisioned throughput settings for read and write activity with the table it is indexing.