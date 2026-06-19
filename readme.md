# Go System Design

A collection of backend engineering and system design projects built in Go.

The goal of this repository is to learn Go through practical, production-style applications rather than isolated coding exercises. Each project focuses on a specific backend engineering concept such as HTTP APIs, concurrency, rate limiting, distributed systems, caching, messaging, and scalability.

Every project includes:

* Problem statement
* Design considerations
* Architecture notes
* Implementation
* Testing
* Future improvements

---

## Learning Goals

* Build idiomatic Go applications
* Understand backend system design fundamentals
* Learn concurrency using goroutines and channels
* Work with HTTP servers and clients
* Design reliable and maintainable services
* Practice production-grade project structure
* Document engineering trade-offs and decisions

---

## Repository Structure

```text
go-system-design/
├── 001-weather/
├── 002-rate-limiter/
├── 003-job-queue/
├── 004-chat-server/
├── 005-url-shortener/
└── ...
```

Each folder represents an independent project and contains its own README describing:

* Requirements
* Design
* Implementation details
* Lessons learned

---

## Project Roadmap

### Foundation

* [ ] 001 Weather API Gateway
* [ ] 002 Rate Limiter [![Live Sandbox](https://img.shields.io/badge/Live-Interactive_Sandbox-7c3aed?style=flat-square)](https://vasu1705.github.io/go-system-design/002-rate-limiter/concept-guide.html)
* [ ] 003 Worker Pool
* [ ] 004 Concurrent File Processor
* [ ] 005 Chat Server

### Backend Services

* [ ] 006 URL Shortener
* [ ] 007 Notification Service
* [ ] 008 API Gateway
* [ ] 009 Search Service
* [ ] 010 Authentication Service

### Distributed Systems

* [ ] 011 Distributed Cache
* [ ] 012 Event Bus
* [ ] 013 Task Queue
* [ ] 014 Log Aggregation Service
* [ ] 015 Service Discovery

---

## Engineering Checklist

Every project should aim to include:

### API Design

* Request validation
* Consistent JSON responses
* Proper HTTP status codes
* Error handling

### Go Practices

* Clear package structure
* Dependency injection
* Context propagation
* Interfaces where appropriate

### Reliability

* Timeouts
* Graceful shutdown
* Structured logging
* Panic recovery

### Testing

* Unit tests
* Integration tests where applicable
* Mock external dependencies

### Documentation

* Architecture overview
* Design decisions
* Trade-offs
* Future improvements

---

## Running Projects

From the repository root:

```bash
go run ./001-weather/cmd
```

Run all tests:

```bash
go test ./...
```

Format code:

```bash
go fmt ./...
```

Static analysis:

```bash
go vet ./...
```

---

## Why This Repository Exists

This repository serves as a public engineering journal and portfolio of my journey learning Go, backend engineering, and system design through hands-on implementation.

The focus is not only on making software work, but also on understanding the architectural decisions, trade-offs, and operational concerns involved in building production systems.
