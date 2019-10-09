# Amazon Elastic Search


### Encryption

Amazon ES domains offer encryption at rest. The feature uses **AWS KMS** to store and manage your encryption keys. 

Enables encryption of the following aspects:
- Indices
- Automated snapshots
- Elasticsearch logs
- Swap files
- All other data in the app directory.

**NOT** encrypted:
- Manual snapshots
- Slow logs and error logs