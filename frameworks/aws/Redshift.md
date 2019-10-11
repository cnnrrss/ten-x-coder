# Amazon Redshift

- Redshift is a columnar SQL Database
- Redshift is a fork of Postgres Database
- Redshift consists of one Leader Node and several Compute Nodes. Each compute node is made up of slices. Each slice is a computing unit.

Redshift delivers fast query performance on datasets ranging in
size from gigabytes to exabytes. Redshift uses columnar storage, data compression, and zone maps to reduce the amount of I/O needed to perform queries. It uses a massively parallel processing (MPP) data warehouse architecture to parallelize and distribute SQL operations to take advantage of all available resources

Amazon Redshift uses machine learning to deliver high throughout based on your workloads. Redshift utilizes sophisticated algorithms to predict incoming query run times, and assigns them to the optimal queue for the fastest processing.

### System Tables and Views

Amazon Redshift has many **system tables** and **views** that contain information about how the system is functioning. 

You can query these system tables and views the same way that you would query any other database tables. 

There are two types of system tables: STL and STV Tables

System tables and views **do not** usethe same consistency model as regular tables. 

#### Ststem Tables 
**STL** - system tables are generated from Amazon Redshift log files to provide a history of the system.

**STV** - STV tables are actually virtual system tables that contain snapshots of the current system data.
- They are based on transient in-mem data and are not persisted to disk

#### System Views 

**SVV** - system viess that contain any reference to transient STV tables

**SVL** - views containing only references to STL tables.

`https://docs.aws.amazon.com/redshift/latest/dg/c_intro_STL_tables.html`

### Data Sources

You can load data from text files in an Amazon S3 bucket, in an Amazon EMR cluster, or on a remote host that your cluster can access using an SSH connection. You can also load data directly from a DynamoDB table.

The maximum size of a single input row from any source is 4 MB.

To export data from a table to a set of files in an Amazon S3, use the UNLOAD command.


### Recommendations

- Addressing uncompressed storage for a singletable is a one-time optimization that _requires_ the table to be rebuilt.
- Consider moving each actively queried db to a separate dedicated cluster.
    - Can reduce resource contention and improve q perf
- Skip compression analysis using `COPY`

### Redshift Spectrm
- Query the data in its original format directly from Amazon S3 in same region
- Spectrum supports structured and semi-structures data formats like: (avro, parquet, txtfile, sequencefile, rcfile, grok, opencsv, orc, ion, regesserde, json)

- Supports compressions and types (gzip - .gz, Snappy - .snappy, bzip2 - .bz2)

### Workload Management WLM

- Configure up to **8** query queues and set the # of queries that can run in each queue concurrently.
- Setup rules to route queries to particular queues based on the user running the query or labels that you specify.
- Configure memory allocated to each queue, so that large queries run in queues with more memory.
- WLM timeout property to limit long-running queries.

**Short Query Acceleration (SQA)**: prioritizes selected short-running queries ahead of longer-running queries in a dedicated space.
- If you enable SQA, you can reduce or eliminate WLM queues that are dedicated to running short queries.

### Time Seres Tables

- Organize data as a sequenbce of time-series tables for a fixed retention period and create a `UNION ALL` view to mask the fact that data is stored in multi tables
- In the sequence, each table should be identical but contain data for different time ranges
- Use `DROP TABLE` instead of running a large-scale `DELETE` and a subsequent `VACUUM` process to reclaim space.


### Tuning

You should assign distribution styles to achieve two goals:
- Collocate the rows from joining tables
When the rows for joining columns are on the same slices, less data needs to be moved during query execution.

- Distribute data evenly among the slices in a cluster.
If data is distributed evenly, workload can be allocated evenly to all the slices.

#### Distribution Styles

When you create a table, you designate one of three distribution styles: **KEY**, **ALL**, or **EVEN**. [Tutorial](https://docs.aws.amazon.com/redshift/latest/dg/tutorial-tuning-tables-distribution.html)

**KEY distribution**

The rows are distributed according to the values in one column. The leader node will attempt to place matching values on the same node slice. If you **distribute a pair of tables** on the **joining keys**, the **leader node collocates** the rows on the slices according to the values in the **joining columns** so that **matching values** from the common columns **are physically stored together**.

**ALL distribution**

A **copy** of the **entire table** is distributed to **every node**. Where EVEN distribution or KEY distribution place only a portion of a table's rows on each node, ALL distribution ensures that every row is collocated for every join that the table participates in.

**EVEN distribution**

The rows are **distributed** across the slices in a **round-robin** fasion, regardless of the values in any particular column. EVEN distribution is **appropriate when a table does not participate in joins** OR **when there is not a clear choice between KEY and ALL**. EVEN distribution is the **default style** of distribution.

### Sort Keys

When you create a table, you can specify one or more columns as the sort key. 

Amazon Redshift stores your data on disk in sorted order according to the sort key. 

How your data is sorted has an important effect on disk I/O, columnar compression, and query performance.

Choose sort keys for the SSB tables based on these best practices:

- If **recent** data is queried **most frequently**, specify the **timestamp column** as the leading column for the sort key.
- If you do **frequent range filtering** or **equality filtering** on one column, **specify that column** as the sort key.
- If you frequently **join** a (**dimension**) table, specify the **join column** as the sort key.

### Improving Queries

- Use **CASE Expression** to perform complex aggregations instead of selecting from the same table multiple times.

- **Avoid** using function in query predicates

- **USE** predicates as much as possible to restrict the dataset

- Use a **WHERE** clause to restrict the dataset

- Use **sort keys** in the **GROUP BY** clause to improve aggregations

- **Use subqueries** in cases where one table in the query is used only for predicate conditions and the subquery returns a small number of rows (less than about 200). The following example uses a subquery to avoid joining the LISTING table.

```sql
select sum(sales.qtysold)
from sales
where salesid in (select listid from listing where listtime > '2008-12-26');
```

- **Avoid** using `select *`. Include only the columns you specifically need.

- In the predicate, use the _least expensive operators_ that you can. [**Comparison Condition**](https://docs.aws.amazon.com/en_pv/redshift/latest/dg/r_comparison_condition.html) operators are preferable to `LIKE` operators. `LIKE` operators are still preferable to `SIMILAR TO` or `POSIX` Operators.

#### Best Practices - Upsert

- **Upsert** - Redshift **doesn't** support a single merge statement to insert and update data from a single data source.

- **Utilize staging table first** - Efficiently update and insert new data by loading your data into a stagign table first.