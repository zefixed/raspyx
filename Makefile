include example.env.
export

migrate-up: ### migration up
	goose -dir migrations postgres '$(PG_URL)' up
.PHONY: migrate-up

migrate-down: ### migration down
	goose -dir migrations postgres '$(PG_URL)' down
.PHONY: migrate-down