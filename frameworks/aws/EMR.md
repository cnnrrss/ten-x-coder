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

- **Hive**: data warehouse and analytics package that runs on top of a Hadoop cluster (DDL). Hive scripts use a SQL-like language called Hive QL that abstracts programming models and supports typical data warehouse interactions. Enables you to **avoid the complexities** of writing Tez jobs based on DAGs or MapReduce programs in a lower level language like Java.

- **Spark**: a distributed processing framework and programming model that helps you do machine learning, stream processing, or graph analytics using Amazon EMR clusters. Similar to Apache Hadoop, Spark is an open-source, distributed processing system commonly used for big data workloads. However, Spark has several notable differences from Hadoop MapReduce. Spark has an optimized directed acyclic graph (DAG) execution engine and actively caches data in-memory, which can boost performance, especially for certain algorithms and interactive queries.

- **Pig**: library that runs on top of Hadoop, providing a **scripting language** that you can use to transform large data sets without having to write complex code in a lower level computer language like Java. The library takes **SQL-like commands** written in a language called Pig Latin and converts those commands into **Tez jobs** based on directed acyclic graphs (DAGs) or **MapReduce programs**.

- **Hadoop**: Java software framework that supports massive data processing across a cluster of instances. It can run on a single instance or thousands of instances. Hadoop uses various processing models, such as MapReduce and Tez.

- **Presto**: is a _fast_ **SQL query engine** designed for interactive analytic queries over large datasets from multiple sources.

- **Flink**: streaming dataflow engine that you can use to run real-timestream processing on high-throughput data sources. Supports event time semantics for out-of-order events, exactly-once semantics, backpressure control, and APIs optimized for writing both streaming and batch applications. (integrates with: Kinesis, Kafka, Twitter, Cassandra, etc..)

- **HBase**: open source, non-relational, distributed database. HBase runs on top of HDFS to provide non-relational db capabilities for the Hadoop Ecosystem

- **HCatalog**: Hive metastore, that allows you to access Hive metastore tables within Pig, Spark SQL, and/or custom MapReduce applications. HCatalog has a REST interface and command line client that allows you to create tables or do other operations.

- **Hue**: "Hadoop User Experience" is open-source web-based GUI for use with Amazon EMR and Apache Hadoop. Hue groups together several different Hadoop ecosystem projects into a configurable interface. Amazon EMR has also added customizations specific to Hue in Amazon EMR. Hue acts as a front-end for applications that run on your cluster, allowing you to interact with applications using an interface that may be more familiar or user-friendly. Hive and Pig editors.

- **Jupyter Notebook**

- **Phoenix**: OLTP and operational analytics. Allowing you to use standard SQL queries and JDBC APIs to work with an Apache HBase backing store.

- **TensorFlow**: an open-source symbolic math library for machine intelligence and deep learning applications. 

- **Tez**: a framework for creating a complex directed acyclic graph (DAG) of tasks for processing data. In some cases, it is used as an alternative to Hadoop MapReduce. For example, Pig and Hive workflows can run using Hadoop MapReduce or they can use Tez as an execution engine.

- **Zeppelin**: as a notebook for interactive data exploration. like Jupyter.

- **ZooKeeper**: is a centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services for replicated node orchestration.

- **Sqoop**:  is a tool for transferring data between Amazon S3, Hadoop, HDFS, and RDBMS databases. 

- **Ganglia**: is a scalable, distributed system designed to monitor clusters and grids while minimizing the impact on their performance. When you enable Ganglia on your cluster, you can generate reports and view the performance of the cluster as a whole, as well as inspect the performance of individual node instances.

- **Oozie**: Workflow Scheduler to manage and coordinate Hadoop jobs. 

- **Livy**
- **Mahout**
- **MXNet** 