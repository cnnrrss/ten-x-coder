# Spark

Spark applications consist of a driver program that run the user’s main function and executes various parallel operations on a cluster.

The main abstraction Spark provides is a resilient distributed dataset (RDD), which is a collection of elements partitioned across the nodes of the cluster that can be operated on in parallel.

## Spark vs MapReduce
Spark is based on MapReduce but unlike MapReduce, it doesn’t shuffle data from one cluster to another, Spark has in-memory processing which makes it faster than MapReduce but still scalable.

Just like MapReduce spark works on distributed computing, it takes the code and Driver program creates a job and submits it to DAG Scheduler.

Spark uses master/slave architecture, the master coordinates and distributes the job and rest all distributed systems are slave worker. The master system is called “Driver

### Why Spark?
**Speed**: Though spark is based on MapReduce, it is 10 times faster than Hadoop when it comes to big data processing.
**Usability**: Spark supports multiple languages thus making it easier to work with.
**Sophisticated Analytics**: Spark provides a complex algorithm for Big Data Analytics and Machine Learning.
**In-Memory Processing**: Unlike Hadoop, Spark doesn’t move data in and out of the cluster.
**Lazy Evaluation**: It means that spark waits for the code to complete and then process the instruction in the most efficient way possible.
**Fault Tolerance**: Spark has improved fault tolerance than Hadoop. Both storage and computation can tolerate failure by backing up to another node.

#### Libraries

**MLlib** is the library that provides machine learning capabilities to spark.
**GraphX** is for Graph creation and processing.
**Spark SQL** and Data frames library are for performing SQL operations on data.
**Spark stream** library is for real-time streaming data processing.

### SDKs

Scala, Java, Python

##### Companies using Spark
Amazon \
Alibaba Taobao \
Baidu \
eBay Inc. \
Hitachi Solutions \
IBM Almaden \
Nokia Solutions and Networks \
NTT DATA \
Simba Technologies \
Stanford Dawn \
Trip Advisor \
Yahoo