# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Full-stack application for OCR/text detection using Google Cloud Vision API. Upload images/PDFs, extract text, and display results on a canvas.

## Tech Stack

- **Backend**: Go 1.21, Gin, GORM, PostgreSQL 16, Google Vision & Storage APIs
- **Frontend**: React 18, TypeScript, Evergreen UI, Konva (canvas), React Router 6
- **Infra**: Docker Compose, AWS CDK, GCP deployment

## Development Commands

### Run the full stack locally
```bash
docker compose up --build
# Frontend: http://localhost:3000
# Backend:  http://localhost:8080
# DB:       localhost:5432
```

### Database migrations (run inside container)
```bash
docker compose exec app /app/db up        # apply migrations
docker compose exec app /app/db down      # revert
docker compose exec app /app/db create -n <name>  # create new migration
docker compose exec app /app/db version   # check current version
```

### Backend
```bash
cd backend
go build ./...          # build
go test ./...           # run all tests
go test ./internal/models/service/...  # run tests for a specific package
```

### Frontend
```bash
cd frontend
yarn install            # install dependencies
yarn start              # dev server (port 3000)
yarn build              # production build
yarn test               # run tests (Jest)
```

### Linting (pre-commit hooks configured)
```bash
# Backend
cd backend && golangci-lint run ./...

# Frontend
cd frontend
npx prettier --check 'src/**/*.{ts,tsx}'
npx eslint 'src/**/*.{ts,tsx}'
```

### Mock generation (backend)
```bash
cd backend && mockery  # generates mocks for repository interfaces
```

## Architecture

### Backend (`backend/`)

Layered architecture with dependency injection (`uber/dig`):

- **`cmd/`** — Entry points: `app/` (HTTP server), `db/` (migration CLI via Cobra)
- **`internal/presentations/`** — HTTP layer: `router/` (Gin routes), `handlers/`, `middlewares/` (CORS, logging, response headers)
- **`internal/models/`** — Business logic: `service/` (DetectText, ImageConversion, Configuration), `entities/`, `value/`
- **`internal/usecases/repositories/`** — Repository interfaces
- **`internal/infrastructures/`** — Implementations: `database/repositories/` (GORM), `google/clients/` (Vision & Storage)
- **`internal/common/`** — Shared: `container/` (DI wiring), `configs/` (env-based config), `loggers/` (Zap)
- **`migration/`** — SQL migration files (golang-migrate)

### API Routes

```
GET    /api/health                 — health check
GET    /api/v1/detect_texts        — list jobs
GET    /api/v1/detect_texts/:id    — get job
POST   /api/v1/detect_texts        — create job (file upload)
DELETE /api/v1/detect_texts/:id    — delete job
GET    /api/v1/signed_urls         — get signed URL for file access
POST   /api/v1/configs/cors        — configure CORS
```

### Frontend (`frontend/`)

- **`pages/`** — App (home/upload), ResultPage (job results with canvas overlay)
- **`components/`** — FileUploadDialog, Image, Rect (Konva shapes), Loader
- **`repositories/`** — API clients extending BaseRepository (GET/POST/DELETE with error handling)
- **`usecases/`** — Business logic (JobUseCase)
- **`models/`** — TypeScript data models (Job, InputFile, OutputFile, ExtractedText, SignedUrl)
- **`configs/`** — `REACT_APP_ENDPOINT_URL` env var (defaults to `http://localhost:8080/api/v1`)

### Routes
- `/` — Home page (upload)
- `/jobs/:jobId` — Result page (displays extracted text over image)

## Configuration

Backend uses env vars (see `backend/.env.example`): DB connection, Google credentials, Storage bucket, signed URL expiration. Loaded via `godotenv` + `caarlos0/env`.

Frontend uses `REACT_APP_ENDPOINT_URL` env var for API endpoint configuration.

## Code Style

- **Backend**: golangci-lint
- **Frontend**: Prettier (no semicolons, trailing commas es5, tab width 2) + ESLint with TypeScript plugin
- **Frontend path alias**: `~/` maps to `src/` root in imports
