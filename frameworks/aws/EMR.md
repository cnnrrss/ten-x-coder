# EMR 

`EMR = Elastic Map Reduce`

EMR integrates with Kinesis for processing real-time streams.

### File system Options

- Enable Ephermal storage using HDFS by distributing the data it stores across instances in the cluster, storing multiple copies of data on different instances to ensure that no data is lost if an instance fails. Ephermal storage can **only** be enabled through HDFS.

- EMRFS to directly access the data stored in S3. EMRFS extends Hadoop to add the ability to directly access data stored in S3 as if it were a file system.

- Local file system storage when each node is created from an Amazon EC2 instance, comes per-configured block of pre-attached disk storage called an instance store.

### Encryption

EMR supports 2 types of encryption:
- LUKS encryption
- SSE-KMS

### Apache Open Source
[Apache](./Apache.md#Apache)