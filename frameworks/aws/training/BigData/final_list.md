# Final List

Last minute review

### Cold

- [x] storage volumes (File Gateway: NFS, traditional, Volume: cached - store in S3 and keep some local; volume - backups, Tape Gateway: Glacier)
- [x] Elastic transcoder
- [x] KPL future objects
    - provide capabilities to use Future objects to validate UserRecords.
- [x] Athean Permissions
    - actions
    - encrypted metadata in GC
    - source data in S3
    - dbs / tbls in GC
- [x] Aurora
    - Aurora fully managed RDS (5x on MySQL, 3x on Postgres)
- [] Kinesis Video Streams (HLS, GetMedia API)
- [x] Kinesis Analytics integrations
    - **sources**: (DS, Firehose, data in S3)
    - **dest**: (DS, Firehose, Lambda)
    - Also, lambda record pre-processing
- [x] IoT Rules / Message Broker / Device Shadow / etc...
    - Device Gateway: manages connections. TLS for in transit.
    - Registry: tracks metadatas
    - Broker: pub/sub messages to and from IoT
    - Rules Engine: gather, process, analyze messages and act on data devices. transform deliver to services. No infra. 1-N actions in parrallel.
    - Device Shadow: persisted, shadow of device to manage device state.
- [x] Datapipeline: • Destinations include **S3, RDS, DynamoDB, Redshift and EMR**. Manages task **dependencies**. **Retries** and notifies on failures. **Cross-region** pipelines. **Precondition** checks. Data sources may be **on-prem**ises. Highly available.
    - **Activities**:
    • EMR
    • Hive
    • Copy
    • SQL
    • Scripts

### Warm
- [x] EMR high network costs / slow performance == HDFS temp storage for processing ?
- [x] KCL uses DynamoDB to track application state, KCL uses Data streams app to create the table and each app name must be unique
- [] CloudWatch vs CloudTrail (monitoring, logging, events, alerts)
- [] Lambda Blueprints / Firehose record format conversions
- [] Firehose failures
- [x] Redshift spectrum features: (query directly from S3, supports structured/semi-structured files [avro, parquet, txt, seq, rcfile, etc...])
- [] Use tags in KDS to track costs from different groups of users / departments
- [] Kinesis Video streams Persistent / Nonpersistent data
- [x] Firehose **JSON** to Parquet / Avro
- [] DynamoDB - distribute upload work by using _sort key_ to load one item from each partition key val, then another from each partition key, and so on.
- [] Refresh SPICE - from sets page, during prep, on schedule, file-based (delete and recreate), reopen dataset or viz
- [] DynamoDB Time Series - Create tables weekly, lower R/W cap to lower value after a week's time
- [x] Cloudsearch - XML / JSON files
- [x] Binary Classification AUC score (.51 is like taking a guess, 1 is good)
- [x] Kinesis Data Streams uses unique DynamoDB table to keep track of state. KCL ProvisionedThroughput errors means u need to increase DynamoDB throughput
- [x] DynamoDB to Redshift (ensure empty values handled, ensure data types matches b/w engines)
- [] ML (Source from Redshift or S3)
- [x] bzip2 supports splitting files

### Must Know
- [x] RCU / WCU DynamoDB (1sc/2ec) 4 KB RCU; 1 KB WCU;
- [] ML Algorithms / Amazon ML data sources
- [] Apache services (Hive, HBase, HCatalog, Phoenix, Hue)
- [] Query Types (Stagger, Tumbling, Sliding, Continuous)
- [x] Redshift Key Distro (ALL, EVEN, KEY, AUTO/None)
- [x] Kinesis Limits (1-128,1-100,5MB 60-900sec,10 GB)
- [] EMR Filesystems (EMRFS, HDFS, Local FS)
- [] Kinesis Analytics supported downstream integrations (KDS, Firehose, Lambda)
- [] Redshift loading data best practices (Unload, Load, Copy, Select, Insert, etc)
- [] Redshift WLM Best Practices
- [] Redshift end-to-end security
- [] Redshift querying best practices
- [] Kinesis stream shardingi
- [] DynamoDB writing best practices
- [] DynamoDB (Global secondary, sparse indexes, bust/adaptive, replication, etc)
- [x] Merge cold shards to make use of unused capacity. Split hot shards to free up resources for keys that target those shards
- [] IoT Policy Actions
