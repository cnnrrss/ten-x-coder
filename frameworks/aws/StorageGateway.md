# AWS Storage Gateway

[What is?](https://docs.aws.amazon.com/storagegateway/latest/userguide/WhatIsStorageGateway.html)

- **Volume Gateway** - Using volume gateways, you can create storage volumes in the AWS Cloud. _Internet Small Computer System Interface (iSCSI)_
    - **Cached volumes**: 
        - You **store** your data in Amazon Simple Storage Service (**Amazon S3**) and retain a copy of frequently accessed data **subsets locally**. 
        - Cached volumes offer a substantial cost savings on primary storage and minimize the need to scale your storage on-premises. 
        - You also retain **low-latency** access to your frequently accessed data.

    - **Stored volumes**: 
        - If you need low-latency access to your entire dataset, first configure your on-premises gateway to store all your data locally. 
        - Then asynchronously back up **point-in-time snapshots** of this data to Amazon S3. 
        - This configuration provides durable and **inexpensive offsite backups** that you can recover to your local data center or Amazon EC2. (For example, if you need replacement capacity for disaster recovery, you can recover the backups to Amazon EC2.)
- **File Gateway** (cost effective, traditional) - You can use a file gateway to ingest files to Amazon S3 for use by object-based workloads and for cost-effective storage for traditional backup applications.
    - Network File System (**NFS**)
    - Server Message Block (**SMB**)
    - Integrate with AWS services:
        - Common access management using AWS Identity and Access Management (IAM)
        - Encryption using AWS Key Management Service (AWS KMS)
        - Monitoring using Amazon CloudWatch (CloudWatch)
        - Audit using AWS CloudTrail (CloudTrail)
        - Operations using the AWS Management Console and AWS Command Line Interface (AWS CLI)
        - Billing and cost management
- **Tape Gateway** (cost effective & durable, long-term) -  If you are looking for a cost-effective, durable, long-term, offsite alternative for data archiving, deploy a tape gateway.
    - cost effective backup data in **GLACIER** or **DEEP_ARCHIVE**