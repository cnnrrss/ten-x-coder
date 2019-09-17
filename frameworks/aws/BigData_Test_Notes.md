# Big Data Test Cheatsheet

Categories:
- Analysis                  15/70   21%
- Processing                17/70   24%
- Storage                   18/70   25%
- Visualization             5/70     7%
- Collection                8/70    11%
- Data Security             7/70    10%

## Analysis

1.1) AWS Kinesis Video - Which of the below options that support
collection, processing and playback of speci#c video clips on-demand? 

- Creation, processing are managed by Kinesis Video Streams while extraction of data
from media sources is supported by Producer libraries
- HTTP Live Streaming (HLS) or GetMedia API supports viewing an Amazon Kinesis video
stream, either for live playback or to view archived video
- Kinesis Video Streams supports live playback or to view archived video through Kinesis
Video Streams console

`https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/how-it-works-kinesis-video-api-producer-sdk.html`


1.2) Secured storage, database, search. Provide capabilities with LEAST overhead

- Amazon DynamoDB provides database layer to host data, relationships and stories
- Amazon CloudSearch to provide Search (Discover) Capability for the messaging app (Cheaper than ES)
- Amazon S3 to provide layered storage for both active/inactive/archive layers of data

1.10) SQL against streaming source using Kinesis Analytics. What kind of use cases can be realized?
- Metrics over time windows, thereby stream data to Amazon S3 or Amazon Redshift through Kinesis stream 
- Aggregate and Process streaming data results to feed real-time dashboards
- Enable Custom Metrics and supports real-time monitoring, noti!cations, and alarms

1.16) To reduce time to insight, optimize costs, and increase %exibility for its analysis, which tool can provide interactive querying capability for the datasets available in CSV, JSON, or columnar data formats such as Apache Parquet and Apache ORC OOTB with S3?
- Athena helps you analyze unstructured, semi-structured, and structured data stored in Amazon S3 and allow run ad-hoc queries using ANSI SQL, without the need to aggregate or load the data into Athena.

2.2)  Please advise why do you think DynamoDB is a right platform to address QuickDialog Business needs. select 3 options.
- Support key-value and document based data models with a provision to have a fexible schema, so each row can have any number of columns at any point in time which can cater new business requirements at relative ease
- Provides two read/write capacity modes for each table: on-demand and provisioned
- Replicate your data automatically across your choice of AWS Regions and automatically scale capacity to accommodate your workloads so that globally distributed applications can access data locally in the selected regions to get single-digit millisecond read and write performance

## Processing

1.3) Improve real-time analytics...
- Use multiple shards to integrate data from di"erent applications, reshard by splitting
hot shards to increase capacity of the stream
- Use CloudWatch metrics to monitor and determine the "hot" or "cold" shards and
understand the usage capacity
- Use multiple shards to integrate data from di"erent applications, reshard by merging
cold shards to reduce cost of the stream

1.4) Please identify the minimum artefacts in EMR to also optimize overall TCO.
- 1 master nodes to coordinate the distribution of data and tasks among other nodes for
processing
- 4-6 core nodes in overall with software components that run tasks and store data to
support full workloads

1.7) There are 2 consumers that are using enhanced fan-out to receive data from the stream. Please detail the implementation speci#cs of
consumers. 
- An enhanced fan-out pipe provides up to 2 MiB/sec of data per shard
- consumers can be built out KCL library or API
- KCL automatically subscribes to consumer to all the shards of the stream

1.17) What kind of tasks are supported by AWS Glue?
- Discovers and catalogs metadata about data stores into AWS Glue catalog
- Populates the AWS Glue Data Catalog with table de!nitions from scheduled crawler programs which classi!er logic to infer the schema, format, and data types of the data
- Generates ETL scripts based on python, scala to transform, #atten, and enrich data from source to target
- Triggers ETL jobs based on a schedule or event and scales resources, as needed, to run jobs

2.3) Requirements:

- Use pipelines to manage the conversion of videos into diferent formats
- Transform videos into diferent formats using a single job ( address reusability) and add new formats with simple changes
- Pre-defned templates to convert videos into specifc formats
- Notify administrators apprised of the status of a job

AWS Elastic Transcoder fulflls the requirements.

1.18) Requirements:
Transform syslog data to CSV, Load the data capture into Redshift, capture transform and delivery failures, backup the syslog streaming data to S3 bucket
- Streaming data is delivered to your S3 bucket !rst. Kinesis Data Firehose then issues an Amazon Redshift COPY command to load data from your S3 bucket to your Amazon Redshift cluster
- The transformation failures and delivery failures are loaded into processing-failed and errors folders in same S3 bucket
- When Redshift is selected as destination, and Source record S3 backup is enabled, and Backup S3 Bucket is de!ned, untransformed incoming data can be delivered to a separate S3 bucket 

2.1) IOT data collected needs to be integrated with diferent services
and fltered to diferent services based on SQL Queries against the data collected OOTB. How can we enable this? select 1 option:

- IOT Rules address fltering and integration with the above services mentioned. (Best solution OOTB)


## Storage 

1.5) What is the best approach to consume all the data captured in the stream is shared with all the applications mentioned above which includes S3, Redshift and Elasticsearch (ES).
- Create separate Kinesis Firehose for different downstream applications like S3, ES, and Redshift

1.12) ORDER_FCT is a Fact Table with billions of rows related to orders | PART_DIM is a part dimension table with billions of records that defines the materials that were ordered.

One of the key requirements includes ORDER_FCT and PART_DIM are joined together in most of order
related queries. ORDER_FCT has many other dimensions to support analysis.
- Distribute the ORDER_FCT and PART_DIM on same key with KEY distribution

13) 2 hot shards, what is basic re-shard strategy?
- SHARD 1 need to be split as SHARD 1A, 1B, SHARD2 into SHARD2A 382..410 and SHARD B as 411..454



## Visualization

1.6) The team is planning to use a Chart to visualize a comparison between a key value and its target value:
- KPI's

1.11) Team wanted to build some charts on a single dimension, grouped dimensions against a single measure and multiple measures and their aggregations and summaries based on X and Y dimensions. 
- Bar Charts
- Combo Charts

## Collection

1.8) AWS IOT Integration, capture the existing persistent representation of the devices, based on canonical data format like JSON.
- Device Shadow is a JSON document used to store and retrieve current state information for a device.
- Device Shadow service provides persistent representations of your devices in the AWS Cloud

1.9)  A lot of kinesis throughput exceptions are thrown. Which of the following option can help us to resolve the issue?
- Increase the read and write throughput of Kinesis Application’s DynamoDB table

1.19) Extend search, document management, integration into Data Lake built on EMR with existing stream:
- data collection, pre-processing, and writing of data into data streams using Kinesis Agents and reading of data using enhanced fan-out consumers built using kinesis Applications and writing to downstream applications using connector libraries. New consumer applications need to be added.

## Data Security

1.14) The team needs a mechanism to enable row-level security, there restricting access to data elements in the dataset. How can this be achieved?
- Creating Data Set Rules for Row-Level Security
- Apply row-level permissions by using a !le or query that contains data set rules
- Choose your permissions data set, and assign permission policies to users and groups
- To apply the data set rules, you add the rules as a permissions data set to your data set.

1.15) The Team is looking for storing persistent data complemented with server-side encryption, read-after-write consistency, and list consistency and enables Data Lake for the enterprise to support analytics.
- EMRFS implementation of HDFS used for reading and writing regular files from Amazon EMR directly to Amazon S3