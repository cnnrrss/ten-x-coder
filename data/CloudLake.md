# CloudLake

- Who is our audience (Data scientist/engineers, business analysts/users, developers)?

- Data input (Hadoop/NoSQL, Flat files, Streams, Apps)?


### Market
**AWS Glue** - Execute Python/PySpark from Glue ETL workflows
**Informatica** - Enterprise grade, costly, burdensome, low-level, requires specific data engineer to manage

### Data Quaity Dimensions

**Completeness**: Is all the intended data being produced in the data set or is any of it missing?

**Uniqueness**: Are any individual pieces of data from the dataset recorded more than once?

**Timeliness**: How long is the time difference between data capture and the real world event being captured?

**Validity**: Is the data presented in the correct and pre-defined format, type or range so as to be applicable to the given analytical task?

**Accuracy**: Does the data matches up with the real world object or event it describes, enabling correct conclusions to be drawn from it?

**Consistency**: Is the given dataset consistent and correlative with different representations of the same information across multiple datasets?

### Features
- **Linting**: 
    - Is document well-formed? 
    - Can we auto-correct? (i.e. append commas to rows, complete brackets/parenthesis, handle new lines)
- **Data statistics**: 
    - **Accuracy**: How do we confirm accuracy (ML?)
    - **Completeness**: % of rows missing column
    - **Consistency**: 
        - Anomaly Detection: Custom business rules to identify records that cross certain thresholds
        - Ad-hoc transformations
    - **Uniqueness**: % of unique rows
    - **Validation**:
        - % of rows with valid value (ex: autodetect emails / Ip addresses and validate)
        - Configure and apply business rules, return % of rows that meet configured criteria
    - **Timeliness**: Is this the most up-to-date version of the data?
        - Streaming data
        - Joining other datasets
- **Data Exploration**:
    - Search
    - Run facets to get to know data
        - Merge

1. Process
In general, most of the software tools assume that the input will not be malformed - none of the assessed tools check for a malformed file, and inthe case that they do, they usually refuse to import the malformed file altogether. In order to fit in the data validation process currentlydocumented, we probably should perform a simple quality check at two stages: Data Acquisition and Data Profiling / Data Quality

### Data Inputs Supported
- Phase 1:
    - Flat Files
        - CSV / TSV
        - JSON
- Phase 2:
    - Hadoop
    - DBs