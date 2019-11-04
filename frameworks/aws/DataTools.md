# AWS Data Tools

AWS has a few database conversion/migration tools.

## AWS Database Migration Service (AWS DMS)

Cloud service makes it easy to migrate relational dbs, warehouses, NoSQL dbs and other types of db stores. Migrate between cloud and on-prem

## AWS Data Sync

Data **transfer** service, simplifies, automates and accelerates moving and replicating data between on prem and AWS storage services over internet or AWS **Direct Connect**

Supports:
- NFS (Network File Service)
- EFS (Elastic File System)
- S3


## AWS Migration Hub

Single place to discover your existing servers, plan migrations and track status of each app migration.

## AWS Schema Conversion Tool (SCT)

Convert your existing db schema from one db engine to another.

You can convert OLTP schema, or data warehouse schema.

You converted schema is suitable for RDS, MySQL DB, Aurora DB cluster, RDS PostgreSQL DB instance or a Redshift cluster.

The converted schema can also be used w/ a db on an EC2 instance or stored as data on an S3 bucket.
