# Data Cleansing

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
**Drake**

## Languages
**Scala / Java**: functional programming
**Python / Pandas / Numpy**: multi-purpose dynamic programming
**R**: Statistical analysis

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
Optimus - Cleansing, pre-processing, feature engineering, exploratory data analysis and easy ML with PySpark backend.
[CleanFrames](https://github.com/funkyminds/cleanframes)

## Manual Classification/Correction

Pay workers to improve data quality.

Mechanical Turk
Crowdflower

## Search
- ElasticSearch
- Lucene / Solr