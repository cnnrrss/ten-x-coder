# Cheat Sheet


## Facts

**VPC Endpoint Gateways**: only DynamoDB & S3 have them. The rest of the services are VPC Endpoint _Interfaces_
**Redshift Spectrum cross account accesss to S3**: Add poolicy to S3 bucket with GET and LIST ops for an IAM role for the Spectrum account

## S3
Cross-account Redshift Spectrum access to an S3 bucket:
    - Add a policy to the S3 bucket allowing S3 `GET` and `LIST` operations for an IAM role for Spectrum on the Redshift account

## EMR
Which technique will help to improve perf and optimize network cost?
- Use HDFS as a temporary data store for processing
    - Use s3distcp to copy data from S3 to HDFS, apply complex transforms on HDFS using spark then copy processed dataset back to S3.

Notifications on consistency views in terms of eventual consistency:
- Enable CloudWatch metrics and Amazon SQS messages in EMRFS for S3 eventual consistency issues

How to track consistency of multi-step sequential ETL processing pipeline?
- Enable Consistent View. Allows EMR clusters to check for list and read-after-write consistency
- EMR uses DynamoDB to store object metadata and track consistency with **EMRFS Storage**

Storing persistent data, SSE, read-after-write consistency, and list consistency and enables DL analytics
- EMRFS implementation of HDFS used for reading and writing regular files from Amazon EMR directly to S3

EMR cluster looking for storage configuration that supports storing temporary data that is continually changing, such as buffers, caches, scratch data and temporary content. How?
- Master and Core nodes running on EC2 that comes with preconfigured block or preattached disk storage called _instance store_
- Master and core nodes running on _local_ file system or local connected disks

## Kinesis
Throughput exceptions, what can help the issue?
- Increase r/w throughput of Kinesis Application's DynamoDB table

### Kinesis Data Streams
Track Data Streams state, how can this be achieved?
- KCL uses unique DynamoDB table
- KCL uses the name of the Amazon Kinesis Data Streams app to create the name of the table, each name is unique.

Best approach to keep traffic between VPC and Kinesis data streams from leaving Amzn?
- Use interface VPC endpoints.

### Kinesis Streams
End to End encryption, how can this be achieved?
- SSE can be enabled trhough KMS which provides the master keys
- SSE can use a CMK for Kinesis that is managed by AWS
- SSE can use a user-specified CMK
- SSE can use a master key imported into the AWS KMS service

Post reshard, what happens?
- Records flowing to parent rerouted to child
- Records in parent shards remain
- Parent is OPEN state before reshard
- Parent in CLOSED state after reshard
- Streams retention period expires, parent shard moves to EXPIRED state

