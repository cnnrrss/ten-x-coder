# AWS Data Pipeline

![](../../assets/aws_data_pipeline.png)

When 'DP' runs a pipeline, it compiles the components to create a set of _actionable instances_. Each instance contains all the information for performing a specific task.

To provide robust data management, DP retries a failed operation. it continues to do so until the task reaches the max # of allowed retries. Attempt objects track the various attempts, retries and failures.

AWS Data Pipeline hands the _instances_ out to **task runners** to process.


**Pipeline Components** -> **Instances** -> **Attempts**

### Pipeline Components
Pipeline components represent the business logic of the pipeline and are represented by the different sections of a pipeline definition.

Specify:
- Data sources
- Activities
- Schedule
- Preconditions of the workflow

Can inherift properties from parent components.

### Instances
When AWS Data Pipeline runs a pipeline, it compiles the pipeline components to create a set of **actionable instances**.

### Task Runners

A task runner is an application that polls AWS Data Pipeline for tasks and then performs those tasks.

Two ways to use Task Runner to process pipeline:

- AWS Data Pipeline installs Task Runner for you on resources that are launched and managed by the AWS Data Pipeline web service.
- You install Task Runner on a computational resource that you manage, such as a long-running EC2 instance, or an on-premises server.

### Attempts

To provide robust data management, AWS Data Pipeline retries a failed operation.

It continues to do so until the task reaches the maximum number of allowed retry attempts. 

Attempt objects track the various attempts, results, and failure reasons if applicable.

### Data Nodes

A Data Node defines the location and type of data that a pipeline activity uses as input or output. Supported data nodes:

**DynamoDBDataNode** - A DynamoDB table that contains data for HiveActivity or EmrActivity to use.

**SqlDataNode** - An SQL table and database query that represent data for a pipeline activity to use.

**RedshiftDataNode** - An Amazon Redshift table that contains data for RedshiftCopyActivity to use.

**S3DataNode** - An Amazon S3 location that contains one or more files for a pipeline activity to use.

### Databases

AWS Data Pipeline supports the following db types:

**JdbcDatabase**: A JDBC database.

**RdsDatabase**: An Amazon RDS database.

**RedshiftDatabase**: An Amazon Redshift database.