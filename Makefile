ENV = .env
BINARY = web

setup:
	@go mod tidy
	@sh ./scripts/setup.sh

gensecret:
	go-token -length 64

dev:
	godotenv -f ${ENV} go run .

build:
	go mod tidy
	go build . -o ${BINARY}

prod:
	./${BINARY}
