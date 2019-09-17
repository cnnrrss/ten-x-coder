# Kinesis

You can add data to an Amazon Kinesis data stream via PutRecord and PutRecords operations, Amazon Kinesis Producer Library (KPL), or Amazon Kinesis Agent.

**note how KPL and kinesis agent are different.KPL is only available on Java and C+

**Producers**

Producers are the systems sending data.
You can send data using:
- KPL ( Kinesis Producer Library)
- Kinesis Agent
- Kinesis API

**Consumers** 

Consumer can be implemented with

- Kinesis API
- Kinesis Client Library ( KCL )
    - This is a pre-built library that helps you easily build Amazon Kinesis Applications for reading and processing data from an Amazon Kinesis stream. This library handles complex issues such as adapting to changes in stream volume, load-balancing streaming data, coordinating distributed services, and processing data with fault-tolerance, enabling you to focus on business logic while building applications.
- Kinesis Connector Library
    - Used to connect Kinesis with other services on AWS. KCL is required to use this Library

## Data Firehouse

Lambda Blueprints

Kinesis Data Firehose provides the following Lambda blueprints that you can use to create a Lambda function for data transformation.

• **General Firehose Processing** — Contains the data transformation and status model described in the previous section. Use this blueprint for any custom transformation logic.
• **Syslog to JSON** — Parses and converts Syslog lines to JSONobjects, using predefined JSON field names.
• **Syslog to CSV** — Parses and converts Syslog lines to CSV format.
• **Kinesis Data Firehose Process Record Streams as source** — Accesses the Kinesis Data Streams records in the input and returns them with a processing status.
• **Kinesis Data Firehose CloudWatch Logs Processor** — Parses and extracts individual log events from records sent by CloudWatch Logs subscription filters.

## Data Streams

A single shard can ingest up to 1 MiB of data per second (including partition keys) or 1,000 records per second for writes. Similarly, if you scale your stream to 5,000 shards, the stream can ingest up to 5 GiB per second or 5 million records per second. 

Each shard can support up to fve read transactions per second 

### Kinesis Connector Library
Used to connect Kinesis with other services on AWS. KCL is required to use this Library

## KCL - Kinesis Client Library

The KCL performs the following tasks:

- Connects to the stream
- Enumerates the shards
- Coordinates shard associations with other workers (if any)
- Instantiates a record processor for every shard it manages
- Pulls data records from the stream
- Pushes the records to the corresponding record processor
- Checkpoints processed records
- Balances shard-worker associations when the worker instance count changes
- Balances shard-worker associations when shards are split or merged

KCL uses a unique Amazon DynamoDB table to keep track of the application's state.

KCL uses the name of the Amazon Kinesis Data Streams application to create the name of the table, each application name must be unique.

## KPL - Kinesis Producer Library

https://docs.aws.amazon.com/streams/latest/dev/developing-producers-with-kpl.html

Validate the transaction by checking the successful insert into the stream by embedding automatic and confgurable retry mechanism

Kinesis Streams provide capabilities to use Future objects to validate
UserRecords. No need to complicate the code by storing in memory/transient storage.

Time-to-live records need to be increases if the UserRecords could not
inserted into stream in time