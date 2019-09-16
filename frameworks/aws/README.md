# AWS

## Compute

**AWS EC2**: Raw compute, managed severs. Provides the virtual application servers, known as instances, to deploy your code on. 
Tags: compute, devops

**AWS ECR**: Elastic Container Registry - is a fully managed Docker container registry that makes it easy for developer to store, manage, and deploy Docker  container images.

**AWS ECS**: Elastic Container Service - is a highly scalable, fast, container management service that makes it easy to run, stop, and manage Docker containers on a cluster of Amazon EC2 instances.

**AWS EKS**: Elastic Kubernetes Sercice - is a managed service that makes it easy for you to run Kubernetes on AWS w/o needing to install and operate your own Kubernetes cluster.

**LightSail**: easiet way to get started with AWS for developers who just need virtual private servers. includes everything: virtual machine, SSD storage, data transfer, DNS management, and static IP for low, predictable price.

**Lambda**: Serverless functions. Runs code without provisioning or managing servers. With Lambda, you can run code for virtually any type of application or backend service - all with zero admin.
Tags: compute, devops

**Elastic Container Service**: A highly scalable, high performance container mgmt service that supports Docker containers and allows you to easily run applications on a managed cluster of EC2 instances.
Tags: compute, devops

**AWS Batch**: enables you to run batch computing workloads on the AWS Cloud. Batch computing is a common way for developers, scientists and engineers to access large amounts of compute resources.

**AWS Elastic Beanstalk**: quickly deploy and manage applications in AWS Cloud w/o worring about the infra that runs those applications. Reduces mgmt complexity w/o restricting choice or control. Simply upload app, and e-beanstalk auto handles the details of capacity and provisioning, lb, scaling, and health monitoring

**Amazon VPC**:

**AWS Serverless Application Repository**:

## Storage

**S3**: Storage
**S3 Glacier**:
**EFS**: Elastic File Sercver
**AWS Backup**:
**FSx**:

## Databases
**RDS**: Relational Database Service (MySQL fork)

**DynamoDB**: NoSqlDB

**ElastiCache**: easy to set up, manage, and scale distributed in-mem cache environments in AWS Cloud. Provides a high perf, resizable and cost-effective in-mem cache, while removing complexity associated w/ deploying and managing a distributed cache env. ElastiCache works w/ both Redis and Memcached engines.

**Redshift**: fast, fully managed, petabyte-scale data warehouse service that makes it simple and cost-effective to efficiently analyze all your data using your existing business intelligence tools. Optimized for datasets ranging from a few hundred GB - PB or more and costs < $1,000/TB/year. 1/10 the cost of most traditional data warehousing solutions.

**QLDB**: Quantum Ledger DB - 

**DocumentDB**: (w/ MongoDB compatibility) is a fast, reliable, and fully managed db service that makes it easy for you to set up, operate, and scale MongoDB-compatible dbs.

Data Sync, Migration Hub, DMS, Schema Conversion?

**Neptune**: fast, reliable Graph DB. Easy to build/run apps that work with highly connected datasets. Core of Neptune is purpose-built, high perf graph db engine optimized for storing billions of relationships and querying the graph w/ milliseconds latency. Support popular graphQL languages Apache TinkerPop, Gremlin and W3C's SPARQL.  Neptune powers graph use cases such as recommendation engines, fraud detection, knowledge graphs, drug discovery, and network security.

#### Topics
(Even, All, Any) key distribution
Sharding process

## Streaming
**SQS**: Simple Queueing Service (like Kafka) - fully managed message queuing service, makes it easy to decouple and scale microservices, dist systems, and serverless apps.

**Amazon Kinesis**: makes it easy to collect, process, and analyze video and data streams in real time. real-time processing of streaming data at massive scale.
    - Data Firehose
        - Load data streams into AWS data stores.
    - Video Streams
        - On-demand playback, custom video processing,
        - Capture, process, and store video streams for analytics and machine learning.
    - Data Streams
        - Build custom applications that analyze data streams using popular stream-processing frameworks.
    - Data Analytics
        - Process and analyze streaming data using SQL or Java.

