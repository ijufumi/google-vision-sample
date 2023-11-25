# Sample application for Google Cloud Vision

This application is sample application to use Google Vision API.

It uses [Text Annotation API](https://cloud.google.com/vision/docs/ocr) for OCR.

## Application structures

```
.
├── backend
│   ├── cmd
│   │   ├── app
│   │   └── db
│   ├── configs
│   ├── docs
│   │   ├── db
│   │   └── infra
│   ├── internal
│   │   ├── common
│   │   │   ├── configs
│   │   │   ├── container
│   │   │   ├── context
│   │   │   ├── loggers
│   │   │   └── utils
│   │   ├── infrastructures
│   │   │   ├── database
│   │   │   └── google
│   │   ├── models
│   │   │   ├── entities
│   │   │   ├── service
│   │   │   └── value
│   │   ├── presentations
│   │   │   ├── handlers
│   │   │   ├── middlewares
│   │   │   └── router
│   │   └── usecases
│   │       └── repositories
│   ├── migration
│   └── scripts
├── deployment
│   ├── aws
│   └── gcp
├── docs
└── frontend
    ├── public
    └── src
        ├── components
        ├── configs
        ├── models
        ├── pages
        ├── repositories
        └── usecases

```

### `backend`

This is `Golang` application.

### `backend/cmd/app`

The `main.go` to launch an application.

### `backend/cmd/db`

The `main.go` to execute `golang-migrate` commands.

### `backend/configs`

This directory is storing configuration files. like google credential file.

### `backend/docs`

This directory is storing system documentation. like PlatUML file.

### `backend/migration`

This directory is for migration files to execute `golang-migrate` commands.

### `backend/internal/common/configs`

This package has configurations of an application.

### `backend/internal/common/container`

This package has a DI container.

### `backend/internal/common/context`

T.B.D

### `backend/internal/common/logger`

T.B.D

### `backend/internal/common/utils`

This package has utilities.

### `backend/internal/infrastructures/database`

This package is for database. such as connection, repositories and entities.

### `backend/internal/infrastructures/google`

This package is for google api. like a storage api and vision api.

### `backend/internal/models/entities`

This package has input / output models for API.

### `backend/internal/models/services`

This package has business logic services.


### `backend/internal/presentations/handlers`

This package has all controllers of http endpoint.

### `backend/internal/presentations/middlewares`

This package has originality middleware of GIN.

### `backend/internal/presentations/router`

This package determines uri path of http endpoint.

### `backend/scripts`

This directory has scripts for several purpose.

### `deployment/aws`

This code creates necessary resources on AWS.

### `deployment/gcp`

This code creates necessary resources on GCP.

### `docs`

This directory has files for `Github pages` or `README.md`.

### `frontend`

This is `React` application.

### `frontend/public`

This directory stores `HTML` and `favicon.ico` and more.

### `frontend/src`

This is codes of frontend.

### `frontend/src/components`

This package is for several components.

### `frontend/src/configs`

This package has a configuration of frontend app

### `frontend/src/models`

This package has `Data Access Object` that is used between backend api.

### `frontend/src/pages`

This package has source codes to show pages

### `frontend/src/repositories`

This package has classes to access the backend api.

### `frontend/src/usecases`

This package is the bridge code between pages and repositories.

## ERD

![ERD](./docs/tables.svg)

## How to

### Run app

```bash
docker compose up --build
```

### Migration

#### Create migration file

```bash
docker compose exec app /app/db create -n [name]
```

Or

```bash
docker compose exec app go run /backend/cmd/db/main.go create -n [name]
```

#### Apply migration file

```bash
docker compose exec app /app/db up
```

Or

```bash
docker compose exec app go run /backend/cmd/db/main.go up
```

#### Revert migration file

```bash
docker compose exec app /app/db down
```

Or

```bash
docker compose exec app go run /backend/cmd/db/main.go down
```

#### Clear all migration file

```bash
docker compose exec app /app/db drop
```

Or

```bash
docker compose exec app go run /backend/cmd/db/main.go drop
```

#### Confirm current migration version

```bash
docker compose exec app /app/db version
```

Or

```bash
docker compose exec app go run /backend/cmd/db/main.go version
```
