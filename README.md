# 📈 StockPulse – Real-Time Inventory & Order Management System

**StockPulse** is a production-grade backend engineering project that simulates a real-world order and inventory management system built using modern backend technologies. Designed for engineers who want to showcase hands-on expertise with **microservices**, **gRPC**, **Kafka**, **Golang**, **Redis**, **Docker**, **Kubernetes**, and **MySQL**.

It now also includes a **React + Tailwind** frontend to visualize and interact with the platform!

---

## ⚙️ Tech Stack

| Layer        | Tech Used                                  |
|--------------|--------------------------------------------|
| Language     | Go (Golang)                                |
| APIs         | REST (Order Service), gRPC (Inventory)     |
| Messaging    | Kafka (event streaming)                    |
| Caching      | Redis                                       |
| Persistence  | MySQL                                      |
| Frontend     | React.js + Tailwind CSS                    |
| DevOps       | Docker, Docker Compose, Kubernetes (k8s)   |

---

## 📌 Features

### ✅ Core Services

- **Order Service** (REST)
  - Accepts customer orders
  - Emits `OrderPlaced` Kafka event
  - Stores order in MySQL

- **Inventory Service** (gRPC + Redis)
  - Handles stock checking/reservation
  - Caches product stock in Redis
  - Syncs with MySQL for persistence

- **Kafka Consumer**
  - Listens to `OrderPlaced` events
  - Triggers stock reservation

### 🖥️ Frontend (React)

- Order form with product selection
- Shows success/failure responses
- Inventory dashboard (upcoming)
- API connectivity with backend

---

## 🧠 Architecture

```mermaid
graph TD
  A[User] -->|REST /order| B(Order Service)
  B -->|Kafka: order.placed| C(Kafka Broker)
  C --> D(Inventory Service via gRPC)
  D -->|Redis + MySQL| E(DB Layer)

  subgraph UI
    A
  end

  subgraph Backend
    B
    C
    D
    E
  end
