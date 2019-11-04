# Kinesis

[Key Concepts](https://docs.aws.amazon.com/streams/latest/dev/key-concepts.html)

### Producers

Producers are the systems sending data.
You can send data using:
- **Kinesis API & PutRecord** and PutRecords operations
- **Amazon Kinesis Producer Library** (KPL)
- **Amazon Kinesis Agent**

**note how KPL and kinesis agent are different. KPL is only available on Java and C+

### Consumers

Kinesis Stream consumers can be implemented with:

- **Kinesis Client Library ( KCL )**
    - This is a pre-built library that helps you easily build Amazon Kinesis Applications for reading and processing data from an Amazon Kinesis stream.
    - This library handles complex issues such as adapting to changes in stream volume, load-balancing streaming data, coordinating distributed services, and processing data with fault-tolerance, enabling you to focus on business logic while building applications.
- **Kinesis Connector Library**
    - Used to connect Kinesis with other services on AWS. KCL is required to use this Library
- **Kinesis Data Streams API** (but KCL is recommended)

## Kinesis Streams
![](../../assets/amazon_kinesis_streams.png)

_TODO_...

### Kinesis Data Streams

Kineses Data Streams record is an _instance_ of the record data structure defined by the KDS service API. KPL user record and KDS stream record are **different** A KDS record contains:
- partition key
- sequence number
- blob of data.

A Kinesis data stream is a set of shards. Each **shard** has a sequence of **data records**. Each **data record** has a **sequence number** that is assigned by Kinesis Data Streams

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

Each shard has a _sequence_ of **data records**. A shard is a uniquely identified sequence of data records in a stream. A stream is composed of one or more shards, each of which has a **fixed unit of capacity**.

- Each shard can support up to **5 transactions per second** for **reads**, up to a maximum total data read rate of **2 MB per second**
- Each shard can support up to **1,000 records per second** for **writes**, up to a maximum total data write rate of **1 MB per second** (including partition keys).
- There is no upper limit on the # of shards you can have in a stream or account.

The total capacity of the stream is the sum of the capacities of its shards.

- **Merge cold shards** to make use of their unused capacity.
- Merge the shards that receive **less data**

#### Additional Limits

- There is no upper limit on the # of streams you can have in an account.
- If you scaled to 5,000 shards, the stream can ingest 5 GiB per second or 5 million records per second.
- Scale up the number of shards in  stream using the Console or UpdateShardCount API
- The default shard limit is 500 shards for the following AWS Regions: US East (N. Virginia), US West (Oregon), and EU (Ireland). For all other Regions, the default shard limit is 200 shards.
- The maximum size of the data payload of a record before base64-encoding is up to 1 MB.
- GetRecords can retrieve up to 10 MiB of data per call from a single shard, and up to 10,000 records per call. Each call to GetRecords is counted as one read transaction.
- Each shard can support up to five read transactions per second. Each read transaction can provide up to 10,000 records with an upper limit of 10 MiB per transaction.
- You can register up to twenty consumers per stream to use enhanced fan-out.

### Partition Key
A **partition key** is used to **group data by shard** within a stream.

Kinesis Data Streams segregates the **data records** belonging to a **stream** into **multiple shards**.

It uses the **partition key** that is associated with each data record to determine **which shard** a given **data record** belongs to.

Partition keys are Unicode strings with a **maximum** length limit of **256 bytes**.

### Sequence Number

_TODO_...

### Retention Period
The retention period is set at the stream level. The **retention period** is the length of time that data records are accessible after they are added to the stream.

`Default retention period`: **24 hours** after creation
You can increase the retention perioud up to **168 hours (7 days)**

### Data Records
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

- Video streams stores incoming media data as **chunks**
- Each **chunk** consists of:
    - media metadata
    - **fragment**
    - video stream specific metadata (fragment #, server / producer timestamps)
- Media data in the fragments is packed into a structured format such as **Matroska (MKV)**

- Server-side encryption is always enabled on Kinesis Video Streams
- Server-side encryption using KMS or CMK allows encrypting data at rest in Vid streams

### Kinesis Connector Library

Helps java developers integrate Kinesis Streams with other AWS services.

The library provides connectors to various AWS services including S3.

Each Amazon Kinesis connector application is a pipeline that understands how records from Kinesis Stream will be handled.

KCL uses a unique Amazon DynamoDB table to keep track of the application's state

KLC uses the name of the Kinesis Data Streams app to create the name of the table, each application name _must_ be unique

### Kinesis agent

Kinesis Agent is a stand-alone Java application that can easily collect and send data to Kinesis Data Streams.

The agent can continuously monitor set of fles (more for log fles) and Aggregation of data is not possible.

#### Kinesis Data Analytics

**Stagger Windows**: A query that aggregates data using _keyed_ time-based windows that open as data arrives. The keys allow for multiple overlapping windows. This is the recommended way to aggregate data using time-based windows, because Stagger Windows reduce late or out-of-order data compared to Tumbling windows.

**Tumbling Windows**: A query that aggregates data using _distinct_ time-based windows that open and close at regular intervals.

**Sliding Windows**: A query that aggregates data _continuously_, using a fixed time or rowcount interval.

**Continuous Query**: query over a stream executes continuously over streaming data. Enables scenarios such as ability for applications to query the stream and generate alerts.