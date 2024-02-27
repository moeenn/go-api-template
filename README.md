# golang native server
A Web server boilerplate created using Golang's `net/http` package. 
- See `go.mod` for other run-time dependencies.
- See `scripts/setup.sh` for development dependencies.

### Install dependencies and setting-up

```bash
# install dependencies
$ make setup

# generate jwt secret token (store inside .env file)
$ make secret
```

### Using application

```bash
# run in development mode
$ make dev

# build for production
$ make build

# run in production
$ make prod
```

### Database migrations

```bash
# create a new migration
$ make name=migration_name new_migration

# run pending migrations
$ make db_migrate

# drop everything and clear database
$ make db_drop
```
