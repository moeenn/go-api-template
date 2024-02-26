# golang native server
A Web server boilerplate created using Golang's `net/http` package. 
- See `go.mod` for other run-time dependencies.
- See `scripts/setup.sh` for development dependencies.

### Install dependencies and setting-up

```bash
# install dependencies
$ make setup

# generate jwt secret token (store inside .env file)
$ make gensecret
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
