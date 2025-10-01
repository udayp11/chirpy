# Chirpy 🐦

A small RESTful API for a microblogging app (“chirps”) built in Go. It supports user accounts, authentication with access/refresh tokens, creating/deleting chirps, listing chirps with optional filtering, and basic admin/reset/metrics endpoints.


## Features

### Users
- Create account
- Update email/password
- Login (access token)
- Refresh token flow
### Chirps
- Create chirp (auth required)
- List chirps (sorted by created_at asc)
- Optional filter by author_id: GET /api/chirps?author_id=123
- Delete chirp (owner-only)
### Admin
- Reset database (dev only)
- Readiness and metrics endpoints
- Webhooks receiver (stub)

## ⚙️ Prerequisites
- Go (1.22+)
- PostgreSQL (or compatible database)
- - Goose (for database migrations): `go install github.com/pressly/goose/v3/cmd/goose@latest`
- (Optional) [sqlc](https://github.com/sqlc-dev/sqlc) — only required if you want to regenerate Go code from SQL queries



## 🚀 Setup & Run
### Configuration
Set environment variables:
- `DB_URL` — Postgres connection string (e.g., `postgres://user:pass@localhost:5432/chirpy?sslmode=disable`)
- `JWT_SECRET` — secret for signing tokens
- `PORT` — optional (default 8080)
- `POLKA_KEY` — webhook shared secret
- `PLATFORM` — environment/platform flag (e.g., `dev` or `prod`)

### Clone the repo
```bash
git clone https://github.com/udayp11/chirpy.git
cd chirpy 
```

### Run database migrations
``` bash
goose -dir ./sql/schema postgres "$DB_URL" up
```
### (Optional) Regenerate DB code with sqlc

If you’ve modified .sql files and want to regenerate type-safe Go code:
```bash
sqlc generate
```
### Start the server
```bash
go run .
```
The server will start (default :8080) and expose REST API endpoints.

### Health Check

Verify the server is running:

curl -i http://localhost:${PORT:-8080}/api/healthz


Expected output:

HTTP/1.1 200 OK
Content-Type: application/json

{"status":"ok"}

## 🌐 Deployment

- Ensure PostgreSQL is set up on your server

- Run Goose migrations

- Build and run the Go binary:

```bash

go build -o chirpy
./chirpy
```


## 🧪 API Overview

### Authentication

| Method | Endpoint        | Auth Required | Body | Response |
|--------|-----------------|---------------|------|----------|
| POST   | `/api/login`    | ❌             | `{"email": string, "password": string}` | `200 OK` → `{"token": "ACCESS_TOKEN", "refresh_token": "REFRESH_TOKEN", "email": "user@example.com"}` |
| POST   | `/api/refresh`  | ✅             | `{"refresh_token": string}` | `200 OK` → `{"token": "NEW_ACCESS_TOKEN"}` |

---

### Users

| Method | Endpoint        | Auth Required | Body | Response |
|--------|-----------------|---------------|------|----------|
| POST   | `/api/users`    | ❌             | `{"email": string, "password": string}` | `201 Created` → `{"id": 1, "email": "user@example.com"}` |
| PUT    | `/api/users`    | ✅             | `{"email": string, "password": string}` | `200 OK` → `{"id": 1, "email": "newemail@example.com"}` |

---

### Chirps

| Method | Endpoint                  | Auth Required | Body | Response |
|--------|---------------------------|---------------|------|----------|
| POST   | `/api/chirps`             | ✅             | `{"body": string}` | `201 Created` → `{"id": 1, "body": "chirp content", "created_at": "Time"}` |
| GET    | `/api/chirps`             | ❌             | Optional query: `author_id` | `200 OK` → `[{"id": 1, "body": "chirp content", "created_at": "Time"}]` |
| GET    | `/api/chirps/{chirpID}`   | ❌             | — | `200 OK` → `{"id": 1, "body": "chirp content", "created_at": "Time"}` |
| DELETE | `/api/chirps/{chirpID}`   | ✅             | — | `204 No Content` |

---

### Chirp Validation

| Method | Endpoint                  | Auth Required | Body | Response |
|--------|---------------------------|---------------|------|----------|
| POST   | `/api/validate_chirp`     | ❌             | `{"body": string}` | `200 OK` → `{"valid": true, "cleaned_body": "cleaned chirp content"}` |

---

### Admin / Operations

| Method | Endpoint       | Auth Required | Body | Response |
|--------|----------------|---------------|------|----------|
| POST   | `/admin/reset` | ✅             | —    | Resets the database (dev only) |
| GET    | `/admin/metrics` | ❌           | —    | Metrics endpoint |

---

### Health Check

| Method | Endpoint        | Auth Required | Body | Response |
|--------|-----------------|---------------|------|----------|
| GET    | `/api/healthz`  | ❌             | —    | `200 OK` → `{"status": "ok"}` |

---

### Webhooks

| Method | Endpoint             | Auth Required | Body | Response |
|--------|----------------------|---------------|------|----------|
| POST   | `/api/polka/webhooks` | ❌           | Implementation-specific | — |

---


## Contributing

Contributions are welcome!

- Fork the repo

- Create a branch: git checkout -b feature/xyz

- Make your changes & commit

- Push to your fork & open a Pull Request


## License

[MIT](https://choosealicense.com/licenses/mit/)
