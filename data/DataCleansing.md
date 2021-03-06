# Data Cleansing

**Chief measures**
Accuracy, Completeness, Consistency, Timeliness, Uniqueness, Validity

**3 major components**:
- **Detection**: corrupt, inaccurate, or anomalies
    - Invalid Types
    - Missing, null, blank columns
    - Duplicate rows
    - Provide validity/quality score (per column/per dataset)
- **Removal**
    - Reject rows with invalid values
    - Filter Unwanted Outliers
- **Correction**
    - Join relavant data on primary key where correct values may exist
    - Data rules, create mapping templates to reference in ETL pipeline for auto-correction
    - Fix Structural Errors: (types, inconsistent labels)
    - Handle Missing Data
    - ML?

Enriching the data, converting

Linting tool, inconsistent timestamps.

#### Why do we need clean data?

Better Data > Fancier Algorithms
Garbage in = Garbage out

If you have a properly cleaned dataset, even simple algorithms are better than machine learning.

####  What is Data Quality?
**Valid** | **Accurate** | **Complete** | **Consistent** | **Uniformed**

### Challenges

**Corruption**: Techniques such as auto-correction are prone to error by assumption which could lead to corrupted data.
**Loss of Data**: if you reject a row because it was determined incomplete, we may be losing valuable data.
**Project costs**: costs typically in the hundreds of thousands of dollars
**Time**: mastering large-scale data-cleansing software is time-consuming
**Security**: cross-validation requires sharing information, giving an application access across systems, including sensitive legacy systems

#### Steps for Data Cleansing
1. Monitor Errors
2. Standardize Your Processes
3. Validate Accuracy
4. Scrub for Duplicate Data
5. Analyze
6. Communicate


## Matching Strategies
Strict vs fuzzy

**Strict**: must be a postal code (i.e. regex)
**Fuzzy**: partial match (i.e. fzf)

## Tools
**AWS Glue ETL**: simple but less powerful 
**Informatica**: \$\$\$
**Talend**: \$\$\$
**Trifacta Wrangler**: (Run on EC2, maintenance $$$)
**Data Ladder**:
**Rulex Robot Data Correction**: https://rulexrdc.com/#/lessons/NufG9lGKsr1zzb4Re7b_7kLVkTrJhpwf

## Opensource Projects
**OpenRefine** (formerly GoogleRefine): free and open source \
[**Drake**](https://github.com/Factual/drake)
[**DataFu**](https://github.com/linkedin/datafu)
**SparkQL** 
    - [CleanFrames](https://github.com/funkyminds/cleanframes)
[**Optimus**](https://github.com/ironmussa/Optimus) - Cleansing, pre-processing, feature engineering, exploratory data analysis and easy ML with PySpark backend.

## Languages
**Scala / Java**: functional programming
**Python / Pandas / Numpy**: multi-purpose dynamic programming
**R**: Statistical analysis
[**Julia**](https://julialang.org)

## [Tools](https://github.com/academic/awesome-datascience#toolboxes---environment)

### Machine Learning
Machine learning is a good approach when we have an idea about what we're looking for.

scikit-learn - Machine Learning in Python
spark mllib
keras
Spark SQL
Tensorflow (Google) - an Open Source Software Library for Machine Intelligence
PyTorch (FB)
CoreML (Apple)
Pyro & Michelangelo (Uber)
[Gluon](https://github.com/gluon-api/gluon-api) + MXNet (Amazon)
DeepMind (Sonnet)

## Manual Classification/Correction

Pay workers to improve data quality.

Mechanical Turk
Crowdflower

## Search
- ElasticSearch
- Lucene / Solr


## Examples

Incorrect email addresses - significant impact on marketing campaigns
    - Email regex, user dedupe, etc.

Inaccurate personal details - Missed sales opportunities, customer complaints

**Improve Data Quality**: 
Prevention, Correction, Failure

Tasks that used to be done by data professionals, such as data experts, now done by operational workers that know the data best, is called selfservice. It requires workflow-driven, easy-to-use tools with an Excel-like UI and smart guidance.

## Wishlist

[Twitter Thread](https://twitter.com/heatherklus/status/1166512298971648000)

- python and r support parity via JupyterLab, RStudio, VS Code
- cloud native 
- individual and team level data access management 
- git access
- spin up and down spark clusters w/ and w/o gpu 
- ability for each user to install any pkg
- some analysis access component (think Airbnb knowledge repo)
- Simple and integrated experiment tracking/comparison
- An API that allows everything to be scripted
- Mobile support so I can fire off a run and monitor it while I’m at lunch/home
- The ability to get a text or email when a model has finished training.
- experiment tracking
- versioning
- project lifecycle management
- spark clusters 
- k8s based
- job scheduler
- install any package
- rich knowledge repo with search
- reusable data connections
- web app publisher/hosting 
- model API publisher/hosting
- Integration with git, Airflow, SAS, H2O, DataRobot, Matlab, …
- SSO, SAML, AAA
