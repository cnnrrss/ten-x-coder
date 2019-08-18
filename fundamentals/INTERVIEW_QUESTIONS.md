# Interview Questions

### System Design
LRU Cache \
Distributed System \
SQL / NoSql \
Load Balancing \
HAProxy \
Cache (Redis/memcached) \
Vertical vs Horizontal scaling \ 
Rate Limiter

Design Twitter \
Design Newsfeed \
Design WhatsApp \
Design LRU Cache

### How Does Questions?

How does context switching work? \
How does diffing work? \
How do filesystems work? \
How does TCP work? \
How does Web Sockets Work? \
How does map reduce work? \
How do distributed systems work?


#### Rate Limiter
**Local rate limiter (single host):**
We may use Token Bucket algorithm  \
We will also need to store buckets in memory and remove/add them on demand.  \
So, we will need a data structure for this, e.g. hash table (ConcurrentHashMap in Java, to deal with synchronization).

**Distributed rate limiter:**
We should make hosts talk to each other and share how many tokens each one of them consumed so far.  \
To implement communication between hosts we may use different approaches: 
- full mesh communication (everyone talks to everyone)
- gossip communication
- distributed cache
- coordination service

For communication protocol we can use TCP or UDP.