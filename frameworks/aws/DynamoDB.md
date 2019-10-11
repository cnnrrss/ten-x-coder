# DynamoDB

- Fully managed NoSql DB
- Supports key-value or document data models. 
    - Enables flexible schema
- No practical storage limits
    - 10 trillion requests per day
    - peaks > 20mil per second
    - petabytes of storage)
- SSD
- Fully resilient and highly available 
    - Automated global replication across 3 regions
- Performance scales - in a linear way
    - specify how many reads/writes you want per second
- Fully integrated with IAM
- Microsecond latency with DynamoDB Accelerator (DAX)
    - In-mem cache
- Real-time data processing with DynamoDB Streams
    - time-ordered sequence of item-level modifications, up to 24 hours

Row == Item in DynamoDB

Rows can have different elements, and different number of elements. No Schema at the table level

#### Burst Capacity
DynamoDB provides some flex in your per-partition throughput provisioning by providing burst capacity. DynamoDB reserves a portion of that unused capacity for later burssts of throughput to handle usage spikes.

- Use **Global Tables**, a fully managed solution across multiple regions, multi-master databases

## Serverless
- Read/write capacity modes
- On-demand mode
- Auto scaling
- Change tracking with triggers

## Enterprise Grade 

- **ACID transactions** - native, server-side support for transations, simplifying the developer experience and enabling enterprise scale/perf benefits to broad set of mission-critical workloads.

- **Encryption at rest** (by default). Uses keys stored in AWS KMS

- **Point-in-time recovery** (PITR) can be enabled for DynamoDB. PITR provides continuous backups, can restore that table to any point in time _to the second_ during the preceding 35 days.
- **On-demand backup and restore**. (not encrypted by default)

### TTL

- Background job checks the TTL attribute of _items_ to see if they are expired. 
- If epoch time val stored in the attribute is less than current time, item is marked as _expired_ and subesquently deleted.
- The processing of expiry and deletion takes place automatically in the backround and **does not** affect read or write traffic to the table.

### Tables

DynamoDB is a collection of Tables
Tables are the **highest** level structure within a Database
Its on tables that you specify the performance requirements

- Write Capacity Units - Number of 1KB blocks / second
- Read Capacity Units - Number of 4KB blocks / second

Eventually Consistent Reads ( Default )
Strongly Consistent (more cost)

This enables DynamoDB to have a fexible schema, so each row can have any number of columns at any point in time.


DynamoDB provides two read/write capacity modes for each table: **on-demand** and **provisioned**. 

For workloads that are less predictable for which you are unsure that you will have high utilization, on-demand capacity mode takes care of managing capacity for you, and you only pay for what you consume. Tables using provisioned capacity mode require you to set read and write capacity.

### Auto-Scaling

Metrics are published to CloudWatch

If the table's consumed capacity exceeds your target utilization (or falls below the target) for a specifc length of time, Amazon CloudWatch triggers an alarm using SNS.

The CloudWatch alarm invokes Application Auto Scaling to evaluate your scaling policy using SNS which issues an UpdateTable request to adjust your table's provisioned throughput.

### Designing a Replica Table / Global Table

DynamoDB **global tables** replicate your data automatically across your choice of AWS Regions and automatically scale capacity to accommodate your workloads. With global tables, your globally distributed requests can access data locally in the selected regions to get single- digit millisecond read and write performance.

- The tables _must_ have the **same partition key** as all of the other replicas.
- The tables _must_ have the **same write capacity** but **not read capacity**
- The table **must** have **DynamoDB Streams enabled**, with the stream containing bot the new and old images of the item.
- The **global secondary indexes** _must_ have the **same** **partition key** and **sort key**
- The table and the **global secondary** names across all replicas must be the same.
- None of the new or existing replica tables in the global table can contain any data.
("Local secondary" names is made up)

### rWCU (replicated write request units)
If you enable on-demand mode on a global table, and you perform **10 writes** to a **local table** that is replicated in **two additional Regions**, you will consume **60 write request units** (10 + 10 + 10 = 30; 30 x 2 = 60).

### Indexes

Allow for quicker lookups, way to query it by primary key (hash value).

**Sparse** - For any item in a table, DynamoDB writes a corresponding index entry only if the index sort key value is present in the item. If the sort key doesn't appear in every table item, the index is said to be sparse. _Sparse indexes are useful for queries over a small subsection of a table._ For example, suppose that you have a table where you store all your customer orders, with the following key attributes:

```
Partition key: CustomerId
Sort key: OrderId
```

**Global secondary index** — An index with a partition key and a sort key that can be different from those on the base table. A global secondary index is considered "global" because queries on the index can span all of the data in the base table, across all partitions. A global secondary index has no size limitations and has its own provisioned throughput settings for read and write activity that are separate from those of the table.

**Local secondary index** — An index that has the same partition key as the base table, but a different sort key. A local secondary index is "local" in the sense that every partition of a local secondary index is scoped to a base table partition that has the same partition key value. As a result, the total size of indexed items for any one partition key value can't exceed 10 GB. Also, a local secondary index shares provisioned throughput settings for read and write activity with the table it is indexing.