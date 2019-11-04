# Glue

Fully managed ETL service that makes it simple and cost-effective to categorize, clean, enrich, and move data reliably between data stores.

#### Use Cases:
- Queries against an Amazon S3 Data Lake
- Analyze Log Data in Your Data Warehouse
- Unified View of Your Data Across Multiple Data Stores
- Event-drive ETL Pipelines

### Lambda

Trigger one or more Glue ETL jobs using Lambda

### Amazon RDS, Redshift, S3, DBs running on EC2

Discover your store associated metadata in the Glue Data Catalog
- Enrich the data by joining with a relational database
- Load this data as tables to your data warehouse

### ETL

Schedule ETL jobs to transform and load the data

### CloudWatch

Logs and notification are pushed to CloudWatch

### Glue Crawlers

Glue lets you set up **crawlers** that can scan data in all kinds of repositories, classify it, extract schema info from it, and store the metadata automatically in the AWS Glue Data Catalog.

AWS Glue acts as the ETL Engine and **Glue Data Catalog** acts as the central metadata repository

### Glue Data Catalog

A persistent metadata store. Managed servce that lets you store, annotate, and share metadata in the AWS Cloud in the same way you would in an Apache _Hive_ metastore.

Provides a uniform repo where disparate systems can store and find metadata to keep track of data in silos.

### Athena, Redshift, EMR

Once catalogued in Glue, your data is immediately available for analysis

Access the Glue Data Catalog for ETL and analytics

### QuickSight & Other BI Tools

Run reports