# Real-Time AI Search & Recommendation Platform

A high-performance, real-time event-driven architecture utilizing **Golang**, **Apache Kafka**, **Vespa.ai**, and **Next.js (React)**.

## System Architecture

```
[ Next.js Frontend ]
        │ (HTTP / Server Actions)
        ▼
[ Golang Backend Service ]
        │ (Produces Events)
        ▼
[ Apache Kafka (Message Broker) ]
        │ (Streams Data / Consumer)
        ▼
[ Vespa.ai (Search & Vector Engine) ] ◄── (Queried directly by Next.js / Go)
```

## Features

- **Asynchronous Telemetry Ingestion:** User interaction events are dispatched from Next.js to Golang microservices and pushed to Kafka topics.
- **Low-Latency Hybrid Vector Search:** Vespa.ai executes combined BM25 text queries and Approximate Nearest Neighbor (HNSW) tensor searches.
- **Idiomatic Go Clean Architecture:** Domain logic, HTTP handlers, Kafka producers, and Vespa Document API clients are loosely coupled and independently testable.
- **Next.js App Router:** Type-safe API proxy routes and server-rendered search interfaces.

## Getting Started

### Local Development with Docker Compose

```bash
docker-compose up --build
```

### Running Backend Unit Tests

```bash
cd backend
go test -v ./...
```
