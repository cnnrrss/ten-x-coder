# Concurrency in Go

**2. Advantages of Goroutinues over threads**
- You can run more goroutines on a typical system than you can threads.
    - Goroutines have built-in primitives to communicate safely between themselves (channels)
    - Goroutines have faster startup time
    - Goroutines are multiplexed onto a smaller number of OS threads
        - OS threads map 1:1
    - Goroutines have growable segmented stacks
    - In java you can run 1000's or 10k's threads. In Go you can run hundreds of thousands or millions of goroutines.
        - Threads consume a lot of memory (≥ 1MB).

**2. Closures**
- Make some concurrent calculations easier to express
- They are just local functions.


**3. What do you know about GO in market trading operations?**

**4. REST and RESTFUL**
REST:
    - Client-server: Rest applications have a server exchanges documents to client over HTTP (typically JSON or XML)
    - Stateless: servers don't maintain any client state. Clients manage their application state. Their requests to servers contain all the information required to process them.
    - Cacheable: server must mark their responses as cacheable or not. 
    - Uniform interface: emphasis on a uniform interface between components
    - Layered-system: components in the system cannot "see" beyond their layer. You can easily add load-balancers and proxies to improve security & performance.

RESTful is typically used to refer to web services implementing such an architecture.

REST - is a style of software architecture. RESTful is an adjective, sometime "REST-like".

**5. What have you done with distributed technologies and systems?**
Mesos - Master-slave: node virtualization and container orchestration platform
Marathon - dynamicall allocate resources across mesos nodes
Chronos -
Zookeeper - 
Kubernetes Pods/Services

**6. What is the most complex system you’ve built with AWS?**


**7. Tell me about your experience with technologies such as MongoDB, Cassandra, Kafka, Docker or Kubernetes?**
MongoDB - We used mongoDB for our 'configuration' db. Users, accounts, campaigns, sites. Hosted on mongoLab. The advantage of mongoDB was that our schema was continually evolving and it gave us the flexibility of experimenting new features, etc...

Cassandra - cluster for our intensive analytics data (pages, views, likes, aggregates) better r/w throughput.

Kafka - was our centralized messaging bus. Almost all servers communicated via event-driven architecture (not client-server, but all internal services).
    - We also utilized the fact that this gave us out of the box metrics and analytics Big Query (MAU, Usage statistics, etc..)

Docker - All our services were built and packages with Docker. We used a mono repo with individual commands so the image / release version was the same but dependencies could be deployed independently leading to smaller builds and containers.

Kubernetes - I worked on a project to migrate our services from AWS/Mesos/etc.. to GCP Google Kubernetes Engine. For our small team, we found this product helped us minimize the amount of dev resources it took us to manage the operations of mesos/ec2/aws.

**8. How do you ensure data is kept secure and what’s your experience with encryption?**

**9. What tools are required to test your web API?**
Postman, curl + jq


**10. What do you know about Microservices?**
UNIX philosophy.

- Higly Maintainable and Testable
- Owned by small team: Break down monolithic components to individual focuses
- Loosely coupled: Encourage continuous delivery by decoupling dependencies.
- Independently deployable: Utilize containers and orchestration infrastructure to package, deploy and run independently.
- Often communicate with an event driven architecture.
- Teams must embrace devops

In theory:
- Simpler to deploy
- Simpler to understand
- Lighter on the IDE
- Reusability (one could argue more boilerplate)
- Faster defect isolation
- Minimized risk of change

Cons:
- Add complexity
- Cascading failures
- Testing can be complicated and tedious (advanced test env)
- Create team / information barriers
- Architectural complexity: Mitigate fault tolerance, network latency, variety of message formats and connections
- Distributed system can lead to duplicated effort and more complexity
- Additional effort into implementing mechanism of comm between the services.

**11. How do you build a scalable system while keeping it secure?**

**12. What is the difference between strings and runes?**
Rune is Golang's byte char. Rune is alias for int32 signed integer
String is a sequence of bytes. 

When iterating of a string go will return a rune as the val which is an int32 data type representation of an ASCII char.

Golang has two additional integer types called byte and rune that are aliases for uint8 and int32 data types respectively -

|Type|	Alias For|Representation|
|----|-------|------|
|byte|	uint8| ASCII|
|rune|	int32| Unicode / UTF-8|

```golang
var firstLetter = 'A' // Type inferred as `rune` (Default type for character values)
var lastLetter byte = 'Z'

var myByte byte = 'a'
var myRune rune = '♥'

fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)
// Output
a = 97 and ♥ = U+2665


// Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
var name = "Steve Jobs"

// Raw String (Can span multiple lines. Escape characters are not interpreted)
var bio = `Steve Jobs was an American entrepreneur.
           He was the CEO and co-founder of Apple Inc.`
```