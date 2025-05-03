# 🚀 Go API Gateway

A high-performance, modular API Gateway built with **Go**, using **Fiber** for routing and **FastHTTP** for external service communication.  
Designed for scalable microservice architectures with built-in authentication, monitoring, and circuit breaker support.

---

## ✨ Key Features

- ⚡ **Blazing fast** routing with Go and Fiber (`v1.23.8`)
- 🔄 Outbound communication via **FastHTTP**
- 🔐 **Optional authentication** per service
- 💥 **Circuit breaker** support to handle failing downstream services gracefully
- 🧩 **Modular service structure** – each service handled in a separate Go file
- 📊 Integrated **Prometheus** and **Grafana** monitoring
- 🧵 In-progress integration with **OpenTelemetry** and **Eager**
- ☁️ **Kubernetes-ready** with LoadBalancer support and base64-encoded secrets
- 🐳 Dockerized and ready to deploy via **Docker Hub**
- ⚙️ Configurable using `.env` (based on `env_example`)

---

## 🛠️ Getting Started

### 1️⃣ Environment Setup

Start by copying the environment example file:

```bash
cp env_example .env

docker build -t yourdockerhub/go-api-gateway:latest .
docker push yourdockerhub/go-api-gateway:latest

## migration
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);