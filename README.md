# 🚀 Go API Gateway

A high-performance, modular API Gateway built with **Go**, using **Fiber** for routing and **FastHTTP** for external service communication.  
Designed for scalable microservice architectures with built-in authentication, monitoring, and circuit breaker support.



---

## 🛠️ Getting Started

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

Secrets (e.g., DB connection string and application key) are managed using Kubernetes Secrets:

apiVersion: v1
kind: Secret
metadata:
  name: go-app-secret
type: Opaque
data:
  APP_KEY: <base64-encoded-key>
  PG_CONNECTION: <base64-encoded-postgres-connection-string>
