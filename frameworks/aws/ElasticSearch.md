# Amazon Elastic Search

Managed service that makes it easy to deploy, operate, and scale Elasticsearch clusters in the AWS Cloud. 

ES is a popular open-source search and analytics engine for: **log analytics**, **real-time app monitoring**, and **clickstream analysis**.

### Encryption

Amazon ES domains offer encryption at rest. The feature uses **AWS KMS** to store and manage your encryption keys.

Enables **encryption** of the following aspects:
- Indices
- Automated snapshots
- Elasticsearch logs
- Swap files
- All other data in the app directory.

**NOT** encrypted:
- Manual snapshots
- Slow logs and error logs