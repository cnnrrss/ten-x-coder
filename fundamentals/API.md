# API design

SOAP vs. REST
- SOAP is a protocol
- REST is an architectural style

- Went to a talk at Google Cloud Next and heard the founder of Swagger talk about API development best practices.

Good API design enables transition from Monolithic to Micro Service.

- OpenAPI (Swagger)

Scaling != "handling more transactions per second"
Scaling == "Making more customers happy per second"

Scaling Transaction
 - Decoupling Layers
 - Caching
 - Compression
 - Database Changes??
 - Streaming / Real-time
 - Smaller Payloads
 - SSL Optimization

Scaling Customers
 - Do all of transaction scaling +
 - Change API Structure
 - Change client design params

### Method Structure Design
- Add specific, remove, unbundle & rebundle
- Do methods / calls fit use cases
- Are you passing a lot of internal IDs around?
- Are some methods more expensive to execute than others?
- Heavy load calls are being made even if data is not needed
- Clients making many calls when the should be making one?

**Unexpected v Expected && Undesired v Desired**

- API call aggregation
- Offboarding calls
- Rate Limiting
    - Make sure key use cases are still feasible
    - Stop fringe high vol users from affecting everyone
    - Make sure documented
- SDK Provision

Out of the Box thinking can save a lot of dollars (and the planet)

https://3scale.net

### SOAP (Simple Object Access Protocol)
- More security, complexity, overhead
- XML
- Cannot be cached (unlike REST)e;
- SOAP has built-in ACID compliance 
 - ACID (Atomicity, Consistency, Isolation, Durability) compliance.
- Successful/retry logic for reliable messaging functionality

