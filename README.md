# Hear! Backend

Backend for the Hear! app using Golang and Fiber.

## Architecture

```mermaid
graph TB
    subgraph Development Tools
        Air[Air - Hot Reload]
        SQLC[SQLC - SQL Compiler]
        Docker[Docker & Docker Compose]
        Make[Make - Build Tool]
    end

    subgraph Application Architecture
        API[API Layer]
        DI[Dependency Injection]
        
        subgraph Core
            Users[Users Module]
            Shared[Shared Module]
        end
        
        subgraph Infrastructure
            DB[(Database)]
            Cache[(Cache)]
        end
    end

    API --> DI
    DI --> Core
    Core --> Infrastructure
    
    Air --> Application
    SQLC --> DB
    Docker --> Infrastructure
    Make --> Application
```

## Development Tools

- Air - Hot reloading during development
- SQLC - Type-safe SQL queries
- Docker & Docker Compose - Containerization
- Make - Build automation
- Go - Main programming language
- PostgreSQL - Database
- Fiber - Web framework

## Development

```bash
air
```

## Make commands

```bash
# Build the project
make build

# Run the project
make run
```