**KCL Library**: Kinesis Consumer Lib
**KPL Library**: Kinesis Producer Lib

## ML
**AWS Rekognition**:
**AWS Polly**:
**AWS Comprehend**:
**AWS SageMaker**:
**AWS Transcribe**:
**Data Pipeline**

#### Topics
SQL, Query Types


## Analytics & Monitoring
**Athena**: Analytics, analyze unstructured, semi-structured and structred data in s3.

**Quicksight**: BI analytics service to build visualizations, perform ad hoc analysis, and quickly get BI from data. Seamlessly discovers AWS data sources, enables orgs to scale to 100_000's of users, and delivers fast/responsive query perf by using in-mem engine (SPICE).

**CloudWatch**: reliable, scalable, flexible, real-time monitoring solution you can start using within minutes. No longer need to setup, manage, and scale your own monitoring infra.
    - CloudWatch Events
        - Send system events from AWS resources to AWS Lambda functions, Amazon SNS topics, streams in Amazon Kinesis, and other target types.
    - CloudWatch Logs
        - Monitor, store, and access your log files from Amazon EC2 instances, AWS CloudTrail, or other sources.

**CloutTrail**: monitor deployments in cloud w/ history of AWS APU calls for your account, including API calls made via AWS Management Console, SDK's, cmd line tools, and higher-level services. Identify which users and accounts, source IPs that called APIs for services that support CloudTrail. 



## Big Data
**Amazon Glue**: serverless fully managed extract, transform, and load (ETL) service

**Lake Formation**:

Apache Hue, Flink, Phoenix, Tez

### Devops

https://aws.amazon.com/getting-started/use-cases/devops/

**AWS CodePipeline**: A continuous delivery service for fast and reliable application updates. CodePipeline builds, test, and deploys your code every time thereis a code change, based onthe release proecess models you define.

**AWS CodeDeploy**: Automates code deployments to any instance, including EC2 and on-premises.

**AWS CodeCommit**: A fully-managed source control service that makes it easy for companies to host secure and highly scalable private Git repos.

**CloudFormation**: An easy way to create and manage a collection of related AWS resources, provisioning and updating them in an orderly and predictable fasion.

**AWS OpsWorks**: A config management service that helps you configure and operate application of all shapes and sizes using Chef.

**AWS Config**: Provides you with an AWS resource inventory, config history, and config change notifications to enable security and governance.

**CloudBuild** - Travis, etc..
KMS
CMK
Client side vs server side

**API Gateway**
**SNS** - Simple Notification Serfice (Email, Text, etc..)
**EMR** - Easily Run and Scale Apache Spark, Hadoop, HBase, Presto, Hive, and other Big Data Frameworks (EMRFS)

#### Topics
Elastic Transcoder
Lambda Blueprints
Enhanced Fan-out/Fan-in

### Security
**Amazon Macie**: Discover, Classify, and Protect Your Data.


### IOT
**Device Registry**
**Device Shadow**


## Certifications

**AWS Data Specialist**


**AWS TCO and Cloud Economics (Digital)**
- [x] [Mod 1: What Is TCO?](training/MOD1_TCO.md#Mod-1:-What-is-TCO?)
- [] [Mod 2: Target Customers](training/MOD1_TCO.md#Mod-2:-Target-Customers)
- [] [Mod 3: Discussing AWS Economics](training/MOD1_TCO.md#Mod-3:-Discussing-AWS-Economics)
- [] [Mod 4: Data Center Economics](training/MOD1_TCO.md#Mod-4:-Data-Center-Economics)
- [] [Mod 5: Delivering the TCO Message](training/MOD1_TCO.md#Mod-5:-Delivering-the-TCO-Message)
- [] [Mod 6: TCO Modeling](training/MOD1_TCO.md#Mod-6:-TCO-Modeling)
- [] [Mod 7: Presenting the results of TCO](training/MOD1_TCO.md#Mod-7:-Presenting-the-results-of-TCO)
- [] [Mod-8: Addressing Common Issues & Objections](training/MOD1_TCO.md#Mod-8:-Addressing-Common-Issues-&-Objections)
- [] [Accreditation](training/MOD1_TCO.md#Accreditation)

