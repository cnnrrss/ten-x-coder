# Big Data

**3 major components**:
- Detect corrupt, inaccurate, or anomalies
- Correction
- Removal

**Strict vs fuzzy**

Strict - must be a postal code \
Fuzzy - partial match

#### Data Quality
- Validity: degree to which the measures conform
- Accuracy: The degree of conformity of a measure to a standard or a true value
- Completeness
- Consistency
- Uniformity

### Paid Tools
AWS Glue ETL \
Informatica \
Talend \
Trifacta Wrangler: (Run on EC2) \
Data Ladder

### Big Data Tools
Impala - interactive BI-like workloads, \
Hadoop - Implementation of MapReduce (DFS) \
Spark - execute batch queries as well, large community, Java/Python/Scala \
Hive - analyzing, summarizing and querying data stored in Hadoop file system \
Presto - BI-type queries, support for concurrent query workloads.
Pig - \
Giraph - \

### Opensource
OpenRefine (formerly GoogleRefine): free and open source \
Drake

#### Custom algorithms
Scala \
Java \
Python/Pandas/Numpy \
R

#### Machine Learning

MAchine learning is a good approach when we have an idea about what we're looking for (hence, not lendlease)

scikit-learn, spark mllib, tensorflow, keras

Spark SQL (ex: clean frames)

Mechanical Turk - Manual classification 

#### Strategies
**Bayes nets** \ 
**Support-vector machines** \
**Decision trees** \
**Hidden Markov models** \ 
**MapReduce** (good when data is regular) \
    - Friends in a network \
    - ranking web pages by importance \
    - For each MapReduce Problem: \
        1. A set of inputs. \
        2. A set of outputs. \
        3. A many-many relationship between the inputs and outputs, which describes which inputs are necessary to produce which outputs.
**Clustering**
    - Hierarchical Clustering
    - K Means Algorithms
    - BFR


**Definition of the Adwords Problem**
Of course, the decision regarding which ads to show must be made on-line. Thus, we are only going to consider on-line algorithms for solving the adwords problem, which is as follows.

• **Given**:
1. A set of bids by advertisers for search queries.
2. A click-through rate for each advertiser-query pair.
3. A budget for each advertiser. We shall assume budgets are for a
month, although any unit of time could be used.
4. A limit on the number of ads to be displayed with each search query.

• **Respond** to each search query with a set of advertisers such that:
1. The size of the set is no larger than the limit on the number of ads
per query.
2. Each advertiser has bid on the search query.
3. Each advertiser has enough budget left to pay for the ad if it is clicked upon.

✦ **Greedy Algorithms**: Many on-line algorithms are greedy, in the sense that they select their action at every step by minimizing some objective function.

✦ **Competitive Ratio**: We can measure the quality of an on-line algorithm by minimizing, over all possible inputs, the value of the result of the on- line algorithm compared with the value of the result of the best possible off-line algorithm.

✦ **Bipartite Matching**: This problem involves two sets of nodes and a set of edges between members of the two sets. The goal is to find a maximal matching – as large a set of edges as possible that includes no node more than once.


**Market Basket** - model of data used to describe a common form of many-many relationship between two kinds of objects. \
**Bloom filter** - The purpose of the Bloom filter is to allow through all stream elements whose keys are in S, while rejecting most of the stream elements whose keys are not in S.
**Page Rank** - 

**DFS / Cluster Computing**
**Feature Extraction**
**Total Info Awareness**
**Bonferroni's Principle**

**Useful to know**
- tf-idf term frequency–inverse document frequency
- hash functions
- indexes
- secondary storage (disk) and its effect on run time of algorithms
- base _e_ of natural logarithms
- multi-way joins & cost

**Finding Similar Items**

#### Search
- ElasticSearch
- Lucene / Solr



### Challenges

**Corruption** 

**Loss of Data**

**Project costs**: costs typically in the hundreds of thousands of dollars

**Time**: mastering large-scale data-cleansing software is time-consuming

**Security**: cross-validation requires sharing information, giving an application access across systems, including sensitive legacy systems

#### Steps
1. Monitor Errors
2. Standardize Your Processes
3. Validate Accuracy
4. Scrub for Duplicate Data
5. Analyze
6. Communicate