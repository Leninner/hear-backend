# Academic Management and Attendance System

A Go-based backend system for managing academic courses, student attendance, and QR code-based enrollment.

## Features

- User management with role-based access control (Admin, Teacher, Student)
- Course and class schedule management
- QR code-based course enrollment
- Attendance tracking with image metadata validation
- Location-based attendance verification

## Architecture

The system follows a clean architecture approach with the following layers:

- Domain: Core business logic and entities
- Service: Business logic implementation
- Repository: Data access layer using sqlc
- API: HTTP endpoints using Fiber

## Database Schema

The system uses PostgreSQL with the following main tables:

- users: User accounts with role-based access
- courses: Course information
- class_schedules: Class schedules for courses
- attendance: Student attendance records
- qr_codes: QR codes for course enrollment

## Getting Started

### Prerequisites

- Go 1.24 or later
- PostgreSQL 12 or later
- Docker and Docker Compose (optional)

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up the database:
   ```bash
   make db-migrate
   ```
4. Generate sqlc code:
```bash
   make sqlc
```
5. Run the application:
```bash
make run
```

### Development

- Run tests:
  ```bash
  make test
  ```
- Generate API documentation:
  ```bash
  make swagger
  ```
- Format code using gofmt:
  ```bash
  gofmt -s -w -d -e -r .
  ```

## API Documentation

The API documentation is available at `/swagger` when running the application.

## License

This project is licensed under the MIT License.
