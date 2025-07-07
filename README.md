# Web Crawler Test Task

This project contains a simple full-stack application written in Go (backend) and React/TypeScript (frontend).

## Prerequisites

- Go 1.20+
- Node.js 18+
- MySQL

## Setup

The repository includes populated `go.sum` and `package-lock.json` files to ensure reproducible builds.

Set the MySQL connection using `WEBCRAWLER_DSN` environment variable. Example:
`user:password@tcp(localhost:3306)/webcrawler?parseTime=true`.

1. **Database**

   Create database `webcrawler` and apply the migration:

   ```sql
   CREATE DATABASE webcrawler;
   USE webcrawler;
   source backend/migrations/001_init.sql;
   source backend/migrations/002_broken_links.sql;
   ```

2. **Backend**

   ```bash
   cd backend
   go mod tidy 
   export WEBCRAWLER_DSN="user:1234@tcp(localhost:3306)/webcrawler?parseTime=true"
   go run main.go
   ```

   The server listens on `:8080`.

3. **Frontend**

   ```bash
   cd frontend
   npm install  
   npm run dev
   ```

   The app will be available at `http://localhost:3000`.

4. **Docker Build**

   Build the full stack container (backend + compiled frontend):

   ```bash
   docker build -t webcrawler .
   ```

## Tests

Run front-end tests:

```bash
cd frontend
npm test
```

Run backend tests:

```bash
cd backend
go test ./...
```

## API

All requests must include header `Authorization: secret-token`.

- `POST /api/urls` – body `{"address": "https://google.com"}`
- `GET /api/urls` – list URLs
  - query params: `page`, `limit`, `search`, `status`, `html`
- `GET /api/urls/:id` – detail
  - includes list of broken links with status codes
- `DELETE /api/urls/:id` – delete
- `POST /api/urls/:id/restart` – re-crawl
- `POST /api/urls/:id/stop` – stop crawling


