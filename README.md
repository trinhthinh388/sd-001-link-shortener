# 🌈 System Design: Link Shortener

## Requirements

### Functional Requirements

The link shortener should be able to do the following operations:

- The server generates a unique short URL for each long URL.
- The server encodes the short URL for readability.
- The server persists the short URL in the data store.
- The server redirects the client to the original long URL against the short URL.

### Non-Functional Requirements

- High Availability
- Low Latency
- High Scalability
- Durable
- Fault Tolerant
