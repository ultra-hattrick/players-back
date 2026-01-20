# Players Service

This microservice handles the retrieval and management of player data for the Hattrick application. It connects to the Hattrick API to fetch player data and stores it in a Postgres database.

## Architecture

The project follows the **Hexagonal Architecture** (also known as Ports and Adapters).

### Directory Structure

- `cmd/`: Entry point of the application.
- `internal/`: Private application code.
    - `core/`: Contains the business logic and domain models.
        - `domain/`: Domain entities (e.g., `Player`).
        - `ports/`: Interfaces defining the input/output ports (e.g., `PlayerRepository`, `PlayerService`).
        - `services/`: Implementation of the business logic (business ports).
    - `adapters/`: Contains the implementation of the ports to interact with external systems.
        - `handler/http/`: HTTP handlers (Driving Adapter).
        - `repository/postgres/`: Database implementation using GORM (Driven Adapter).

## Prerequisites

- Go 1.22+
- Docker & Docker Compose (for DB)

## Configuration

The application is configured using environment variables. You can define them in a `.env` file in the root directory. Copy the `.env.example` file and fill in your values.

```bash
cp .env.example .env
```

### Environment Variables

| Variable | Description | Example |
| :--- | :--- | :--- |
| `GIN_MODE` | Gin mode (debug, release) | `debug` |
| `LOG_LEVEL` | Logging level | `info` |
| `PORT` | Application port | `8082` |
| `DB_POSTGRES_HOST` | Database host | `127.0.0.1` |
| `DB_POSTGRES_PORT` | Database port | `5432` |
| `DB_POSTGRES_USER` | Database user | `postgres` |
| `DB_POSTGRES_PASS` | Database password | `secret` |
| `DB_POSTGRES_NAME` | Database name | `mis_proyectos` |
| `DB_POSTGRES_SCHEMA` | Database schema | `ultra_hattrick` |
| `CONSUMER_KEY` | Hattrick Consumer Key | `your_key` |
| `CONSUMER_SECRET` | Hattrick Consumer Secret | `your_secret` |
| `OAUTH1_TOKEN` | Hattrick OAuth Token | `your_token` |
| `OAUTH1_TOKEN_SECRET` | Hattrick OAuth Token Secret | `your_token` |
| `BASE_RESOURCE_URL` | Hattrick API URL | `https://chpp.hattrick.org/chppxml.ashx` |

## Running the application

1.  **Set up the environment variables**:
    Create a `.env` file as described above.

2.  **Run with Go**:
    ```bash
    make run
    ```

## Testing

To run the unit tests:

```bash
go test ./...
```

or

```bash
go test -v ./internal/core/services/...
```

or

```bash
make test
```

## API Documentation

Swagger documentation is available at `/swagger/index.html` when the application is running.
