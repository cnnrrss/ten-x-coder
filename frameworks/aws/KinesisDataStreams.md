# Kinesis Data Streams Overview

- Data retention is **24 hours** by default, can go up to **7 days**
- Ability to reprocess / replay data
- Multiple applications can consume the same stream (like kafka)
- Real-time processing with scale of throughput
- One stream is made of many different shards
- Billing is per shard provisioned, _can have as many shards as you want_

### Kinesis Stream Records

- **Data Blob**: data being sent, serialized as bytes. Up to **1 MB**. Can represent anything
- **Record Key**: sent alongside records, group records in shard. Use highly distributed key to avoid a "hot partition"
- **Sequence Number**: unique identifier for each record put in shards, added by kinesis after ingestion.

### Data Stream Limits
- **Producer**:
    - 1MB/s or 1000 messages/s at write PER SHARD
    - "ProvisionedThroughputException" otherwise
- **ConsumerClassic**:
    - 2MB/s at read PER SHARD across _ALL_ consumers
    - 5 API calls per second PER SHARD across _ALL_ consumers
- **Consumer Enhanced Fan-Out**:
    - 2MB/s at read PER SHARD, PER ENHANCED CONSUMER
        - 2MB/s is _dedicated_ to the consumer, does not compete with other consumers.
    - No API calls needed (push model)
- **Data Retention**:
    - 24 hours by default
    - Can be exteneded to 7 days

### Kinesis Producers
- Kinesis SDK
- Kinesis Producer Library (KPL)
- Kinesis Agent
- 3rd party libraries (Spark, Log4J, Flume, Kafka Connect, NiFi, Appenders)

Other AWS Managed Sources:
- CloudWatch Logs
- AWS IoT
- Kinesis Data Analytics

#### Kinesis Producer SDK - PutRecord(s)
- APIs that are used are PutRecord (one) and PutRecords (many records)
- PutRecords uses batching and increases throughput => less HTTP requests
- `ProvisionedThroughputExceeded` if we go over the limits
- Use case: low throughput, higher latency, simple API, AWS Lambda

#### AWS Kinesis API – Exceptions
- `ProvisionedThroughputExceeded` Exceptions
- Happens when sending more data (exceeding MB/s or TPS for any shard)
- Make sure you don’t have a hot shard (such as your partition key is bad and too much data goes to that partition)

**Solution**:
- Retries with backoff
- Increase shards (scaling)
- Ensure your partition key is a good one

### KPL Library
- Easy to use C++ / Java library
- Building high perf, long-running processes
- Automated and configurable _retry_ mechanism
- Sync / Async API (better perf for async, duh)
- Submits metrics to CloudWatch for monitoring
- **Batching**: (both turned on by default) - increase throughput, decrease cost:
    - **Collect**: collect records and write to multiple shards in the same `PutRecords` API call
    - **Aggregate**: increased latency
        - Capability to _store multiple records in one record_ (go over 1000record/s limit)
- Compression _must be implemented by the user_ (custom code in Producer)

### Kinesis Agent
- Monitor Log files and sends them to Kinesis Data Streams (or Firehose)
- Java-based agent built on top of KPL
- Install in Linux based servers

**features**:
- write to multiple directories and multi streams
- routing feature based on dir/log file
- _pre-process data_ (single line , csv to json, log to json)
- Agent handles file rotation, checkpointing, _retry_ upon failures
- Emits metrics to CloudWatch for monitoring


### Kinesis Security
- Control access / authorization using IAM policies
- Encryption in flight using HTTPS endpoints
- Encryption at rest using KMS
- Client side encryption must be manually implemented (harder)
- VPC Endpoints available for Kinesis to access within VPC

### AWS Kinesis Data Firehose
- Fully Managed Service, no administration
- **Near Real Time** (60 seconds latency minimum for non full batches)
- Load data into **Redshift / Amazon S3 / ElasticSearch / Splunk**
- Automatic scaling
- Supports many data formats
- Data Conversions from JSON to Parquet / ORC (only for S3)
- Data Transformation through AWS Lambda (ex: CSV => JSON)
- Supports compression when target is Amazon S3 (**GZIP, ZIP, and SNAPPY**)
- Only GZIP -  is the data is further loaded into Redshift
- Pay for the amount of data going through Firehose
- _Spark / KCL do not read from KDF_

#### Firehose Buffer Sizing

- Buffer Size (ex: 32MB): if that buffer size is reached, it’s flushed
- Buffer Time (ex: 2 minutes): if that time is reached, it’s flushed
- Firehose can automatically increase the buffer size to increase throughput
- High throughput => Buffer Size will be hit
- Low throughput => Buffer Time will be hit