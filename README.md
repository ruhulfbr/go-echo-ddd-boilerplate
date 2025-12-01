<h1 align="center">Welcome to <span style="color:mediumseagreen">GO Echo DDD boilerplate</span></h1>

It's an API Skeleton project based on Echo framework.
My aim is reducing development time on default features that you can meet very often when your work on API.

## What's inside:

- Registration
- Authentication with JWT
- CRUD API for posts
- Migrations
- Request validation
- Environment configuration
- Docker development environment

## Usage

1. Copy .env.dist to .env and set the environment variables. There are examples for all the environment variables except
   COMPOSE_USER_ID, COMPOSE_GROUP_ID which are used by the linter. To get the current user ID, run in terminal:

   `echo $UID`

   In the .env file set these variables:

   `COMPOSE_USER_ID="username in current system"` - your username in system

   `COMPOSE_GROUP_ID="user uid"` - the user ID which you got in the terminal

2. Run your application using the command in the terminal:

   `docker-compose up`
3. Make requests to register a user (if necessary) and login.
4. After the successful login, copy a token from the response, use as `apiKey` in a form:
   `Bearer {token}`. For example:
```
   Bearer
   eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODk0NDA5NjYsIm9yaWdfaWF0IjoxNTg5NDM5OTY2LCJ1c2VyX2lkIjo1fQ.f8dSG3NxFLHwyA5-XIYALT5GtXm4eiH-motqtqAUBOI
```

Now, you are able to make requests which require authentication.
## Directories

```
├── cmd/
│   └── service/                 # Main service entry point
│       ├── bootstrap/           # App bootstrap (DB, server, handlers)
│       ├── wiring/              # Dependency wiring (repositories, services)
│       └── main.go              # Main entry file
├── config/
├── internal/
│   ├── common/                  # Shared utilities, config, and errors
│   │   ├── errors/
│   │   └── utils/
│   ├── domain/                  # Business domain logic
│   │   ├── post/
│   │   └── user/
│   ├── http/                    # HTTP layer
│   │   ├── handlers/
│   │   ├── middleware/
│   │   ├── requests/
│   │   ├── responses/
│   │   └── routes/
│   ├── infrastructure/          # Database, Redis, logging, and repositories
│   │   ├── database/
│   │   ├── logger/
│   │   ├── models/
│   │   ├── redis/
│   │   └── repositories/
│   └── services/                # Application services
│       ├── auth/
│       ├── oauth/
│       ├── post/
│       └── token/
├── go.mod
├── go.sum
├── Dockerfile
├── compose.yml
└── .env.dist
```

