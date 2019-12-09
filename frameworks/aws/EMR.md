# EMR

`EMR = Elastic Map Reduce`

EMR integrates with Kinesis for processing real-time streams.

Managed Hadoop framework on EC2 instances

Incldues Spark, HBase, Presto, Flink, Hive & more

EMR Notebooks

Several integration points with AWS

## EMR Cluster
- **Master node**: manages the cluster
    - Single EC2 instance
- **Core node**: Hosts HDFS data and runs tasks
    - Can be scaled up & down, but with some risk
- **Task node**: Runs tasks, does not host data
    - No risk of data loss when removing
    - Good use of _spot instances_

### EMR Usage
- Transient vs Long-running Clusters

### Integration
- **EC2** for the instances that comprise the cluster
- **VPC** to configure virtual network in which you launch instances
- **S3** to store input and output data
- **Cloudwatch** to monitor cluster performance and configure alarms
- **IAM** to configure permissions
- **CloudTrail** to audit _requests_ made to the service
- **Data Pipeline** to schedule and start your clusters

### File system Options

- **HDFS**: enables **ephermal storage**. Distributing the data it stores across instances in the cluster, storing multiple copies of data on different instances to ensure that no data is lost if an instance fails. **Ephermal storage** can **only** be enabled through **HDFS**.

- **EMRFS**: access S3 as if it were HDFS
    - EMRFS Consistent View – Optional for S3 consistency

- **Local file system**: storage when each node is created from an Amazon EC2 instance, comes per-configured block of pre-attached disk storage called an instance store.

- **EBS** for HDFS

### Encryption

EMR supports 2 types of encryption:
- LUKS encryption
- SSE-KMS

- AWS KMS Customer Master Keys (CMKs) for EMRFS Encryption
- S3 client-side encryption with custom materials provider


**S3**: server-side or client-side encryption
- SSE-S3, SSE-KMS, CSE-KMS, or CSE-custom

**Note**: S3 server-side encryption with KMS (SSE-KMS) is not available when using EMR version <= 4.4

### Apache Open Source
[Apache](./Apache.md#Apache)