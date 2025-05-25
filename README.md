# Raspyx

This is a REST API built using the Go programming language, leveraging the [Gin](https://github.com/gin-gonic/gin) framework for handling HTTP requests. The API provides several endpoints for working with schedule.

## 📖 Table of Contents

- [📄 Description](#-description)
- [✨ Features](#-features)
- [🛠️ Technologies](#%EF%B8%8F-technologies)
- [📥 Installation and Setup](#-installation-and-setup)
  - [🖥️ Local Setup](#%EF%B8%8F-local-setup)
  - [📦 Docker Setup](#-docker-setup)
- [🌐 Base endpoints](#-base-endpoints)
- [✅ Testing](#-testing)
- [📜 License](#-license)

## 📄 Description

This API provides functionality to work with schedule. It includes features for creating, retrieving, updating, and deleting data.

## ✨ Features

- Supports GET, POST, PUT, DELETE methods
- Authentication using JWT
- Database connection with PostgreSQL
- Database migration with goose
- Caching with Redis
- Message streaming with Kafka
- Request and error logging
- Swagger documentation
- Grafana and Prometheus monitoring
- CI/CD integration with Jenkins

## 🛠️ Technologies

- Go (1.24)
- Gin
- JWT
- PostgreSQL
- Goose
- Redis
- Kafka
- Docker
- Grafana
- Prometheus
- Jenkins

## 📥 Installation and Setup

1. Clone the repository:

```bash
git clone https://github.com/zefixed/raspyx.git
cd raspyx
```

> ❗ Before using rename `.env.example` to `.env` and set up your parameters

### 🖥️ Local Setup

To run the application locally, follow these steps:

1. Install go

   Arch

   ```bash
   yay -Sy go
   ```

   Debian

   ```bash
   sudo apt install golang
   ```

2. _(Optional)_ Installing swag for docs generating

   > 💡 _This step is only needed if you want to regenerate documentation (e.g. change host in main.go)_

   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

   ```bash
   make swag
   ```

3. Run app

   ```bash
   make all
   ```

   > 💡 `make db-admin` _creates admin:admin user_

   > ❗ _admin:admin user should only be used to create another user with admin rights and should be deleted after its creation_

### 📦 Docker Setup

To run the application with Docker, follow these steps:

1. Run the docker-compose:

   ```bash
   docker compose up --build -d
   ```

## 🌐 Base Endpoints
After starting the service, the following endpoints will be available:

1. **API**  
   Base URL: `http://localhost:8080/raspyx`  
   [API Documentation](http://localhost:8080/raspyx/swagger/index.html)

2. **Monitoring**
   - Prometheus metrics: http://localhost:9090  
     *(Available only on local installations. To enable it in a Docker setup, uncomment the `ports` section in `docker-compose.yaml`.)*
   - Grafana dashboard: http://localhost:3000  
     *Default username and password: admin/admin*

3. **Health Checks**
   - http://localhost:8080/raspyx/ping


## ✅ Testing

To test the API, you can use [Postman](https://www.postman.com/) or [cURL](https://curl.se/). You can also set up unit tests in the project using:

```bash
go test -v ./...
```

## 📜 License

This project is licensed under the GNU License v3 - see the [LICENSE](LICENSE) file for details.
