ENV = .env
BINARY = web
MIGRATIONS_PATH = migrations

include ${ENV}

setup:
	@sh ./scripts/setup.sh

secret:
	@go-token -length 64

new_migration:
	@migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq $(name)

db_migrate:
	@migrate -path '${MIGRATIONS_PATH}' -database $(DB_CONNECTION) -verbose up

db_drop:
	@migrate -database $(DB_CONNECTION) drop

dev:
	godotenv -f ${ENV} go run .

build:
	go mod tidy
	go build . -o ${BINARY}

prod:
	./${BINARY}

lint:
	staticcheck ./... && errcheck ./...
