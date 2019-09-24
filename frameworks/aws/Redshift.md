# Amazon Redshift

- Redshift is a columnar SQL Database
- Redshift is a fork of Postgres Database
- Redshift consists of one Leader Node and several Compute Nodes. Each compute node is made up of slices. Each slice is a computing unit.

Redshift delivers fast query performance on datasets ranging in
size from gigabytes to exabytes. Redshift uses columnar storage, data compression, and zone maps to reduce the amount of I/O needed to perform queries. It uses a massively parallel processing (MPP) data warehouse architecture to parallelize and distribute SQL operations to take advantage of all available resources

Amazon Redshift uses machine learning to deliver high throughout based on your workloads. Redshift utilizes sophisticated algorithms to predict incoming query run times, and assigns them to the optimal queue for the fastest processing.

### Tuning

You should assign distribution styles to achieve two goals:
- Collocate the rows from joining tables
When the rows for joining columns are on the same slices, less data needs to be moved during query execution.

- Distribute data evenly among the slices in a cluster.
If data is distributed evenly, workload can be allocated evenly to all the slices.

##### Distribution Styles

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