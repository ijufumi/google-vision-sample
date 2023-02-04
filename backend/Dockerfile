FROM golang:1.19-bullseye as build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app /app/cmd/app/
RUN go build -o db /app/cmd/db/

FROM debian:bullseye-slim as deploy

WORKDIR /app

RUN apt-get update \
&& apt-get install -y ca-certificates \
&& rm -rf /var/cache/apt

COPY --from=build /app/app /app/app
COPY --from=build /app/db /app/db
COPY /app/migrate /app/migrate
COPY /app/scripts /app/scripts

RUN chmod +x /app/scripts/entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/app/scripts/entrypoint.sh"]