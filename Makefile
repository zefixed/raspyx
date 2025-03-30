include .env.example
export

all: db-create migrate-up
.PHONY: all

migrate-up: ### migration up
	goose -dir migrations postgres '$(PG_URL)' up
.PHONY: migrate-up

migrate-down: ### migration down
	goose -dir migrations postgres '$(PG_URL)' down
.PHONY: migrate-down

db-create: ### Creation postgres db docker instance
	docker run -p 5432:5432 --restart unless-stopped --name $(APP_NAME)db -e POSTGRES_PASSWORD=root -d postgres
	docker exec $(APP_NAME)db sh -c 'until pg_isready -U postgres; do sleep 0.5; done'
	docker exec -it $(APP_NAME)db psql -U postgres -c "CREATE DATABASE $(APP_NAME)db;"
.PHONY: db-create

db-delete: ### Deletion postgres db docker instance
	docker stop $(APP_NAME)db
	docker rm $(APP_NAME)db
.PHONY: db-delete

swag: ### Generation swagger documentation
	swag init -g cmd/app/main.go
.PHONY: swag
