# System Design Interview

For better information, this place will never beat [this](https://github.com/donnemartin/system-design-primer) repo.

### Approach
**Step 1** — Understand the Goals. Clarifying ambiguities early in the interview is critical. ...
**Step 2** — Establish the Scope. ...
**Step 3** — Design for the Right Scale. ...
**Step 4** — Start High-Level, then Drill-Down. ...
**Step 5** — Data Structures and Algorithms (DS&A) ...
**Step 6** — Tradeoffs.

### Tips
- Don't go into details too early
- You don't want to be talking all the time
- (Flexibility) Don't have a set architecture in mind (Avoid silver bullets)
- K.I.S.S (don't focus on one service)
- Form your thoughts. Don't be specifically tied to one technology (NoSQL: Cassandra, Messaging Queue: Kafka)

### Technologies
- Be aware of current technologies
    - Dbs:
        - NoSQL
            - Amazon Dynamo DB
            - MongoDB
            - Cassandra
            - Redis
        - SQL
            - MySQL
            - Posgres
    - Gateway
        - AWS ELB
        - HAPROXY
    - Heartbeats
        - Zookeeper
    Protocols:
        - HTTP
        - TCP
        - UDP
        - gRPC

### Vocabulary
**API Gateway** - single entry point for all clients. Handles requests in one of two ways. Some requests are simply proxied/routed to the appropriate service. It 