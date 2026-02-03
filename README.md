# Go Fiber Simple REST API

A simple REST API built with Go Fiber framework, featuring JWT authentication, PostgreSQL database, and Swagger documentation.

## Features

- 🚀 Built with [Fiber](https://gofiber.io/) - Express-inspired web framework
- 🔐 JWT Authentication with token versioning
- 👤 User management with admin roles
- 📚 Books and Authors CRUD operations
- 🗄️ PostgreSQL database with GORM
- 📝 Swagger/OpenAPI documentation
- ✅ Request validation
- 🔄 Auto database migration

## Tech Stack

- **Go** 1.25.5+
- **Fiber** v2 - Web framework
- **GORM** - ORM library
- **PostgreSQL** - Database
- **JWT** - Authentication
- **Swagger** - API documentation
- **Validator** v10 - Request validation

## Project Structure

```
.
├── app/
│   ├── dto/              # Data Transfer Objects
│   ├── handlers/         # HTTP handlers
│   ├── middleware/       # Custom middlewares
│   ├── models/           # Database models
│   └── queries/          # Database queries
├── docs/                 # Swagger documentation
├── pkg/
│   ├── routes/           # API routes
│   ├── server/           # Server configuration
│   └── utils/            # Utility functions
├── platform/
│   └── database/         # Database connection
├── .env                  # Environment variables
├── main.go               # Application entry point
└── go.mod                # Go modules
```

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) 1.25.5 or higher
- [PostgreSQL](https://www.postgresql.org/download/) 12 or higher
- [Git](https://git-scm.com/downloads)

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/Alfinaditya/go-fiber-simple-rest-api.git
cd go-fiber-simple-rest-api
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Setup PostgreSQL Database

Create a new PostgreSQL database:

```sql
CREATE DATABASE your_database_name;
CREATE SCHEMA your_schema_name;
```

### 4. Configure Environment Variables

Create a `.env` file in the root directory:

```bash
touch .env
```

Add the following environment variables:

```env
# Database Configuration
DB_POSTGRES_HOST=localhost
DB_POSTGRES_PORT=5432
DB_POSTGRES_USER=your_postgres_user
DB_POSTGRES_PASSWORD=your_postgres_password
DB_POSTGRES_NAME=your_database_name
DB_POSTGRES_SCHEMA=your_schema_name

# JWT Configuration
JWT_SECRET_KEY=your_super_secret_jwt_key_here
```

**Example:**

```env
DB_POSTGRES_HOST=localhost
DB_POSTGRES_PORT=5432
DB_POSTGRES_USER=postgres
DB_POSTGRES_PASSWORD=postgres123
DB_POSTGRES_NAME=fiber_api_db
DB_POSTGRES_SCHEMA=public

JWT_SECRET_KEY=my_secret_key_12345
```

### 5. Run the application

```bash
go run main.go
```

The server will start on `http://localhost:3001`

### 6. Build for production

```bash
go build -o app
./app
```

## API Documentation

Once the application is running, visit the Swagger documentation at:

```
http://localhost:3001/docs/
```

## Authentication

This API uses JWT (JSON Web Tokens) for authentication. To access protected endpoints:

1. **Register** a new user via `/api/auth/register`
2. **Login** via `/api/auth/login` to receive a JWT token
3. Include the token in the `Authorization` header for protected routes:

```
Authorization: Bearer <your_jwt_token>
```

## Development

### Generate Swagger Documentation

If you make changes to the API handlers, regenerate the Swagger docs:

```bash
# Install swag CLI
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init
```

### Hot Reload (Optional)

For development with hot reload, you can use [Air](https://github.com/cosmtrek/air):

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

## Troubleshooting

### Database Connection Issues

If you encounter database connection errors:

1. Ensure PostgreSQL is running
2. Verify database credentials in `.env`
3. Check if the database and schema exist
4. Ensure PostgreSQL is accepting connections on the specified port

### Port Already in Use

If port 3001 is already in use, you can change it in `pkg/server/start_server.go`:

```go
log.Fatal(app.Listen(":3001")) // Change to your preferred port
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

**Alfin Aditya**
- GitHub: [@Alfinaditya](https://github.com/Alfinaditya)

## Acknowledgments

- [Fiber](https://gofiber.io/) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Swagger](https://swagger.io/) - API documentation

---

⭐ If you find this project useful, please consider giving it a star!