What kind of monitoring can be enabled using [CloudWatch](#CloudWatch)

### Kinesis Video Streams
Logs?
- _CloudWatch_ collects and processes _raw data_ from Kinesis Video Streams into readable, near real-time metrics
- _CloudTrail_ captures all _API calls_
- Enable continuous delivery of _CloudTrail_ events, stored in S3 in log format

### Kinesis Firehose
Logs captured in S3 and integrated with Glue. What formats are supported with record format conversion enabled in Firehose?
- JSON to Parquet
- JSON to Apache ORC <- faster columnar formats that save space and enable faster queries.

Requirements: Transform syslog->csv, load to Redshift, capture failures, backup data to s3:
- Streaming data delivered to S3 first, then `COPY` to Redshift
- Transform failures in _processing-failed_ and _errors_ folders
- When redshift is destination, and source S3 backup enabled, backup bucket defined, untransformed data can be delivered to separate S3 bucket

Log Files -> Apache ORC, How?
- Kinesis Data Streams can transform using _Record Format_ techniques

### KCL
What kind of monitoring is enabled by the KCL library?
- **Per-KCL-Application Metrics**, aggregated across all KCL workers w/in scope of the app
- **Per-Worker Metrics**, aggregated across all record processors consuming data
- **Per-Shard Metrics**, aggregated across a single record processor

### Kinesis Data Analytics
What kind of downstream apps are supported?
- Lambda
- Kinesis Data Streams
- Kinesis Firehose

What kind of _automated_ monitoring can be enabled?
- Watch single metric overtime, threshold alerts in **CloudWatch** Alarms
- Share log files between accounts, monitor **CloudTrail** log files in real time by sending them to CloudWatch Logs using **CloudTrail Log Monitoring**
- Monitor, store, and access your log files from AWS **CloudTrail** or other sources using AWS **CloudWatch Logs**
- Match events and route them to take action using **CloudWatch Events**

## Athena
What kind of permissions are needed?
- Athena actions
- Permissions to S3 locations with underlying data
- Permissions for resources stored in Glue Data Catalog
- Permissions for encrypted metadata in Glue Data Catalog

Encryption for data in S3 and results:
- SSE with **SSE-S3** for encrypted datasets in S3 and results
- SSE with (**SSE-KMS**) for encrypted datasets in S3 and results
- CSE with (**CSE-KMS)** for encrypted datasets in S3 and results

## Redshift
Best practices for improving queries:
- Configure up to eight query queues to run concurrently
- Set up rules to route queries to particular queues
- SQA prioritizzes short-running queries ahead of longer ones
- If SQA enabled, you can reduce the WLM queues that are dedicated to running short queries

Security Features
- Users inbound access to Redshift cluster manager thru security groups
- Access management to Redshift cluster managed through IAM
- Encryption can be enabled for data in transit, loaded, cluster and SSL connections

Best practices to load _time series_ tables:
- Organize data as sequence of time-series tables for a fixed retention period and create `UNION ALL` view to hide the fact they're in separate tables
- In sequence, each table should be identical but contain different time ranges
- Use `DROP TABLE` instead of `DELETE`

Best practices to improve performance and decrease the operating costs of the Redshift cluster:
- Uncompressed storage for a single table is a one-time optimization, requires table rebuild
- Consider moving each actively queried db to separate dedicated cluster, reduces contention.
- Skip compression analysis using `COPY`

Implement Security end to end:
- Redshift Management Console is controlled by AWS account priveleges
- Define a cluster security group

Table design best practices:
- Best SORT key
- Best Distro style
- Automatic Compression
- Define primary key and foreign key constraints between tables

Best practices loading data into tables:
- Use copy command to load multiple files from S3, EMR, DynamoDB
- Manage consistency using a manifest file
- In order to reduce need for VACCUM, load in sort key order
- Load the data in sequential blocks accorting to sort order

How can data from other services be joined/queried alongside Redshift through Redshift Spectrum?
- Spectrum accesses external dbs in Athena and EMR using _external schema and tables_
- Spectrum references external db in an _external data catalog_
- For external schemas, Redshift needs auth to access Athena and S3 using IAM Roles / Policies

What tables/views can help access perf related info for diagnosis?
- STL system tables generated from Redshift log files to provide _history_ of system
- STV virtual system tables contain snapshots of current system data provide _snapshot_
- _System catalogs_ store schema metadata such as info about tables & columns

Upsert data?
- Efficiently update and insert new data by _loading into intermediate_ table first
- Load data into a _staging table_ then join with your target for an `UPDATE` and `INSERT` stmt


## DynamoDB

Facts:
- Maximum **size** of item **400KB**
- Data types supported:
    - Scalar Types: (string, number, binary, boolean, null)
    - Document Types: List, Map
    - Set Types: string set, number set, binary set
- Primary key must be decided at creation
- Attributes can be added over time
- Each partition can only contain
    - Max 3000 RCU
    - Max 1000 WCU
    - Max of 10GB
- \# of Partitions = MAX(Capacity, Size)
    - Capacity = (Total RCU / 3000) + (Total WCU / 1000)
    - Size = Total Size / 10 GB
- WCU and RCU spread evenly between partitions
- BatchWriteItem
    - Up to 25 PutItem / DeleteItem in one call
    - Up to 16 MB of data written
    - Up to 400 KB of data per item
    - Save latency by reducing # of API calls
    - Operations done in parallel for efficiency
    - Batch can fail, retry failed items

Example RCU / WCU Question:
- 80KB record size
- 400 writes / sec
- 1800 (eventually consistent) reads / sec
`1 WCU = 1 KB / s so we need 80KB * 400 / s = 32000 WCU` \
`1 RCU = 2 ec-reads/s of 4 KB so.. 1800 * 80 / 8 = 18000 RCU`

Scaling
- If RCU issues, can use Dax (cache in DynamoDB / hot key problem)
- Exponential backoff when exception is encountered
- Distribute partition keys as well as possible

Define steps to restore backup of a DynamoDB table available in S3 to a development environment:
- Defines a data node using DynamoDB, specified as input to EMRActivity
- Defines S3DataNode
- Precondition to check whether S3 objects with given prefix is present
- Use EMRActivity to process the restoration

Distributing write activity efficiently during data upload:
- distribute upload work by using _sort_ key to load one item from partition key value, then another, so on..

Hot partitions, to accomodate _uneven_ access patterns, how will DDB use capcity to prevent throttling?
- _Adaptive_ Capacity

Encrypted data at rest, how you gonna do it?
- Encrypts data in table using 256-bit AES
- Encryption at rest integrates with AWS Key Management Service (KMS) for managing encryption key to encrypt the table
- AWS Owned Customer Master Key (CMK) – Default encryption type used to encrypt the table, owned by DynamoDB (no additional charge).
- AWS Managed CMK – The key is stored in your account and is managed by AWS KMS (harges apply).
- You can audit the encryption and decryption of your DynamoDB table by examining the DynamoDB API calls to AWS KMS using AWS CloudTrail.

Two Types of Capacity
- **Burst Capacity**: when you are not using a partition's throughput, DynamoDB reserves a portion of that unused capacity for later bursts of throughput to handle usage spike
- **Adaptive Capacity**: enables your application to continue reading and writing to _hot_ partitions without being throttled

## Query Types
**Stagger Windows** - _keyed_ time-based windows (ex: capture aggregate every _15 mins_) (can be overlapping)
**Tumbling Windows** - _disctinct_ time-based windows (non overlapping)
**Sliding windows** - _continuously_ using a _fixed time_ or _rowcount_ interval
**Continuous** - query _over_ stream executes _continuously_, enables scenarios such as ability to continuously query and generate alerts (Ex: alert if throughput latency is above defined max threshold)

## ML
Ml creating data schemas, what are the sources
- Amazon ML datasources can be created only for RDS, Redshift, S3 services
- Metadata of files in S3, tables, views and collections in dbs are the data sources
- AttributeType includes Binary, Categorical, Numeric and Text datatypes

**Question**: How much revenue can be generated for each team?
- Linear regression algo, regression model
- RMSE metric to provide accuracy to the model

**Question**: Recommend new products based on customer interests?
- Multi-nominal logistic regresssion algo
- Cross-validation technique to train
- Macro-average F1 score to provide accuracy

**Question**: Reviews by customers or bots?
- Binary classification
- Area Under the Curve (AUC) provide accuracy
- Cross Validation

Amazon ML is limited to 100GB of data


[Types of ML Models](https://docs.aws.amazon.com/machine-learning/latest/dg/types-of-ml-models.html)

|Algorithm|Classification|Accuracy|Usage|
|---------|--------------|--------|-----|
|Logistic Regression|Binary|Area under the curve (AUC)|"Is this email spam or not spam?" <br> "Will the customer buy this product?" <br> "Is this product a book or a farm animal?" |
|Multi-nominal logistic regression|Multi-class|Macro-average F1 score|"Is this product a book, movie, or clothing?" <br> "Is this movie a romantic comedy, documentary, or thriller?" <br> "Which category of products is most interesting to this customer?"|
|Linear Regression|Regression Model|Standard root mean square error (RMSE)|"What will the temperature be in Seattle tomorrow?" <br> "For this product, how many units will sell?" <br> "What price will this house sell for?"|

## Quicksight
How do you refresh datasets?
- Refresh a SPICE DS from _Data Sets Page_
- Refresh a SPICE DS during _Data Prep_
- Refresh a SPICE DS on a _Schedule_
- Refresh file based data, delete and recreate the DS
- Refresh data from db, reopen DS
- Refresh data from db, reopen visualization

Which data repos CANNOT be directly connected to Quicksight?
- Neptune
- DynamoDB
- Elasticsearch

[Row level security](https://docs.aws.amazon.com/quicksight/latest/user/restrict-access-to-a-data-set-using-row-level-security.html#create-row-level-security) of elements of dataset, how can this be achieved?
- Data Set Rules for Row-level security
- Apply row-level permissions by using a file or query that contains the data set rules
- Choose your permissions data set, and assign permission policies to users and groups
- To apply the data set rules, you add the rules as a permissions data set to your data set

Submited data set for deletion, what happens next?
- Receive warning that you have analyses that use this data
- Data set deletion does _not_ delete dependent analyses
- Next time you open analyses, you need to select new data set

TODO Test 2 -  55, 57, 60, 61, 62

## CloudWatch
Monitoring of Kinesis streams (KDS):
- Metrics configured for streams auto coll/pushed to CW every _minute_
- Metrics archived for _two weeks_
- Basic stream level data captured every _minute_ , no extra charge
- Shard-level data sent every _minute_ for _additional cost_

## Glue
Key functionality/components of Glue:
- Autogen Pyspark/Scala to perform ETL thru Glue ETL Operations
- Crawlers scan data
- Glue Data Catalog, persistent metadata store

Heuristic to determine root for table of a dir structure:
- Similar schema for tbl1&2, and same src `s3://bckt1/dir1/` -> crawler may create 1 table with 2 partitions
- Similar schema for tbl1&2 and src `s3://bckt1/tbl1|tbl2/` -> crawler may create 2 tables

Tasks supported by Glue?
- Glue catalog (metadata store)
- Crawler: classifier logic to infer the schema
- Generate ETL scripts to transform/flatten/enrich
- Trigger ETL jobs on a schedule

## Storage Gateway
Want cloud-backed storage volumes and retain copy of frequently accessed data subsets locally. How to achieve?
- Volume Gateway
- Cached Volumes

## Apache
SSH Tunnel to service that supports interactive data exploration to EMR master node:
- Presto

## Data Pipeline
What are the key items of scheduled Pipeline?
- _Pipeline components_ represent the business logic
- When DP runs a pipeline, it compiles the _pipeline components_ to create a set of _actionable instances_
- DP retries a failed operation. It continues to do so until the task reaches the max # of allowed attempts
- DP hands the _instances_ out to _task runners_ to process

## Elasticsearch
When encryption is enabled, what all aspects of domain are encrypted?
- Automated snapshots
- ES logs
- Swap files

## Elastic Transcoder
What are the capabilities?
- Provides mgmt console to create _Jobs, Pipelines, and Presets_
- Enables a _RESTful_ service uses HTTPS and JSON
- Allows app code make requests directly to Elastic transcoder
- _Must write necessary code_ to sign and auth your

## IOT
Integrate IOT data collected into the application needs to be based on canonical data like JSON. How can this be achieved?
- Device Shadow is a JSON doc used to store and retrieve state info for a device
- Device Shadow provides _persistent_ representations of your devices

Appropriate authentication for IOT devices:
- IoT devices use X.509 certs for auth
- Mobile Apps use Amazon Cognito Identities for auth
- Web/desktop use IAM or federated entities for auth

What artefacts and their functions need to be used to complete IOT setup w/ data stored in DynamoDB?
- Gateway to secure and communicate w/ AWS IoT
- Rules engine to provide integration services with DynamoDB
- Message broker to provide secure mechanism for IoT to pub and sub msg from each other
- Shadow service to provide persistent reps of devices in Cloud

## General Questions
LEAST overhead that fulfills requirements: DB with archiving/expiring, Search, Storage
- DB = DynamoDB (fully managed, archiving, copying and expiry)
- Search = CloudSearch
- Storage = S3

## VPC Endpoints
DynamoDB & S3 are the only technologies that have VPC Endpoint Gateways, the rest are VPC Endpoint Interfaces.