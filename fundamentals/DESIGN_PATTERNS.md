# Design Patterns


Types:
- Creation
- Structural
- Behavioral.

    Different from an algorithm because applying the same pattern to two separate problems could yield a different result. An algorithm is more like a recipe
    https://en.wikipedia.org/wiki/Software_design_pattern

    There are also "domain specific" patterns:
- Web
- Ui design patterns
- Platform specific
- MVC
- MVVM
- etc..


### Fluent-Style
iA fluent interface is a method for designing OO APIs based extensively on **method chaining** with the goal of making the readability of the source code close to that of ordinary written prose. \
Essentially creating a domain-specific language within the interface. An example is JMock testing framework:

```bash
mock.expects(once()).method("m").with( or(stringContains("hello"),
                                          stringContains("howdy")) );
```


### Creation
i**Factory Method**
Define an interface for creating a single object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.

**Abstract Factory**
Provide an interface for creating families of related or dependent objects without specifying their concrete classes.

**Builder**
Separate the construction of a complex object from its representation, allowing the same construction process to create various representations.

**Prototype**
Specify the kinds of objects to create using a prototypical instance, and create new objects from the 'skeleton' of an existing object, thus boosting performance and keeping memory footprints to a minimum.

**Singleton**
Ensure a class has only one instance, and provide a global point of access to it.
NOTE: Normally makes it difficult to test.

**Dependency Injection**
A class accepts the objects it requires from an injector instead of creating the objects directly.

**Object Pool**
Avoid expensive acquisition and release of resources by recycling objects that are no longer in use. Can be considered a generalisation of connection pool and thread pool patterns.

**Lazy Initialization**
Tactic of delaying the creation of an object, the calculation of a value, or some other expensive process until the first time it is needed. This pattern appears in the GoF catalog as "virtual proxy", an implementation strategy for the Proxy pattern.

**Multiton**
Ensure a class has only named instances, and provide a global point of access to them.

**Resource acquisition is initialization**
Ensure that resources are properly released by tying them to the lifespan of suitable objects.

### Structural
**Adapter/Wrapper/Translator**
Convert the interface of a class into another interface clients expect. An adapter lets classes work together that could not otherwise because of incompatible interfaces. The enterprise integration pattern equivalent is the translator.
**Bridge**
Decouple an abstraction from its implementation allowing the two to vary independently.	
**Composite**
**Decorator**
**Facade**
**Flyweight**
**Proxy**


### Behavioral
**Chain of Responsibility**
**Command**
**Iterator**
**Mediator**
**Memento**
**Observer**
**State**
**Strategy Pattern**
[Go example](https://github.com/tmrts/go-patterns/blob/master/behavioral/strategy.md)

Strategy behavioral design pattern enables an algorithm's behavior to be selected at runtime.

It defines algorithms, encapsulates them, and uses them interchangeably.

Implementation of an interchangeable operator object that operates on integers.

Rules of Thumb:
- Strategy pattern is similar to Template pattern except in its granularity.
- Strategy pattern lets you change the guts of an object. Decorator pattern lets you change the skin

**Template Method**
**Visitor**

#### Synchronization Patterns
**Condition Variable**
**Lock/Mutex**
**Monitor**
**Read-Write Lock**
**Semaphore**
i
#### Concurrency Patterns
**N-Barrier**
**Bounded Parallelism**
**Broadcast**
**Coroutine (Goroutines)**
**Generators**
**Reactor**
**Parallelism**
**Producer / Consumer**
i
### Messaging Patterns (algorithms?)
**Fan-In**
Funnels tasks to a work sink (e.g. server)

**Fan-Out**
Distributes tasks among workers (e.g. producer)

**Futures & Promises**

**Publish/Subscribe**

**Push & Pull**

### Stability Patterns
**Bulkheads**
**Circuit-Breaker**
[Go example](https://github.com/sony/gobreaker)
**Deadline**
**Fail-Fast**
**Handshaking**
**Steady-State**
i
### Profiling Patterns
**Timing Functions**

### Idioms
*i*Functional Options**
Functional options are a method of implementing clean/eloquent APIs in Go. Options implemented as a function set the state of that option.

### Anti-Patterns
 - Cascading Failures
