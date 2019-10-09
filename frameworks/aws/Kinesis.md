# Kinesis

[Key Concepts](https://docs.aws.amazon.com/streams/latest/dev/key-concepts.html)

You can add data to an Amazon Kinesis data stream via PutRecord and PutRecords operations, Amazon Kinesis Producer Library (KPL), or Amazon Kinesis Agent.

**note how KPL and kinesis agent are different.KPL is only available on Java and C+

## Kinesis Streams
![](../../assets/amazon_kinesis_streams.png)

### Kinesis Data Streams

Kineses Data Streams record is an _instance_ of the record data structure defined by the KDS service API. KPL user record and KDS stream record are **different** A KDS record contains:
- partition key
- sequence number
- blob of data.

A Kinesis data stream is a set of shards. Each **shard** has a sequence of **data records**. Each **data record** has a sequence number that is assigned by Kinesis Data Streams

```
- Shard 1
    - Data Records: [{seq#:1}, {seq#:2}, {seq#:3}]

- Shard 2
    - Data Records: [{seq#:4}, {seq#:5}, {seq#:6}]
- Shard 3 ...
```

**Data Streams API** -  Data Streams API can be consumer to the Data Stream but cannot deaggregate the data.

- Future Objects: Kinesis Streams provide capabilities to use Future objects to validate UserRecords. No need to complicate code by storing memory/transient storage. Examine failures using the Future objects that are returned from addUserRecord method.

- Automatic and configurable retry mechanism

- Time-to-live: on records if records could not be inserted after RecordMaxBufferedTime, time to-live can be extended if UserRecords could not be inserted into stream in time.

#### Shards

Each shard has a sequence of data records. A shard is a uniquely identified sequence of data records in a stream. A stream is composed of one or more shards, each of which has a **fixed unit of capacity**.

. Each shard can support up to **5 transactions per second** for **reads**, up to a maximum total data read rate of **2 MB per second** and up to **1,000 records per second** for **writes**, up to a maximum total data write rate of **1 MB per second** (including partition keys). 

The total capacity of the stream is the sum of the capacities of its shards.

- **Merge cold shards** to make use of their unused capacity.
- Merge the shards that receive **less data**

#### Partition Key
A **partition key** is used to **group data by shard** within a stream.

Kinesis Data Streams segregates the **data records** belonging to a **stream** into **multiple shards**. 

It uses the **partition key** that is associated with each data record to determine **which shard** a given **data record** belongs to.

Partition keys are Unicode strings with a **maximum** length limit of **256 bytes**.

#### Sequence Number

#### Retention Period
The retention period is set at the stream level. The **retention period** is the length of time that data records are accessible after they are added to the stream. 

`Default retention period`: **24 hours** after creation
You can increase the retention perioud up to **168 hours (7 days)**

#### Data Records
A data record is the unit of data stored in a Kinesis data stream. 

Composed of:
- Sequence number
- Partition key
- Data blob (immutable sequence of bytes)
    - KDS does not inspect/interpret/modify the data blob
    - Can be up to 1 MB

#### Server-side Encryption

Kinesis Data Streams can automatically encrypt sensitive data as a producer enters it into a stream.

Kinesis Data Streams uses **AWS KMS master keys** for encryption
    - `Note`: Using server-side encryption incurs AWS Key Management Service (AWS KMS) costs

### Kinesis Video Streams

### Producers

Producers are the systems sending data.
You can send data using:
- KPL (Kinesis Producer Library)
- Kinesis Agent
- Kinesis API

### Consumers

Consumer can be implemented with

- Kinesis Client Library ( KCL )
    - This is a pre-built library that helps you easily build Amazon Kinesis Applications for reading and processing data from an Amazon Kinesis stream. This library handles complex issues such as adapting to changes in stream volume, load-balancing streaming data, coordinating distributed services, and processing data with fault-tolerance, enabling you to focus on business logic while building applications.
- Kinesis Connector Library
    - Used to connect Kinesis with other services on AWS. KCL is required to use this Library
- Kinesis Data Streams API (but KCL is recommended)
- KCL uses a unique Amazon DynamoDB table to keep track of the application's state.
- KCL uses name of Kinesis Data Streams application to create the name of the table, each application name must be unique.

### Kinesis Connector Library 

Helps java developers integrate Kinesis Streams with other AWS services.

The library provides connectors to various AWS services including S3. 

Each Amazon Kinesis connector application is a pipeline that understands how records from Kinesis Stream will be handled.

#### Kinesis agent

Kinesis Agent is a stand-alone Java application that can easily collect and send data to Kinesis Data Streams. 

The agent can continuously monitor set of fles (more for log fles) and Aggregation of data is not possible.

#### Kinesis Data Analytics

**Stagger Windows**: A query that aggregates data using keyed time-based windows that open as data arrives. The keys allow for multiple overlapping windows. This is the recommended way to aggregate data using time-based windows, because Stagger Windows reduce late or out-of-order data compared to Tumbling windows.

**Tumbling Windows**: A query that aggregates data using distinct time-based windows that open and close at regular intervals.

**Sliding Windows**: A query that aggregates data continuously, using a fixed time or rowcount interval.

**Continuous Query**: query over a stream executes continuously over streaming data. Enables scenarios such as ability for applications to query the stream and generate alerts.