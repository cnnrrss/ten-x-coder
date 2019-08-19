# Design a Rate Limiter

- Limit the number of requests an entity can send to an API within a time window e.g., 15 requests per second.

- The rate limiting should work for a distributed setup, as the APIs are accessible through a cluster of servers.

- How would you handle throttling (soft and hard throttling etc.).