# Entity service for golang

This repository serves as a skeleton for you to kickstart your next golang based project 
for anything within the TEC services backend.

## Get started

1. Git clone this repository
    ```bash
    # still not using SSH? Get better, start using SSH.
    git clone git@github.com:esportsclub/entity-service-golang.git <project_name>
    ```
2. Open `go.mod` and replace the first git url in first line with your projects' git link
3. Install all dependencies
   ```bash
   go mod tidy
   ```
4. Copy `.env.example` to `.env`
   ```
    cp .env.example .env
    ```
5. Fill in the `.env` with appropriate values so that connections are established
6. 🚀 Time to code

## Documentation

Docs have been written inside the code itself that follows and adheres to 
[GoDoc specifications](https://go.dev/blog/godoc). Every function and every file including
the `structs` have been documented thoroughly.

## Packages used

### Gorilla Mux - [Website](https://gorilla.github.io) | [Docs](https://github.com/gorilla/mux#examples)

An express.js like route handler and server that handles HTTP routes and responses.

### Go DotENV - [Docs](https://github.com/joho/godotenv)

Allows you to load `.env` files from your local folder and proxy `os.GetEnv()` calls. 

### MongoDB - [Website](https://www.mongodb.com/docs/drivers/go/current/) | [Docs](https://www.mongodb.com/docs/drivers/go/current/quick-reference/)

Self-explanatory - allows you to talk and make queries to mongodb.

### Redis - [Website](https://github.com/redis/go-redis) | [Docs](https://redis.uptrace.dev)

Self-explanatory - allows you to talk and make queries to redis.

## Directory structure

```
.
├── common
├── database
│   └── models
├── handlers
└── services
```
