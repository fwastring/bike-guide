# Bike Guide API

A backend API for a bike guide application that provides information about bike routes and stops.

## Features

- PostgreSQL with PostGIS for storing and querying geographic data
- CouchDB for user data, authentication, and application state
- Redis for session management
- RESTful API built with Fiber
- Authentication middleware

## Prerequisites

- Go 1.23 or higher
- Docker and Docker Compose
- PostgreSQL with PostGIS extension
- CouchDB
- Redis

## Getting Started

1. Clone the repository
2. Navigate to the backend directory
3. Run `go mod tidy` to install dependencies
4. Start the databases using Docker Compose:
   ```
   docker-compose up -d
   ```
5. Run the application:
   ```
   go run cmd/app/main.go
   ```

## API Endpoints

### Public Endpoints

- `GET /` - Welcome message
- `POST /auth/login` - User login
- `POST /auth/logout` - User logout

### Protected Endpoints

- `GET /api/users/me` - Get current user information
- `GET /api/bikes/routes` - Get bike routes
- `GET /api/bikes/stops` - Get bike stops

## Environment Variables

The application can be configured using the following environment variables:

### Server Configuration

- `SERVER_PORT` - Server port (default: 3000)

### PostgreSQL Configuration

- `POSTGRES_HOST` - PostgreSQL host (default: localhost)
- `POSTGRES_PORT` - PostgreSQL port (default: 5432)
- `POSTGRES_USER` - PostgreSQL user (default: postgres)
- `POSTGRES_PASSWORD` - PostgreSQL password (default: postgres)
- `POSTGRES_DB` - PostgreSQL database name (default: gis_data)

### CouchDB Configuration

- `COUCHDB_HOST` - CouchDB host (default: localhost)
- `COUCHDB_PORT` - CouchDB port (default: 5984)
- `COUCHDB_USER` - CouchDB user (default: admin)
- `COUCHDB_PASSWORD` - CouchDB password (default: password)
- `COUCHDB_PROTOCOL` - CouchDB protocol (default: http)

### Redis Configuration

- `REDIS_HOST` - Redis host (default: localhost)
- `REDIS_PORT` - Redis port (default: 6379)
- `REDIS_PASSWORD` - Redis password (default: "")
- `REDIS_DB` - Redis database number (default: 0)

## License

MIT
