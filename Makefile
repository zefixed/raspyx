include .env
export

all: clear redis-create db-create migrate-up db-admin build run
.PHONY: all

migrate-up: ### migration up
	goose -dir migrations postgres '$(PG_URL)' up
.PHONY: migrate-up

migrate-down: ### migration down
	goose -dir migrations postgres '$(PG_URL)' down
.PHONY: migrate-down

db-create: ### Creating postgres db docker instance
	docker run -p $(PG_PORT):5432 --network=raspyx_raspyx-network --restart unless-stopped --name $(APP_NAME)db -e POSTGRES_PASSWORD=$(PG_PASSWORD) -v $(APP_NAME)_postgres_data:/var/lib/postgresql/data -d postgres:17-alpine
	docker exec $(APP_NAME)db sh -c 'until pg_isready -U postgres; do sleep 0.5; done'
	docker exec $(APP_NAME)db psql -U postgres -tAc "SELECT 1 FROM pg_database WHERE datname='$(APP_NAME)db'" | grep -q 1 || docker exec $(APP_NAME)db psql -U postgres -c "CREATE DATABASE $(APP_NAME)db"
.PHONY: db-create

db-delete: ### Deleting postgres db docker instance
	docker stop $(APP_NAME)db
	docker rm $(APP_NAME)db
.PHONY: db-delete

redis-create: ### Creating Redis docker instance
	docker run -p $(REDIS_PORT):6379 --network=raspyx_raspyx-network --restart unless-stopped --name $(APP_NAME)redis -e REDIS_PASSWORD=$(REDIS_PASSWORD) -d redis:7-alpine \
	redis-server --requirepass $(REDIS_PASSWORD)
.PHONY: redis-create

redis-delete: ### Deleting Redis docker instance
	docker stop $(APP_NAME)redis
	docker rm $(APP_NAME)redis
.PHONY: redis-delete

build: ### Building app
	go build -o ${APP_NAME} ./cmd/app
.PHONY: build

run: ### Running app
	./${APP_NAME}
.PHONY: run

clear: ### Cleaning up
	docker stop $(APP_NAME) $(APP_NAME)db $(APP_NAME)redis || true
	docker rm $(APP_NAME) $(APP_NAME)db $(APP_NAME)redis || true
.PHONY: clear

db-admin: ### Creating admin user in db
	ADMIN_PASSWORD_HASH=$$(htpasswd -nbB admin admin | cut -d: -f2 | sed 's/$$/\\$$/g'); \
	docker exec -it $(APP_NAME)db psql -U $(PG_USER) -d $(APP_NAME)db -c "INSERT INTO users (uuid, username, password_hash, access_level) VALUES ('00000000-0000-0000-0000-000000000000', 'admin', '$$ADMIN_PASSWORD_HASH', 99) ON CONFLICT DO NOTHING"
.PHONY: db-admin

swag: ### Generating swagger documentation
	swag init -g cmd/app/main.go
.PHONY: swag

grafana-create: ### Creating grafana docker instance
	docker run -d --name=$(APP_NAME)grafana -p $(GRAFANA_PORT):3000 --network=raspyx_raspyx-network -v $(APP_NAME)_grafana_data:/var/lib/grafana grafana/grafana
.PHONY: grafana-create

grafana-delete: ### Deleting grafana docker instance
	docker stop $(APP_NAME)grafana
	docker rm $(APP_NAME)grafana
.PHONY: grafana-delete

prometheus-create: ### Creating prometheus docker instance
	docker run -d --name=$(APP_NAME)prometheus -p $(PROMETHEUS_PORT):9090 --network=raspyx_raspyx-network --add-host=host.docker.internal:host-gateway -v $(APP_NAME)_prometheus_data:/prometheus -v ./prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus:v3.4.0
.PHONY: prometheus-create

prometheus-delete: ### Deleting prometheus docker instance
	docker stop $(APP_NAME)prometheus
	docker rm $(APP_NAME)prometheus
.PHONY: prometheus-delete

