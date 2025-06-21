# Go Gin Rate Limiting Demo

This project demonstrates a simple **HTTP API** built using **Golang** and the **Gin** web framework, showcasing the implementation of **rate limiting middleware** without requiring any external database or storage.  
It includes multiple **stateless dummy endpoints** to simulate real-world use cases such as time lookup, UUID generation, and client IP detection — all protected by configurable rate limit rules.

---


## ✨ Features

This demo application is designed to be lightweight, self-contained, and easy to extend. It provides a set of public API endpoints that are protected using **rate limiting** logic, ideal for learning or validating throttling behavior in a real-world API scenario.  
Below are the key features included in the project:
- **⚙️ Built with Gin** — minimal, fast, and idiomatic web framework in Go.
- **🚦 Rate Limiting Middleware** — configurable using token bucket algorithm (`golang.org/x/time/rate`).
- **🔒 Per-IP Rate Limiting** — request limits are enforced per client IP.
- **🧪 Stateless & Testable** — all endpoints are fully testable and don't rely on external services.
- **🔁 Simple Endpoint Examples**:
  - `GET /api/ping` — health check
  - `GET /api/time` — returns current server time
- **🧪 Unit tested** — includes basic tests for rate limit success and failure scenarios.
---


## 🤖 Tech Stack

This project utilizes a concise and efficient set of technologies to demonstrate rate limiting in a Go-based API. The stack is intentionally minimal to ensure clarity, speed, and ease of experimentation.

| **Component**             | **Description**                                                                             |
|---------------------------|---------------------------------------------------------------------------------------------|
| **Language**              | Go (Golang), a statically typed, compiled language known for concurrency and efficiency     |
| **Web Framework**         | Gin, a fast and minimalist HTTP web framework for Go                                        |
| **Rate Limiting**         | `golang.org/x/time/rate` — token-bucket rate limiter to control API usage frequency         |

---

## 🧱 Architecture Overview

This project adopts a **modular and clean architecture** pattern to ensure scalability, readability, and testability. Components are separated based on their responsibilities, promoting a clear structure that aligns with common Go project conventions.

```bash
📁 go-api-rate-limiter/
├── 📂cmd/                                  # Contains the application's entry point.
├── 📂internal/                             # Core domain logic and business use cases, organized by module
│   └── 📂handler/                          # HTTP handlers that process incoming API requests
├── 📂pkg/                                  # Reusable utility and middleware packages shared across modules
│   └── 📂middleware/                       # Request processing middleware
│       ├── 📂headers/                      # Manages request headers like CORS and security
│       └── 📂ratelimiter/                  # Implements API rate limiting based on IP, path, and method
├── 📂routes/                               # Route definitions, groups APIs, and applies middleware per route scope
└── 📂tests/                                # Contains unit or integration tests for business logic
```

---

## 🛠️ Installation & Setup  

Follow the instructions below to get the project up and running in your local development environment.  

### ✅ Prerequisites

Make sure the following tools are installed on your system:

| **Tool**                                                      | **Description**                           |
|---------------------------------------------------------------|-------------------------------------------|
| [Go](https://go.dev/dl/)                                      | Go programming language (v1.20+)          |
| [Make](https://www.gnu.org/software/make/)                    | Build automation tool (`make`)            |


### 🔁 Clone the Project  

Clone the repository:  

```bash
git clone https://github.com/yoanesber/Go-API-Rate-Limiter.git
cd Go-API-Rate-Limiter
```

### ⚙️ Configure `.env` File  

Create a `.env` file in the root of the project to configure basic runtime settings:  

```properties
# Application configuration
ENV=PRODUCTION
API_VERSION=1.0
PORT=1000
IS_SSL=FALSE
```

## 🚀 Running the Application  

Follow these simple steps to run the application locally:  

- **Notes**:  
  - All commands are defined in the `Makefile`.
  - To run using `make`, ensure that `make` is installed on your system.
  - Ensure you have `Go` installed on your system

### 📦 Install Dependencies

Make sure all Go modules are properly installed:  

```bash
make tidy
```

### 🧪 Run Unit Tests

```bash
make test
```

### 🔧 Run Locally

Run the app locally:

```bash
make run
```

### 🟢 Application is Running

Now your application is accessible at:
```bash
http://localhost:1000
```

---

## 🧪 Testing Scenarios  

Below are the key scenarios to validate the rate limiting behavior of the API:  

### Scenario 1: Successful Login

**Endpoint**: 

```http
GET http://localhost:1000/api/ping
```

**Rate Limit:**: Possible 2 requests every 5 seconds, with a maximum burst of 2 requests

**✅ Expected Response**:

```json
{
    "message": "pong"
}
```

### Scenario 2: Exceed Limit and Block

**Endpoint**: 

```http
GET http://localhost:1000/api/time
```

**Rate Limit**: Possible 2 requests every 5 seconds, with a maximum burst of 2 requests

**Test Flow**:
- Send 5 quick successive requests
- First 1–2 requests → `200 OK`
- Remaining requests → `429 Too Many Requests`

**❌ Expected Response**:
```json
{
    "error": "Too Many Requests",
    "message": "You have exceeded the rate limit. Please try again later."
}
```
