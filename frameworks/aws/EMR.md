# EMR - Elastic Map Reduce

EMR integrates with Kinesis for processing real-time streams.

### File system Options

- Enable Ephermal storage using HDFS by distributing the data it stores across instances in the cluster, storing multiple copies of data on different instances to ensure that no data is lost if an instance fails. Ephermal storage can **only** be enabled through HDFS.

- EMRFS to directly access the data stored in S3. EMRFS extends Hadoop to add the ability to directly access data stored in S3 as if it were a file system.

- Local file system storage when each node is created from an Amazon EC2 instance, comes per-configured block of pre-attached disk storage called an instance store.

### Encryption

EMR supports 2 types of encryption:
- LUKS encryption
- SSE-KMS

### Apache

https://docs.aws.amazon.com/emr/latest/ReleaseGuide/emr-kinesis.html

EMR service provides various apps under Hadoop ecosystem for big data processing.

- Hive: data warehouse and analytics package that runs on top of a Hadoop cluster (DDL). Hive scripts use a SQL-like language called Hive QL that abstracts programming models and supports typical data warehouse interactions. Enables you to avoid the complexities of writing Tez jobs based on DAGs or MapReduce programs in a lower level language like Java.

- Spark:

- Pig:

- Hadoop:

- Presto: SQL Query engine, (DML)

- Flink: streaming dataflow engine that you can use to run real-timestream processing on high-throughput data sources.

- HBase: open source, non-relational, distributed database. HBase runs on top of HDFS to provide non-relational db capabilities for the Hadoop Ecosystem

- HCatalog: Hive metastore

- Hue: "Hadoop User Experience" is open-source web-based GUI for use with Amazon EMR and Apache Hadoop.

- Jupyter Notebook

- Phoenix: OLTP and operational analytics. Allowing you to use standard SQL queries and JDBC APIs to work with an Apache HBase backing store.

- TensorFlow:

- Tez: DAG for data processing

- Zeppelin: notebook product, like Jupyter

- ZooKeeper: replicated node orchestration

Others:
- Sqoop
- Ganglia
- Livy
- Mahout
- MXNet
- Oozie