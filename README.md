# go-http-psql

A sample of go http project directory structure.

We only use simple package to study basic knowledge, includes:

1. net/http
2. jackc/pgx
3. julienschmidt/httprouter

In golang, 
- Don't use MVC. Instead, split your code into different services.
- Don't put too much code in `func init()`

## Run PostgreSQL in docker

```bash
docker run --name pg13 \
-e POSTGRES_PASSWORD=Hello123@ \
-e POSTGRES_USER=hello \
-p 5432:5432 \
-d postgres:13.0-alpine
```

