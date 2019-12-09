## KCL - Kinesis Client Library

**Note**: You wouldn't use KCL with Lambda. It would be implemented within a daemon or long running consumer application.

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

Can be monitored with CloudWatch.

#### Enhanced Fan-Out

This feature enables consumers to receive records from a stream with throughput of up to **2 MiB of data per second** per shard.

Throughput is **dedicated**, which means that consumers that use enhanced fan-out **don't have to contend with other consumers** that are receiving data from the stream.
