![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Docker](https://img.shields.io/badge/docker-ready-blue)

# RestTest
A collection of simple RESTful applications implemented in different programming languages. The goal is to compare implementations side-by-side, focusing mainly on performance, but also on readability and ease of maintenance

## Features
- JSON serialization example
- PostgreSQL integration with read/write endpoints
- Redis caching with fallback logic
- Simulated concurrency endpoint
- Health check
- Graceful shutdown support

## Table of contents
- Tech Stack
- Setup
- Endpoints
- Project Structure
- License

## Tech Stack

Each language may include one or more web frameworks with relevant libraries

### Go:
- Frameworks:
    - [chi](Go/chi) - HTTP router - [Project repository](https://github.com/go-chi/chi)
- Libraries:
    - pgx - PostgreSQL driver
    - go-redis - Redis client
    - net/http - HTTP Server
### Rust:
- Frameworks:
    - [axum](Rust/axum) - Web framework - [Project repository](https://github.com/tokio-rs/axum)
- Libraries:
    - tokio - Async runtime

### Zig
- Frameworks:
    - [zap](Zig/zap) - HTTP Server - [Project repository](https://github.com/zigzap/zap)

### C++
- Frameworks:
    - [crow](C++/crow) - Micro web framework - [Project repository](https://github.com/CrowCpp/Crow)

## Setup
### Prerequisites
- Docker
- Docker Compose

### Quick Start
To build and run all implementations and supporting tool (eg. wrk for benchmarking), simply run:
```bash
docker-compose up -d --build
```
This will:
- Build all language implementations automatically via their respective Dockerfiles
- Start services in isolated containers
- Start wrk for benchmarking
- Save the results in json form inside [wrk-results](wrk-results)
> [!NOTE]
> Please make sure ports 8080-8083 on your machine are available

## Endpoints
All implementations aim to expose the same set of endpoints:

| Method | Endpoint           | Description                                           |
|--------|--------------------|-------------------------------------------------------|
| GET    | `/health`          | Returns a simple health check response                |
| GET    | `/user/json`       | Returns a static user JSON                            |
| GET    | `/user/db/{id}`    | Fetches a user from PostgreSQL                        |
| POST   | `/user/db`         | Writes a new user to PostgreSQL                       |
| GET    | `/user/cache/{id}` | Fetches a user from Redis or falls back to PostgreSQL |
| GET    | `/user/concurrent` | Simulates concurrent access (delayed)                 |

### Sample Payload for POST `/user/db`
```json
{
  "username": "alice",
  "email": "alice@example.com"
}
```

## Project Structure
```markdown
RestTest/
├── go/
│ └── chi/ # Go with chi framework
├── rust/
│ └── axum/ # Rust with axum
├── cpp/
│ └── crow/ # C++ with Crow
├── zig/
│ └── zap/ # Zig with Zap
├── wrk-scripts/ # Benchmarking scripts
├── wrk-results/ # JSON output from benchmarking
├── docker-compose.yml
└── README.md
```
As you can see, each implementation follows this pattern: `<Language>/<framework>/` (Note capitalization). \
Inside each you will find:
- Source code - main, routes, handlers, db/cache logic
- Dockerfile - for container definition

## Contributing
Want to add a new language or framework? Follow the existing structure:

- Create a new folder under `<language>/<framework>`
- Implement the required endpoints with matching behavior
- Use a Dockerfile to build the service
- Update `docker-compose.yml` accordingly
- Provide your benchmarking scripts in [wrk-scripts](wrk-scripts) following the other scripts' behavior

PRs and issues for performance optimizations or project structure are also welcome!

## License
This project is licensed under the MIT license. See [LICENSE](LICENSE) for details.