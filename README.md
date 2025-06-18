# Hear Backend

A robust backend API for a university attendance management system built with Go, following clean architecture principles and best practices.

## 🏗️ Architecture

This project follows **Clean Architecture** (Hexagonal Architecture) with clear separation of concerns:

```
backend/
├── cmd/                    # Application entry point
├── internal/              # Private application code
│   ├── auth/             # Authentication module
│   ├── users/            # User management module
│   ├── courses/          # Course management module
│   ├── classrooms/       # Classroom management module
│   ├── attendance/       # Attendance tracking module
│   ├── qrcode/           # QR code generation module
│   ├── di/               # Dependency injection
│   ├── shared/           # Shared utilities and middleware
│   └── infrastructure/   # Database and external services
├── db/                   # Database migrations and queries
└── docs/                 # API documentation
```

### Module Structure
Each module follows the same pattern:
```
module/
├── domain/          # Business logic and entities
├── application/     # Use cases and handlers
├── infrastructure/  # Database repositories
└── setup.go        # Module initialization
```

## 🚀 Features

- **JWT Authentication** with access and refresh tokens
- **Role-based Access Control** (Admin, Teacher, Student)
- **Clean Architecture** with dependency injection
- **PostgreSQL** with SQLC for type-safe queries
- **Fiber** web framework for high performance
- **Comprehensive API** with OpenAPI documentation
- **Database migrations** with proper versioning

## 🛠️ Tech Stack

- **Language**: Go 1.24+
- **Framework**: Fiber v2
- **Database**: PostgreSQL
- **ORM**: SQLC (SQL Compiler)
- **Authentication**: JWT
- **Documentation**: Swagger/OpenAPI
- **Development**: Air (hot reload)

## 📋 Prerequisites

- Go 1.24 or higher
- PostgreSQL 12 or higher
- Make (optional, for convenience commands)

## 🚀 Quick Start

### 1. Clone the repository
```bash
git clone <repository-url>
cd backend
```

### 2. Install dependencies
```bash
go mod download
```

### 3. Set up environment variables
```bash
# Create .env file (optional, defaults are provided)
export JWT_SECRET="your-secret-key-here"
```

### 4. Set up the database
```bash
# Start PostgreSQL (using Docker)
docker-compose up -d postgres

# Run migrations
make migrate-up
```

### 5. Generate SQL code
```bash
make sqlc-generate
```

### 6. Run the application
```bash
# Development mode with hot reload
make dev

# Or run directly
go run cmd/main.go
```

The API will be available at `http://localhost:8080`

## 🛠️ Development

### Development Mode with Air
```bash
# Install Air (if not already installed)
go install github.com/cosmtrek/air@latest

# Run with hot reload
make dev
```

### Database Operations
```bash
# Run migrations
make migrate-up

# Rollback migrations
make migrate-down

# Generate SQL code
make sqlc-generate

# Reset database
make db-reset
```

### Testing
```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

## 📊 API Endpoints

### Authentication (Public)
- `POST /api/auth/login` - User login
- `POST /api/auth/register` - User registration
- `POST /api/auth/refresh` - Token refresh
- `POST /api/auth/logout` - User logout

### User Management (Admin Only)
- `POST /api/admin/users` - Create user
- `GET /api/admin/users/:id` - Get user by ID
- `GET /api/admin/users` - Get all users

### Course Management (Admin & Teachers)
- `POST /api/courses` - Create course
- `GET /api/courses` - Get all courses
- `GET /api/courses/:id` - Get course by ID
- `PUT /api/courses/:id` - Update course
- `DELETE /api/courses/:id` - Delete course

### Classroom Management (Admin & Teachers)
- `POST /api/classrooms` - Create classroom
- `GET /api/classrooms` - Get all classrooms
- `GET /api/classrooms/:id` - Get classroom by ID
- `PUT /api/classrooms/:id` - Update classroom
- `DELETE /api/classrooms/:id` - Delete classroom

### QR Code Management (Teachers Only)
- `POST /api/qrcodes` - Create QR code
- `GET /api/qrcodes/:id` - Get QR code by ID
- `GET /api/qrcodes/code/:code` - Get QR code by code
- `PUT /api/qrcodes/:id` - Update QR code
- `DELETE /api/qrcodes/:id` - Delete QR code

### Attendance Management (Teachers & Students)
- `POST /api/attendance` - Create attendance record
- `GET /api/attendance/:id` - Get attendance by ID
- `GET /api/attendance/student/:studentId` - Get student attendance
- `GET /api/attendance/class-schedule/:classScheduleId` - Get class schedule attendance
- `PUT /api/attendance/:id` - Update attendance
- `DELETE /api/attendance/:id` - Delete attendance

## 🔐 Authentication

The API uses JWT tokens for authentication:

1. **Login/Register** to get access and refresh tokens
2. **Include Bearer token** in Authorization header for protected routes
3. **Use refresh token** to get new access tokens when expired

### Token Structure
- **Access Token**: 15 minutes validity
- **Refresh Token**: 7 days validity
- **Claims**: User ID, Email, Role, Token Type

### Role-Based Access
- **Admin**: Full access to all endpoints
- **Teacher**: Course, classroom, QR code, and attendance management
- **Student**: Attendance viewing and creation

## 🗄️ Database

### Schema Overview
- **users**: User accounts and authentication
- **refresh_tokens**: JWT refresh token storage
- **universities**: University information
- **faculties**: Faculty departments
- **courses**: Course definitions
- **classrooms**: Physical classroom locations
- **course_sections**: Course sections with teachers
- **class_schedules**: Class scheduling
- **attendance**: Attendance records
- **qr_codes**: QR codes for attendance

### Migrations
Migrations are managed with SQL files in `db/migrations/`:
- `000001_init_schema.up.sql` - Initial schema
- `000002_add_refresh_tokens.up.sql` - Refresh tokens table

### SQLC
Type-safe SQL queries are generated using SQLC:
- Query files in `db/queries/`
- Generated Go code in `internal/infrastructure/db/`

## 🏗️ Project Structure

### Core Modules

#### Auth Module
- JWT token generation and validation
- Password hashing with bcrypt
- Refresh token management
- Authentication middleware

#### Users Module
- User CRUD operations
- Role management
- Password validation

#### Courses Module
- Course management
- Faculty associations
- Academic year tracking

#### Classrooms Module
- Physical classroom management
- Location tracking
- Capacity management

#### Attendance Module
- Attendance tracking
- Status management (present, absent, late)
- Class schedule associations

#### QR Code Module
- QR code generation
- Expiration management
- Course associations

### Shared Components

#### Middleware
- Authentication middleware
- Role-based authorization
- Request logging

#### Response Handling
- Standardized API responses
- Error handling
- Success/error formatting

## 🚀 Deployment

### Docker
```bash
# Build image
docker build -t hear-backend .

# Run container
docker run -p 8080:8080 hear-backend
```

### Environment Variables
```bash
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=hear
DB_PASSWORD=hear
DB_NAME=hear

# JWT
JWT_SECRET=your-secret-key

# Server
PORT=8080
```

## 📚 API Documentation

API documentation is available via Swagger UI:
- **URL**: `http://localhost:8080/swagger/`
- **OpenAPI Spec**: `http://localhost:8080/swagger/doc.json`

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific module tests
go test ./internal/auth/...
```

## 📝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Run the test suite
6. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🤝 Support

For support and questions, please open an issue in the repository.
