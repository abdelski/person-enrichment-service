# Person Enrichment Service

A REST API service that enriches person data with age, gender, and nationality from external APIs and stores it in PostgreSQL.

## Features

- Create persons with automatic data enrichment
- Get person by ID
- Update person details
- Delete persons
- Filter persons with pagination
- Swagger documentation
- Configurable via environment variables

## API Endpoints

| Method | Endpoint           | Description                          |
|--------|--------------------|--------------------------------------|
| POST   | /api/v1/persons    | Create a new person                  |
| GET    | /api/v1/persons/:id| Get person by ID                     |
| PUT    | /api/v1/persons/:id| Update person                        |
| DELETE | /api/v1/persons/:id| Delete person                        |
| GET    | /api/v1/persons    | Get filtered list of persons         |

## Getting Started

### Prerequisites

- Go
- PostgreSQL
- Make
- Docker

### Installation

1. Clone the repository
3. Run make target to start all the services:
   ```bash
   make docker-run