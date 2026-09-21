# ⚡ Altera

**Altera** is a modern full-stack application built with a focus on clean architecture, type-safe APIs, secure authentication, and scalable backend design.

The project uses **Go, PostgreSQL, Docker, and database migrations** to create a reliable backend foundation while keeping development and deployment environments reproducible.

> 🚧 **Altera is currently under active development.** Features, architecture, and APIs may change as the project evolves.

---

## ✨ Features

### 🔐 Authentication

Altera provides a secure authentication foundation for managing users and protected application resources.

Planned and implemented authentication functionality includes:

- User registration
- User login
- Password hashing
- Token-based authentication
- Protected API routes
- Authenticated user sessions
- User-specific resource access

---

### 🗄️ PostgreSQL Database

Altera uses **PostgreSQL** as its primary relational database.

The database layer is designed around:

- Structured relational data
- Explicit schema migrations
- Referential integrity
- Reproducible development environments
- Clear separation between application and database logic

Database schema changes are tracked through version-controlled migrations rather than manual database modifications.

---

### 🐳 Docker Development Environment

PostgreSQL runs inside Docker to keep local development predictable and isolated.

```text
┌─────────────────────────────┐
│         Altera App          │
│                             │
│        Go Backend API       │
└──────────────┬──────────────┘
               │
               │ SQL
               ▼
┌─────────────────────────────┐
│          Docker             │
│                             │
│        PostgreSQL           │
│          :5433              │
└─────────────────────────────┘
```

The PostgreSQL container is exposed on:

```text
127.0.0.1:5433
```

Port `5433` is used because the default PostgreSQL port `5432` may already be occupied by a local PostgreSQL installation.

---

## 🛠️ Tech Stack

| Layer | Technology |
| --- | --- |
| Backend | Go |
| Database | PostgreSQL |
| Database Migrations | SQL Migrations |
| Containerization | Docker |
| API Architecture | REST |
| Configuration | Environment Variables |
| Version Control | Git / GitHub |

The technology stack will continue to expand as development progresses.

---

## 🏗️ Architecture

Altera follows a layered backend architecture designed to keep responsibilities separated.

```text
Client
   │
   │ HTTP / REST
   ▼
┌────────────────────────────┐
│          Router            │
│                            │
│ Routes & Middleware        │
└─────────────┬──────────────┘
              │
              ▼
┌────────────────────────────┐
│         Handlers           │
│                            │
│ HTTP Request / Response    │
└─────────────┬──────────────┘
              │
              ▼
┌────────────────────────────┐
│       Business Logic       │
│                            │
│ Application Services       │
└─────────────┬──────────────┘
              │
              ▼
┌────────────────────────────┐
│        Data Layer          │
│                            │
│ PostgreSQL Queries         │
└─────────────┬──────────────┘
              │
              ▼
┌────────────────────────────┐
│        PostgreSQL          │
│                            │
│ Docker Container           │
└────────────────────────────┘
```

This structure helps keep HTTP concerns, business rules, and database operations independent from one another.

---

## 📁 Project Structure

A typical Altera project structure follows:

```text
Altera/
│
├── backend/
│   ├── cmd/
│   │   └── ...
│   │
│   ├── internal/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   ├── services/
│   │   ├── repository/
│   │   └── ...
│   │
│   ├── migrations/
│   │   ├── 000001_*.up.sql
│   │   ├── 000001_*.down.sql
│   │   └── ...
│   │
│   ├── .env
│   └── ...
│
├── docker-compose.yml
│
└── README.md
```

The exact structure may evolve as additional Altera features are implemented.

---

## 🚀 Getting Started

### Prerequisites

Make sure the following tools are installed:

- Go
- Docker Desktop
- Git
- PostgreSQL migration tooling

Verify your installations:

```powershell
go version
docker --version
docker compose version
migrate -version
```

---

## 1. Clone the Repository

```powershell
git clone <your-altera-repository-url>
cd Altera
```

---

## 2. Start PostgreSQL

Start the development database:

```powershell
docker compose up -d
```

Verify that the container is running:

```powershell
docker compose ps
```

The development PostgreSQL server should be available at:

```text
127.0.0.1:5433
```

---

## 3. Configure Environment Variables

Create the backend environment file if one is not already present.

```env
DATABASE_URL=postgres://<username>:<password>@127.0.0.1:5433/altera?sslmode=disable
```

Add any authentication secrets or additional configuration required by the application.

> Never commit `.env` files containing real credentials or secrets.

---

## 4. Run Database Migrations

Apply all available migrations:

```powershell
migrate -path migrations -database "$env:DATABASE_URL" up
```

If `DATABASE_URL` is not loaded into the current PowerShell session, it can be assigned first:

```powershell
$env:DATABASE_URL="postgres://<username>:<password>@127.0.0.1:5433/altera?sslmode=disable"
```

Then run:

```powershell
migrate -path migrations -database "$env:DATABASE_URL" up
```

---

## ↩️ Testing Migration Rollbacks

During development, migration rollback behaviour should also be verified.

Roll back the latest migration:

```powershell
migrate -path migrations -database "$env:DATABASE_URL" down 1
```

