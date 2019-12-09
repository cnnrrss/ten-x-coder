# Athena

Athena is serverless, no infra, setup or maintenance. You pay for the queries you run (Like BQ).

Athena scales automagically-executing queries in parallel-so results are fast, even on large datasets and complex queries.

You can run **ad-hoc queries** using ANSI SQL, w/o need to aggregate or load data into Athena.

Use Athena to generate reports, explore data w/ BI tools, or SQL clients connected with a **JDBC** or **ODBC** driver.

Presto under the hood

### Usage Examples
- Ad-hoc queries of web logs
- Querying staging data before loading to Redshift
- Analyze CloudTrail / CloudFront / VPC / ELB etc logs in S3
- Integration with Jupyter, Zeppelin, RStudio notebooks
- Integration with QuickSight
- Integration via ODBC / JDBC with other visualization tools

### Data Formats

Athena helps you analyze unstructured, semi-structured and structured data stored in S3.

Examples:
- CSV
- JSON
- _Columnar_: Apache Parquet, Apache ORC
    - Save LOTS of money by using columnar formats
    - Save 30-90% and get better performance

### Integrations

- Athena **integrates** with **Amazon QuickSight** for easy data visualization.
- Athena **does not** integrate with **Amazon RDS**.
- Athena **integrates** with the **AWS Glue Data Catalog**, which ofers a **persistent metadata store** for your data in Amazon S3.
    - Allows you to create tables and query data in Athena
    - Central metadata store available throughout your AWS account
    - Integrated with the ETL and data discovery features of AWS Glue.

### Security
- Access Control
    - IAM, ACLs, S3 bucket policies
    - AmazonAthenaFullAccess / AWSQuicksightAthenaAcce

- Encrypt results at rest in S3 staging directory
    - Server-side encryption with S3-managed key **(SSE-S3)**
    - Server-side encryption with KMS key **(SSE-KMS)**
    - Client-side encryption with KMS key **(CSE-KMS)**

- **Cross-account access** in S3 bucket policy possible

- Transport Layer Security (**TLS**) encrypts **in-transit** (between Athena and S3)

### Permissions

To run queries in Athena, you must have the appropriate permissions for:
- **Athena actions**
- The Amazon **S3** locations where the **underlying data** is stored that you are going to query in AthenaData
- The resources that you store in **AWS Glue  Catalog**, such as **databases** and **tables**, tat you are going to query in Athena
- The **encrypted metadata** in the **AWS Glue Data Catalog** (if you migrated to using that metadata in Athena and the metadata is encrypted)