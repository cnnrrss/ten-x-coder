# Final List

Last minute review

### Cold

- [] storage volumes
- [] Elastic transcoder
- [] KPL future objects
- [] Athean Permissions
- [] Aurora / Quicksight
- [] Kinesis Video Streams (HLS, GetMedia API)

### Warm
- [x] EMR high network costs / slow performance == HDFS temp storage for processing ?
- [x] KCL uses DynamoDB to track application state, KCL uses Data streams app to create the table and each app name must be unique
- [] CloudWatch vs CloudTrail (monitoring, logging, events, alerts)
- [] Lambda Blueprints / Firehose record format conversions
- [] Firehose failures
- [x] Redshift spectrum features: (query directly from S3, supports structured/semi-structured files [avro, parquet, txt, seq, rcfile, etc...])
- [] Use tags in KDS to track costs from different groups of users / departments
- [] Kinesis Video streams Persistent / Nonpersistent data

### Must Know
- [] RCU / WCU DynamoDB
- [] ML Algorithms / Amazon ML data sources
- [] Apache services (Hive, HBase, HCatalog, Phoenix, Hue)
- [] Query Types
- [] Redshift Key Distro
- [] Kinesis Limits
- [] EMR Filesystems (EMRFS, HDFS, Local FS)
- [] Kinesis Analytics supported downstream integrations (KDS, Firehose, Lambda)
- [] Redshift loading data best practices
- [] Redshift WLM Best Practices
- [] Redshift end-to-end security
- [] Redshift querying best practices
- [] Kinesis stream shardingi
- [] DynamoDB writing best practices
- [] DynamoDB (Global secondary, sparse indexes, bust/adaptive, replication, etc)
- [x] Merge cold shards to make use of unused capacity. Split hot shards to free up resources for keys that target those shards
- [] IoT Policy Actions
