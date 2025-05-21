include .env
export

RED="\\033[31m"
GREEN="\\033[32m"
RESET="\\033[0m"

define delete-docker-service
	@if docker inspect ${APP_NAME}$1 >/dev/null 2>&1; then \
		docker stop ${APP_NAME}$1 >/dev/null 2>&1; \
		docker rm ${APP_NAME}$1 >/dev/null 2>&1; \
		echo -e "${GREEN}${APP_NAME}$1 deleted${RESET}"; \
	else \
		echo -e "${RED}${APP_NAME}$1 does not exist${RESET}"; \
	fi
endef

all: clear network-create redis-create db-create migrate-up db-admin-create prometheus-create grafana-create build run
.PHONY: all

migrate-up: ### migration up
	@OUT=$$(goose -dir migrations postgres '${PG_URL}' up 2>&1); \
	if echo "$$OUT" | grep "successfully migrated" >/dev/null 2>&1; then \
	  	echo -ne "${GREEN}${APP_NAME}db "; \
	  	echo "$$OUT" | awk 'END {for (i=4; i<=NF; i++) printf "%s%s", $$i, (i<NF ? " " : "")}'; \
		echo -e "${RESET}"; \
	elif echo "$$OUT" | grep "no migrations to run" >/dev/null 2>&1; then \
		echo -ne "${GREEN}${APP_NAME}db "; \
	  	echo "$$OUT" | awk 'END {for (i=4; i<=NF; i++) printf "%s%s", $$i, (i<NF ? " " : "")}'; \
		echo -e "${RESET}"; \
	else \
		echo -e "${RED}${APP_NAME}db migration error"; \
		echo -e "$$OUT${RESET}"; \
	fi
.PHONY: migrate-up

migrate-down: ### migration down
	@OUT=$$(goose -dir migrations postgres '$(PG_URL)' down 2>&1); \
	if echo "$$OUT" | grep "OK" >/dev/null 2>&1; then \
	  	echo -ne "${GREEN}${APP_NAME}db migrated down to "; \
	  	echo "$$OUT" | awk 'END {for (i=4; i<=4; i++) printf "%s%s", $$i, (i<NF ? " " : "")}'; \
		echo -e "${RESET}"; \
	else \
		echo -e "${RED}${APP_NAME}db migration error"; \
		echo -e "$$OUT ${RESET}"; \
	fi
.PHONY: migrate-down

db-create: ### Creating postgres db docker instance
	@if docker inspect ${APP_NAME}db >/dev/null 2>&1; then \
  		echo -e "${RED}${APP_NAME}db already exists${RESET}"; \
  	else \
		docker run -p $(PG_PORT):5432 --network=${APP_NAME}-network --restart unless-stopped --name $(APP_NAME)db -e POSTGRES_PASSWORD=$(PG_PASSWORD) -v $(APP_NAME)_postgres_data:/var/lib/postgresql/data -d postgres:17-alpine >/dev/null 2>&1; \
		docker exec $(APP_NAME)db sh -c 'until pg_isready -U postgres; do sleep 0.5; done' >/dev/null 2>&1; \
		docker exec $(APP_NAME)db psql -U postgres -tAc "SELECT 1 FROM pg_database WHERE datname='$(APP_NAME)db'" | grep -q 1 || docker exec $(APP_NAME)db psql -U postgres -c "CREATE DATABASE $(APP_NAME)db" >/dev/null 2>&1; \
		echo -e "${GREEN}${APP_NAME}db created${RESET}"; \
	fi
.PHONY: db-create

db-delete: ### Deleting postgres db docker instance
	$(call delete-docker-service,db)
.PHONY: db-delete

redis-create: ### Creating Redis docker instance
	@if docker inspect ${APP_NAME}redis >/dev/null 2>&1; then \
  		echo -e "${RED}${APP_NAME}redis already exists${RESET}"; \
  	else \
		docker run -p $(REDIS_PORT):6379 --network=${APP_NAME}-network --restart unless-stopped --name $(APP_NAME)redis -e REDIS_PASSWORD=$(REDIS_PASSWORD) -d redis:7-alpine >/dev/null 2>&1; \
		echo -e "${GREEN}${APP_NAME}redis created${RESET}"; \
	fi
.PHONY: redis-create

redis-delete: ### Deleting Redis docker instance
	$(call delete-docker-service,redis)
.PHONY: redis-delete

build: ### Building app
	@echo -e "${GREEN}${APP_NAME} is building...${RESET}"; \
	OUT=$$(go build -o ${APP_NAME} ./cmd/app 2>&1); \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then \
	  	echo -e "${GREEN}${APP_NAME} builded successfully${RESET}"; \
	else \
	  	echo -ne "${RED}${APP_NAME} building error: "; \
	  	echo -e "$$OUT${RESET}"; \
	fi
