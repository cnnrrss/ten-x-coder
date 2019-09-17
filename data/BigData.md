# Big Data

[Awesome BD](https://github.com/onurakpolat/awesome-bigdata)


### Big Data Tools
Impala: interactive BI-like workloads,
Hadoop: Implementation of MapReduce (DFS)
Spark: execute batch queries as well, large community, Java/Python/Scala
Hive: analyzing, summarizing and querying data stored in Hadoop file system
Presto: BI-type queries, support for concurrent query workloads.
Beam: 
Pig: high-level data flow scripting language that supports standalone scripts and provides an interactive shell which executes on Hadoop 
Giraph: 

### Big Data & ML Strategies

**Bayes nets**:

**Support-vector machines**:

**Decision trees**:

**Hidden Markov models**:

**MapReduce**: (good when data is regular)
    - Friends in a network
    - ranking web pages by importance
    - For each MapReduce Problem:
        1. A set of inputs.
        2. A set of outputs.
        3. A many-many relationship between the inputs and outputs, which describes which inputs are necessary to produce which outputs.

**Clustering**:
    - Hierarchical Clustering
    - K Means Algorithms
    - BFR

#### Definition of the Adwords Problem
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


**Market Basket** - model of data used to describe a common form of many-many relationship between two kinds of objects.

**Bloom filter** - The purpose of the Bloom filter is to allow through all stream elements whose keys are in S, while rejecting most of the stream elements whose keys are not in S.

**Page Rank** - 

**DFS / Cluster Computing**

**Feature Extraction**

**Total Info Awareness**

**Bonferroni's Principle**

### Useful to know
- tf-idf term frequency–inverse document frequency
- hash functions
- indexes
- secondary storage (disk) and its effect on run time of algorithms
- base _e_ of natural logarithms
- multi-way joins & cost