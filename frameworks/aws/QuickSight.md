# Amazon QuickSight

- To join tables both are based on the same SQL database data source.
- To join tables from _different_ data sources, create the join **before** importin to Amazon Quicksight (Quicksight does not provide facility to join tables from different data sources)
- QuickSight allows configure join type (inner, outer, left, right)
- If you choose a table and made changes to the fields (ex: changing field name), these changes are discarded when you add tables using the join interface

## Data Sources

#### Relational Data Sources
- Athena
- Aurora
- Redshift & Redshift Spectrum
- S3 & S3 Analytics
- Apache Spark >= 2.0
- MariaDB >= 10.0
- SQL Server >= 2012
- MySQL >= 5.1
- PostgreSQL >= 9.3.1
- Presto >= 0.167
- Snowflake
- Teradata >= 14.0

Amazon **Redshift clusters**, Amazon **Athena databases**, and **Amazon Relational Database Service (RDS)** instances **must be in AWS**. Other database instances must be in one of the following environments to be accessible from Amazon QuickSight:

- Amazon EC2
- On your local network
- In a data center or some other internet-accessible environment

#### File Data Sources
- CSV/TSV - Delimited text files
- ELF/CLF - Extended and common log format files
- JSON - Flat or semi-structured data files
- XLSX - Microsoft Excel files

Files in Amazon S3 that have been compressed with **zip**, or **gzip** (www.gzip.org), can be **imported as-is**.

If you used another compression program for files in Amazon S3, or if the files are on your local network, you need to unzip them before importing them.

#### JSON Data Sources

Amazon QuickSight natively supports JSON flat files and JSON semi structure data files.

You can either upload a JSON file or connect to your Amazon S3 bucket that contains JSON data. Amazon QuickSight automatically performs schema and type inference on JSON files and embedded JSON objects. Then it flattens the JSON, so you can analyze and visualize application-generated data.


### Data Preparation

The process of transforming raw data for use in an analysis. Includes changes like the following:
- Filtering out data so you can focus on what's important to you
- Renaming fields to make them readable
- Changing data types so they are more useful
- Adding calculated fields to enhance analysis

### SPICE

SPICE is Quicksight's "_Super-fast, Parallel, In-memory Calculation Engine_". SPICE is engineered to quickly perform advanced calculations and serve data.

### Data Analysis

The basic workspace for creating and interacting with visuals, creating _Sheets_ and _Stories_ that you can plas as a slideshow.

### Dashboards

A dashboard is a **read-only snapshot** of an analysis that you can share with other Quicksight users for reporting purposes.

Dashboards let users view and filter the dashboard visuals without changing the underlying data.

### Security

Restrict access to a data set using [row level security](https://docs.aws.amazon.com/quicksight/latest/user/restrict-access-to-a-data-set-using-row-level-security.html). This can be done before or after the dataset is shared. Only the people you shared with can see the data. By adding row-level security you can further control their access.

**Steps to Row level security**
- Create data set rules for Rls (grant access, or deny access)
- Apply row-level permissions by using a file or query that contains the data set rules
- Choose your permissions data set, and assign policies to users and groups.
- To apply the data set rules, add the permissions data set to the data set.

#### Managing Users

In Enterprise edition, you can manage users through any of the following.

- Add and remove Microsoft Active Directory groups to create and deactivate user accounts
- Federated Logins
- Inviting Users by Email

### Chart Types

**Tabular Reports** - customized table view of your data.

**Heat Maps** - use to show a measure for the **intersection of two dimensions**, with color-coding to easily differentiate where values fall in the range.
- Heat maps can also be used to show the count of values for the intersection of two dimensions.

**Line Chart** - used to compare changes in measure values over a period of time for following scenarios:
- One measure over a period of time (gross sales by month)
- Multiple measures over a period of time (gross sales and net sales by month)
- One measure for a dimension over a period of time (# of flight delays per day by airline)

**Tree Maps** - used to visualize one or two measures for a dimension

**KPI** - key metrics / goals