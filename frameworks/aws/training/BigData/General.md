# Big Data Specialty General Knowledge

## Domain 1: Collection

### Kinesis Streams Overview

Data retention is 24 hours by default, can go up to 7 days

### \# of Shards

- Producer
Write 1 MB per second PER SHARD
1000 Records per second PER SHARD
(otherwise "ProvisionedThroughputException)

- Consumer Classic
Read 2 MB per second PER SHARD
5 API calls per second PER SHARD

- Consumer Enhanced Fan-Out:
2 MB/s at read PER SHARD, PER ENHANCED CONSUMER
No API calls needed (push model)


### KPL
- Long-running produccers
- Submit metrics to CloudWatch
- Batching (both turned on by default) - increase throughput, decrease cost
    - Collect Records and write to multi shards in same PutRecords API call
    - Aggregate - increase latency
        - Increase payload size
- Automated and configurable retry mechanism
- Syncronous or Asynchronous API
- Compression must be implemented by the user

### Kinesis Firehose
**Sources**:
    - SDK, KPL, Kinesis Agent, IoT rules actions, Kinesis Data Streams, CloudWatch Logs & Events
        - Direct `PUT`

**Targets**:
    - S3, Elasticsearch, Redshift, Splunk
        - For Redshift, first s3 then `COPY` cmd to Redshift

**Lambda Blueprints**:
    Transform data in Firehose
    Many "blueprints available" (ex: CSV -> JSON)

**Compression**:
    - Disabled
    - GZIP
    - Snappy
    - Zip

**Source Records**, **Transformation failures**, and **Delivery failures** can be recorded in S3

Buffer Size and Buffer Time can be set,  enables new **_near_** real time
    - **Buffer size**: (1 - 128 MB)
    - **Buffer time** (60 - 900 seconds )

|Firehose|Streams|
|--------|-------|
|Fully managed, send to S3, Splunk, Redshift, ElasticSearch|Going to write custom code (producer / cosumer)|
|Near real time (lowest buffer time is 1 minute)|Real time (~200ms latency for classic, ~70ms latency for enhanced fan-out)|
|Automated scaling|Must manage scaling (shard splitting / merging)|
|No data storage|Data Storage for 1 to 7 days, replay capability, multi consumers|
|Serverless data transformations with Lambda|Use w/ Lambda to insert data in real-time to Elasticsearch (for example)|


**Agent** can do CSV to Json transform for you