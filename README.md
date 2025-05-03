<style>
  .title {
    width: 100%;
    height: 100px;
    padding: 20px;
    position: relative;
    margin-bottom: 40px;
    display: flex;
    overlfow: hidden;
    border-radius: 16px;
    flex-direction: column;
    justify-content: center;
    background-color: #0068d6;
  }
  .title::before {
    content: "";
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border-radius: 16px;
    pointer-events: none;
    position: absolute;
    background-image: radial-gradient(88% 100% at top,rgba(255,255,255,.5),rgba(255,255,255,0));
  }
  .title::after {
    content: "";
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    border-radius: 16px;
    pointer-events: none;
    position: absolute;
    background-image: url("./assets/noise.webp");
    background-size: 30%;
    opacity: .2;
  }

  .title h1 {
    font-size: 2.8rem;
    font-weight: 700;
    border: none;
    color: #f0f0f0;
    margin: 0;
  }
</style>
<div class="title">
  <h1>🌈 System Design</h1>
</div>

# Link Shortener

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
