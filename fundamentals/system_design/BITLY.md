# Design Bitly / TinyURL


- How to generate a unique ID for each URL?
`id  := crc32(url)`

- How would you generate unique IDs at scale (thousands of URL shortening request coming every second)?
    - Let's assume 10k/sec
    - 90% of all requests should respond in 10ms
    - Set db connection pool Max:20, Min:5, Idle 10000

- How would your service handle redirects?


- How would you support custom short URLs?


- How to delete expired URLs etc?
    - Allow them to register a subdomain
    - Each user assigned a subdomain (or alternatively the longURL host)
    - spotify.bit.ly/e43k3hf
    
- How to track click stats?

### Design
Service that given a long url, returns a short code for it
    - Handles redirects
    - Worker pool
    - Concurrent instances for horizontal scaling
    - Auto scale based on traffic (or if using a messaging queue, topic backlog)