Reapply it:

```powershell
migrate -path migrations -database "$env:DATABASE_URL" up
```

This verifies that both the `.up.sql` and `.down.sql` migrations behave correctly.

> ⚠️ Be careful with `down` migrations once the database contains important data. Rollbacks may delete tables, columns, or records depending on the migration.

---

## 5. Run the Backend

From the backend directory:

```powershell
go run ./cmd/...
```

Or run the application's specific entry point:

```powershell
go run ./cmd/api
```

The exact command depends on the current Altera project structure.

---

## 🗃️ Database Migrations

Database migrations are stored as paired SQL files.

```text
migrations/
├── 000001_initial_schema.up.sql
├── 000001_initial_schema.down.sql
├── 000002_example.up.sql
└── 000002_example.down.sql
```

### Apply all migrations

```powershell
migrate -path migrations -database "$env:DATABASE_URL" up
```

### Roll back one migration

```powershell
migrate -path migrations -database "$env:DATABASE_URL" down 1
```

### Roll back all migrations

```powershell
migrate -path migrations -database "$env:DATABASE_URL" down
```

### Check migration version

```powershell
migrate -path migrations -database "$env:DATABASE_URL" version
```

Migration files should always be committed alongside the application code that depends on them.

---

## 🔒 Security

Altera is being developed with security considered throughout the application architecture.

### Authentication

Protected endpoints should require authenticated requests before accessing private resources.

Authentication responsibilities include:

- Credential validation
- Secure password hashing
- Token verification
- Authentication middleware
- Expired or invalid token rejection

---

### Authorization

Authentication answers:

```text
Who is making this request?
```

Authorization determines:

```text
Is this user allowed to perform this action?
```

Database operations involving private resources should therefore be scoped to the authenticated user whenever appropriate.

---

### Input Validation

Incoming API data should be validated before reaching application or database logic.

Validation should cover areas such as:

- Required fields
- Maximum string lengths
- Allowed values
- Numeric ranges
- Email formatting
- IDs
- Request body structure

Invalid input should return predictable client errors rather than reaching the database layer.

---

### Password Security

Passwords should never be stored in plaintext.

Only secure password hashes should be persisted.

```text
Password
   │
   ▼
Password Hashing
   │
   ▼
Secure Hash
   │
   ▼
PostgreSQL
```

The original password should never be recoverable from the stored database value.

---

### Secrets

Secrets should be supplied through environment variables.

Examples include:

```env
DATABASE_URL=
JWT_SECRET=
```

Secrets must not be:

- Hardcoded into Go source files
- Included in frontend bundles
- Committed to Git
- Written to application logs

---

## 🧪 Development

Before committing changes, verify the project builds and tests successfully.

```powershell
go test ./...
```

Build the project:

```powershell
go build ./...
```

Format Go source files:

```powershell
go fmt ./...
```

Run static analysis:

```powershell
go vet ./...
```

A healthy development cycle should therefore include:

```powershell
go fmt ./...
go vet ./...
go test ./...
go build ./...
```

---

## 🐳 Docker Commands

Start services:

```powershell
docker compose up -d
```

View running services:

```powershell
docker compose ps
```

View logs:

```powershell
docker compose logs
```

Follow logs:

```powershell
docker compose logs -f
```

Stop services:

```powershell
docker compose down
```

Stop services and remove associated volumes:

```powershell
docker compose down -v
```

> Removing volumes deletes persisted development database data.

---

## ⚙️ Environment Variables

| Variable | Purpose |
| --- | --- |
| `DATABASE_URL` | PostgreSQL connection string |
| `JWT_SECRET` | Secret used for authentication tokens |
| Additional variables | Added as Altera's services expand |

Example development connection:

```env
DATABASE_URL=postgres://<username>:<password>@127.0.0.1:5433/altera?sslmode=disable
```

---

## 🛣️ Development Roadmap

Altera is under active development.

Current and planned work includes:

- [x] Go backend foundation
- [x] PostgreSQL development database
- [x] Dockerized database environment
- [x] Database migration system
- [ ] Complete authentication flow
- [ ] Authorization middleware
- [ ] Core application API
- [ ] Request validation
- [ ] Structured error handling
- [ ] Logging
- [ ] Automated testing
- [ ] Frontend integration
- [ ] Production deployment
- [ ] CI/CD pipeline

The roadmap will evolve as the application's core feature set is implemented.

---

## 🚢 Production Considerations

Before a production release:

- Use production database credentials
- Use HTTPS
- Rotate development secrets
- Use strong authentication secrets
- Configure restrictive CORS policies
- Validate all incoming requests
- Apply authorization checks to protected resources
- Configure production database backups
- Run migrations through a controlled deployment process
- Add structured application logging
- Add automated tests
- Configure CI/CD checks
- Avoid exposing PostgreSQL directly to the public internet

---

## 📜 License

This project is currently developed for educational, portfolio, and software engineering purposes.

---

## 👨‍💻 Author

**Chethiya Bandara**

Computer Science · Software Engineering

Altera is an ongoing project exploring **Go backend engineering, PostgreSQL, containerized development, database design, API architecture, and secure application development**.
