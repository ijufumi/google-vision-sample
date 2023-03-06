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

#### Apply migration file

```bash
docker compose exec app /app/db up
```

#### Revert migration file

```bash
docker compose exec app /app/db down
```

#### Clear all migration file

```bash
docker compose exec app /app/db drop
```

#### Confirm current migration version

```bash
docker compose exec app /app/db version
```