.PHONY: build

run: ### Running app
	@./${APP_NAME}
.PHONY: run

db-admin-create: ### Creating admin user in db
	@ADMIN_PASSWORD_HASH=$$(htpasswd -nbB admin admin | cut -d: -f2 | sed 's/$$/\\$$/g'); \
	if docker exec -it $(APP_NAME)db psql -U $(PG_USER) -d $(APP_NAME)db -c "SELECT * FROM users WHERE uuid = '00000000-0000-0000-0000-000000000000'" | grep admin >/dev/null 2>&1; then \
		echo -e "${RED}admin already exists, DO NOT FORGET TO DELETE HIM${RESET}"; \
	else \
		docker exec -it $(APP_NAME)db psql -U $(PG_USER) -d $(APP_NAME)db -c "INSERT INTO users (uuid, username, password_hash, access_level) VALUES ('00000000-0000-0000-0000-000000000000', 'admin', '$$ADMIN_PASSWORD_HASH', 99)" >/dev/null 2>&1; \
		echo -e "${GREEN}admin created, ${RED}DO NOT FORGET TO DELETE HIM${RESET}"; \
	fi
.PHONY: db-admin-create

db-admin-delete: ### Creating admin user in db
	@docker exec -it $(APP_NAME)db psql -U $(PG_USER) -d $(APP_NAME)db -c "DELETE FROM users WHERE uuid = '00000000-0000-0000-0000-000000000000'" >/dev/null 2>&1; \
	echo -e "${GREEN}admin deleted${RESET}"
.PHONY: db-admin-delete

swag: ### Generating swagger documentation
	@OUT=$$(swag init -g cmd/app/main.go 2>&1); \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then \
	  	echo -e "${GREEN}documentation successfully generated${RESET}"; \
	else \
	  	echo -e "${RED}documentation generating error: "; \
	  	echo -e "$$OUT${RESET}"; \
	fi
.PHONY: swag

grafana-create: ### Creating grafana docker instance
	@if docker inspect ${APP_NAME}grafana >/dev/null 2>&1; then \
  		echo -e "${RED}${APP_NAME}grafana already exists${RESET}"; \
  	else \
  	  	docker run -d --name=$(APP_NAME)grafana -p $(GRAFANA_PORT):3000 --network=${APP_NAME}-network -v $(APP_NAME)_grafana_data:/var/lib/grafana grafana/grafana >/dev/null 2>&1; \
		echo -e "${GREEN}${APP_NAME}grafana created${RESET}"; \
	fi
.PHONY: grafana-create

grafana-delete: ### Deleting grafana docker instance
	$(call delete-docker-service,grafana)
.PHONY: grafana-delete

prometheus-create: ### Creating prometheus docker instance
	@if docker inspect ${APP_NAME}prometheus >/dev/null 2>&1; then \
  		echo -e "${RED}${APP_NAME}prometheus already exists${RESET}"; \
  	else \
		docker run -d --name=$(APP_NAME)prometheus -p $(PROMETHEUS_PORT):9090 --network=${APP_NAME}-network --add-host=host.docker.internal:host-gateway -v $(APP_NAME)_prometheus_data:/prometheus -v ./prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus:v3.4.0 >/dev/null 2>&1; \
		echo -e "${GREEN}${APP_NAME}prometheus created${RESET}"; \
	fi
.PHONY: prometheus-create

prometheus-delete: ### Deleting prometheus docker instance
	$(call delete-docker-service,prometheus)
.PHONY: prometheus-delete

network-create:
	@if docker network ls | grep "${APP_NAME}-network" >/dev/null 2>&1; then \
  		echo -e "${RED}${APP_NAME}-network already exists${RESET}"; \
  	else \
  	  	docker network create ${APP_NAME}-network >/dev/null 2>&1; \
  	  	echo -e "${GREEN}${APP_NAME}-network created${RESET}"; \
  	fi
.PHONY: network-create

network-delete:
	@if docker network inspect ${APP_NAME}-network >/dev/null 2>&1; then \
  		docker network rm "${APP_NAME}-network" >/dev/null 2>&1; \
  		echo -e "${GREEN}${APP_NAME}-network deleted${RESET}"; \
  	else \
  	  	echo -e "${RED}${APP_NAME}-network does not exist${RESET}"; \
  	fi
.PHONY: network-delete

clear: db-delete redis-delete grafana-delete prometheus-delete network-delete ### Cleaning up
.PHONY: clear
