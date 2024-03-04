#! /bin/sh

echo "Installing project dependencies"
go mod tidy

# linting dependencies
go install honnef.co/go/tools/cmd/staticcheck@latest
go install github.com/kisielk/errcheck@latest
go install github.com/jgautheron/goconst/cmd/goconst@latest

# list out installation commands for binary dependencies
go install -v github.com/joho/godotenv/cmd/godotenv@latest
go install -v github.com/moeenn/go-token@latest
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
