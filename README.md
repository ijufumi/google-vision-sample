# Sample application for Google Cloud Vision

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
