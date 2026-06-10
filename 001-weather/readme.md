# Project 01 - External API Gateway

## Goal

Build a production-style REST API in Go that fetches data from external services, transforms the response, and exposes a clean JSON API.

This project focuses on learning Go fundamentals used in backend engineering:

* HTTP servers
* Routing
* JSON serialization/deserialization
* HTTP clients
* Error handling
* Context and timeouts
* Dependency injection
* Testing
* Project structure

---

## Functional Requirements

### Endpoint

GET /country/{name}

Example:

GET /country/india

### Behaviour

1. Accept country name from user.
2. Call external REST API.
3. Parse external JSON response.
4. Transform data into internal response model.
5. Return simplified JSON response.

Example Response:

{
"name": "India",
"capital": "New Delhi",
"population": 1400000000,
"region": "Asia"
}

---

## Non Functional Requirements

* Request timeout support
* Structured logging
* Proper HTTP status codes
* Graceful error handling
* Configuration through environment variables
* Unit test coverage

---

## Project Structure

cmd/
internal/
handlers/
services/
clients/
models/
config/
logger/
tests/

---

## Learning Objectives

* Understand net/http
* Work with JSON payloads
* Create reusable HTTP clients
* Use contexts and cancellation
* Separate business logic from handlers
* Write unit tests

---

## Success Criteria

* Endpoint returns correct response
* Invalid inputs handled properly
* External API failures handled gracefully
* Timeout handling works
* Unit tests pass
* Linting passes
* README updated with architecture notes

---

## Future Enhancements

* Response caching
* Rate limiting
* Metrics
* OpenAPI/Swagger
* Docker support
* Circuit breaker pattern
