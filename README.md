# 💰 Qotera Backend — Financial Intelligence Engine

Qotera is a high-performance, containerized financial backend designed to automate expense tracking and budgeting by syncing transaction data from various sources (specifically SMS providers like Telebirr and CBE). Built with **Go** and **Fiber**, it provides a secure, efficient API for the Qotera mobile application.

---

## 🚀 Key Features

*   **Secure Authentication**: JWT-based auth with OTP verification flow.
*   **Transaction Syncing**: Specialized endpoints for syncing and deduplicating SMS-derived transactions.
*   **Budgeting Engine**: Category-wise limit setting and real-time "Spent vs Limit" calculations.
*   **Analytics**: Aggregated financial snapshots (Daily/Weekly/Monthly) powered by optimized Postgres queries.
*   **Cloud Ready**: One-click deployment to **Render** via Blueprint infrastructure-as-code.
*   **Dockerized Stack**: Fully contained development environment with Postgres and Redis.

---

## 🛠️ Technology Stack

*   **Language**: [Go](https://go.dev/) (Golang)
*   **Web Framework**: [Fiber](https://gofiber.io/) (Fast & Lightweight)
*   **ORM**: [GORM](https://gorm.io/) (PostgreSQL)
*   **Caching/OTP**: [Redis](https://redis.io/)
*   **Database**: [PostgreSQL](https://www.postgresql.org/)
*   **Infrastructure**: [Render](https://render.com/) (Blueprints) & [Docker](https://www.docker.com/)

---

## 📦 Project Structure

```text
.
├── cmd/                # Entry points (Server, Seeding, etc.)
├── internal/
│   ├── domain/         # Domain models (User, Transaction, Budget)
│   ├── handler/        # HTTP Handlers (Fiber controllers)
│   ├── middleware/     # Auth & Logger middleware
│   ├── repository/     # GORM data access layers
│   └── service/        # Core business logic
├── pkg/
│   ├── config/         # Environment configuration
│   ├── database/       # Connection helpers (Postgres, Redis)
│   └── utils/          # JWT, Validation, and Helpers
├── docs/               # Auto-generated Swagger documentation
└── render.yaml         # Infrastructure as Code
```

---

## 🔨 Getting Started

### Prerequisites
- Go 1.21+
- Docker & Docker Compose

### Local Development
1.  **Clone the repo**:
    ```bash
    git clone https://github.com/kaybee/qotera-backend.git
    cd qotera-backend
    ```

2.  **Start Services** (Postgres & Redis):
    ```bash
    docker-compose up -d
    ```

3.  **Run migrations and seed data**:
    ```bash
    go run cmd/seed/main.go
    ```

4.  **Run the server**:
    ```bash
    go run cmd/server/main.go
    ```
    The API will be available at `http://localhost:3000`.

---

## 📄 API Documentation
The API documentation is available via Swagger. Once the server is running, visit:
`http://localhost:3000/swagger/index.html`

---

## ☁️ Deployment
This repository is configured for effortless deployment on **Render**. Simple connect your repository to Render and it will automatically provision your Database, Redis, and Web Service using the `render.yaml` blueprint.

---

## 📜 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
